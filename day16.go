package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

type nodo struct{
	nome string
	flowRate int
}

type list struct{
	head *nodes
	tail *nodes
}

type nodes struct{
	item nodo
	next *nodes
}

type stato struct{
	n nodo 
	tempo int 
	bitSet int
}


var stati map[stato]int
var graph map[nodo]map[nodo]int
var mappaNodi map[nodo]int

func main(){
	 start := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	adjacentList := make(map[nodo]*list,0)
	listaNodi := make([]nodo,0)

	for scanner.Scan(){
		line := scanner.Text()
		inputs := strings.Split(line," ")
		/*
		 [1] NOME
		 [4] rate
		 [9] > valvole
		*/
		inputs[4] = inputs[4][strings.Index(inputs[4],"=")+1:len(inputs[4])-1]
		
		value,_ := strconv.Atoi(inputs[4])
		tempNodo := nodo{inputs[1],value}
		adjacentList[tempNodo] = nil

		for i := 9; i < len(inputs) ;i++{
			temp := inputs[i];
			if i != len(inputs) - 1{
				temp = temp[:len(temp)-1]
			}
			nuovoNodo := nodo{temp,-1}

			if adjacentList[tempNodo] == nil{
				lista := new(list)
				lista.head = &(nodes{nuovoNodo,nil})
				lista.tail = lista.head
				adjacentList[tempNodo] = lista
			}else{
				lista := adjacentList[tempNodo]
				temp := new(nodes)
				temp.item = nuovoNodo
				temp.next = lista.head
				lista.head = temp
				adjacentList[tempNodo] = lista
			}
		}
		listaNodi = append(listaNodi,tempNodo)
	}

	for _,v := range adjacentList{
		f := v.head
		for ;f!= nil;{
			for _,v := range listaNodi{
				if f.item.nome == v.nome{
					f.item.flowRate = v.flowRate
					break
				}
			}
			f = f.next
		}
	}
	graph = FloydWarshall(adjacentList)

	nodiInteressanti := make([]nodo,0)

	for _,v := range listaNodi{
		if v.flowRate != 0 || v.nome == "AA"{
			nodiInteressanti = append(nodiInteressanti,v)
		}
	}

	mappaNodi = make(map[nodo]int)

	for k,v := range nodiInteressanti{
		mappaNodi[v] = k
	}
	stati = make(map[stato]int)
	b := (1 << len(nodiInteressanti))-1

	max := 0

	for i := 1; i < b/2; i += 2{
		c:= 0
		temp1 := i
		for ;temp1 != 0;{
			s := temp1%10
			if s == 1{
				c++
			}
			temp1 /= 10
		}

		if c < len(nodiInteressanti)/2 - 5 || c > len(nodiInteressanti) + 5 {
			continue
		}

		temp := DFS(26,nodo{"AA",0}, i) + DFS(26,nodo{"AA",0},b^i)

		if temp > max{
			max = temp
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Parte 1 : " ,DFS(30,nodo{"AA",0}, 0))
	fmt.Println("Parte 2 (se da risultato errato rimuovere linea 125) : " ,max)
	fmt.Println("Tempo: ", elapsed)
}

func DFS(tempo int, nodoCorrente nodo, nodiVisitati int) int{
	if tempo <= 0{
		return 0
	}
	max := 0

	for k,v := range graph[nodoCorrente]{
		if k.flowRate == 0{
			continue
		}
		if tempo-v-1 <= 0{
			continue
		}
		if k == nodoCorrente{
			continue
		}
		
		temp := 0

		if val, ok := stati[stato{k,tempo-v,nodiVisitati}]; ok{
			if val > max{
				max = val
			}
			continue
		}	

		if  (nodiVisitati & (1 << mappaNodi[k])) == 0{
			tempNodiVisitati := nodiVisitati
			tempNodiVisitati += 1 << mappaNodi[k]
			temp = DFS(tempo-v-1,k,tempNodiVisitati)
			temp += k.flowRate*(tempo-v-1)
		}
		tempNotOpen := DFS(tempo-v,k,nodiVisitati)

		if temp > max {
			max = temp
		}
		if tempNotOpen > max{
			max = tempNotOpen
		}
		max2 := temp 
		if tempNotOpen > max2{max2 = tempNotOpen}
		stati[stato{k,tempo-v,nodiVisitati}] = max2
	}
	return max
}
/*
func setValue(nodoCorrente nodo,tempo int, nodiVisitati int,val int){
	for k,v := range stati{
		if nodiVisitati == v.bitSet && nodoCorrente == v.n && tempo == v.tempo{
			stati[k].valore = val
			return
		}
	} 
}

func contains(nodoCorrente nodo, tempo int, nodiVisitati int) int{
	for _,v := range stati{
		if nodoCorrente == v.n && tempo == v.tempo && nodiVisitati == v.bitSet {
			return v.valore
		}
	} 

	var temp stato 
	temp.n = nodoCorrente
	temp.bitSet = nodiVisitati
	temp.tempo = tempo
	temp.valore = -1
	stati = append(stati,temp)
	return -1
}*/

func visitato(n nodo, nodiVisitati []nodo) bool{
	for _,v := range nodiVisitati{
		if n == v{
			return true
		}
	}
	return false
}


func FloydWarshall(adjacentList map[nodo]*list) map[nodo]map[nodo]int{
	sol := make(map[nodo]map[nodo]int)

	for k,_ := range adjacentList{
		sol[k] = make(map[nodo]int)
		for j,_ := range adjacentList{
			if k == j{
				sol[k][j] = 0
			}else if !lookFor(j,adjacentList[k]){
				sol[k][j] = 1
			}else{
				sol[k][j] = 999999
			}
		}
	}
	for k,_ := range adjacentList{
		for j,_ := range adjacentList{
			for i,_ := range adjacentList{
				if sol[j][i] > sol[j][k] + sol[k][i]{
					sol[j][i] = sol[j][k] + sol[k][i]
				}
			}
		}
	}
	return sol
}


func lookFor(x nodo ,lista *list) bool{
	temp := lista.head

	for;temp != nil; temp = temp.next{
		if temp.item == x{
			return false
		}
	}
	return true
}
