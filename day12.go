package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)
type nodo struct{
	h int
	x int 
	y int
}

type arco struct{
	p1 nodo
	p2 nodo
}
type punto struct{
	n nodo
	dist int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	matrix := make([][]string,0)
	listaArchi := make([]arco,0)
	edges := make(map[nodo]punto)
	startingPoints := make([]nodo,0)
	var ris nodo
	for scanner.Scan(){
		line := scanner.Text()
		arr := strings.Split(line,"")
		matrix = append(matrix,arr)
	}
	// i = y, j = x
	for i := 0; i < len(matrix); i++{
		for j := 0; j < len(matrix[0]); j++{
			var nuovoNodo nodo
			nuovoNodo.h = int(matrix[i][j][0])
			nuovoNodo.x = j;
			nuovoNodo.y = i;
			var p punto
			p.n = nuovoNodo
			if nuovoNodo.h == 97{
				startingPoints = append(startingPoints,nuovoNodo)
			//	partenza = nuovoNodo
				p.n = nuovoNodo
				p.dist = 0
			}else{p.dist = 1000000}
			if nuovoNodo.h == 69{
				nuovoNodo.h = 122
				fmt.Println(nuovoNodo.x," ",nuovoNodo.y)
				ris = nuovoNodo
				p.n = nuovoNodo
				p.dist = 1000000
			}
			edges[nuovoNodo] = p
			if j > 0{
				if nuovoNodo.h+1 >= int(matrix[i][j-1][0]){
					if (matrix[i][j-1] != "E" || (matrix[i][j-1] == "E") && nuovoNodo.h+1 >= 122){
						var nuovoNodo1 nodo
						nuovoNodo1.h = int(matrix[i][j-1][0])
						if matrix[i][j-1] == "E"{
							nuovoNodo1.h = 122
						}
						nuovoNodo1.x = j-1;
						nuovoNodo1.y = i;
						nuovoArco := arco{
							p1: nuovoNodo,
							p2: nuovoNodo1,
						}
						listaArchi = append(listaArchi,nuovoArco)
					}
				}
			}

			if j < len(matrix[0])-1{
				if nuovoNodo.h+1 >= int(matrix[i][j+1][0]) || (matrix[i][j+1] == "E" && nuovoNodo.h+1 >= 122){
					if (matrix[i][j+1] != "E" || (matrix[i][j+1] == "E") && nuovoNodo.h+1 >= 122){
						var nuovoNodo1 nodo
						nuovoNodo1.h = int(matrix[i][j+1][0])
						if matrix[i][j+1] == "E"{
							nuovoNodo1.h = 122
						}
						nuovoNodo1.x = j+1;
						nuovoNodo1.y = i;	
						nuovoArco := arco{
							p1: nuovoNodo,
							p2: nuovoNodo1,
						}
						listaArchi = append(listaArchi,nuovoArco)
					}
				}
			}

			if i > 0{
				if nuovoNodo.h+1 >= int(matrix[i-1][j][0]) || (matrix[i-1][j] == "E" && nuovoNodo.h+1 >= 122){
					if (matrix[i-1][j] != "E" || (matrix[i-1][j] == "E") && nuovoNodo.h+1 >= 122){
						var nuovoNodo1 nodo
						nuovoNodo1.h = int(matrix[i-1][j][0])
						if matrix[i-1][j] == "E"{
							nuovoNodo1.h = 122
						}
						nuovoNodo1.x = j;
						nuovoNodo1.y = i-1;	
						nuovoArco := arco{
							p1: nuovoNodo,
							p2: nuovoNodo1,
						}
						listaArchi = append(listaArchi,nuovoArco)
					}
				}
			}

			if i < len(matrix)-1{
				if nuovoNodo.h+1 >= int(matrix[i+1][j][0]) || (matrix[i+1][j] == "E" && nuovoNodo.h+1 >= 122){
					if (matrix[i+1][j] != "E" || (matrix[i+1][j] == "E") && nuovoNodo.h+1 >= 122){
						var nuovoNodo1 nodo
						nuovoNodo1.h = int(matrix[i+1][j][0])
						if matrix[i+1][j] == "E"{
							nuovoNodo1.h = 122
						}
						nuovoNodo1.x = j;
						nuovoNodo1.y = i+1;
						nuovoArco := arco{
							p1: nuovoNodo,
							p2: nuovoNodo1,
						}
						listaArchi = append(listaArchi,nuovoArco)
					}
				}
			}

		}
	}
	res := make([]punto,0)
	for q := 0; q < len(startingPoints);q++{
		nodiFrontiera := make([]nodo,0)
		nodiFrontiera = append(nodiFrontiera,startingPoints[q])
		temp := edges[startingPoints[q]]
		temp.dist = 0
		edges[startingPoints[q]] = temp
		for ;len(nodiFrontiera) != 0 ;{
			for i := 0; i  < len(listaArchi); i++{
				if(listaArchi[i].p1 == nodiFrontiera[0]){
					if edges[listaArchi[i].p2].dist > edges[nodiFrontiera[0]].dist+1 {
						var temp punto
						temp.n = nodiFrontiera[0]
						temp.dist = edges[nodiFrontiera[0]].dist+1
						edges[listaArchi[i].p2] = temp
						nodiFrontiera = append(nodiFrontiera,listaArchi[i].p2)
					}
				}
			}
			nodiFrontiera = nodiFrontiera[1:]
		}
		res  = append(res,edges[ris])
		//fmt.Println(res[q])
		for k,_ := range edges{
			temp := edges[k]
			temp.dist = 1000000
			edges[k] = temp
		}
	}
	for k,_ := range res{
		fmt.Println(res[k])
	}
}
