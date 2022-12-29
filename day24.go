package main

import(
	"fmt"
	"bufio"
	"os"
)

type pos struct{
	x,y int
}


type nodo struct{
	currentPosition pos 
	currentMap map[pos][]string
	depth int
}

type state struct{
	blizzards [][]string
	currentPosition pos
}

var maxX,maxY,minX,minY int
var states []state


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	mappa := make(map[pos][]string)

	minX = 1
	minY = 1
	states = make([]state,0)
	startingPoint := pos{1,0}
	var endingPoint pos

	j:=0
	for scanner.Scan(){
		line := scanner.Text()
		for i := 0; i < len(line);i++{
			mappa[pos{i,j}] = make([]string,0)
			mappa[pos{i,j}] = append(mappa[pos{i,j}],line[i:i+1])
			if line[i:i+1] == "E"{
				endingPoint = pos{i,j}
			}

			if line[i:i+1] == "S"{
				startingPoint	 = pos{i,j}
			}
		}
		j++
	}
	maxY = endingPoint.y - 1
	maxX = endingPoint.x 
	coda := make([]nodo,0)

	coda = append(coda,nodo{startingPoint,mappa,0})
	fmt.Println(startingPoint)
	for;coda != nil;{

		n := coda[0]
		//fmt.Println(n.currentPosition)
		if n.currentPosition == endingPoint{
			break
		}
		tempMappa := aggiornaMappa(n.currentMap)

		var temp state
		temp.currentPosition = n.currentPosition
		temp.blizzards = make([][]string,0)
		for i := 0; i <= maxY;i++{
			t := make([]string,0)
			for d := 0; d <= maxX; d++{
				myTempString := ""
				f := tempMappa[pos{d,i}]
				for k := 0; k < len(f);k++{
					myTempString += f[k]
				}
				t = append(t,myTempString)
			}
			temp.blizzards = append(temp.blizzards,t)
		}
		if stateAlreadyPresent(temp){
			coda = coda[1:]
			continue
		}
		states = append(states,temp)
		if _,ok := tempMappa[n.currentPosition]; !ok && n.depth+1 < 20{
			coda = append(coda,nodo{n.currentPosition,tempMappa,n.depth+1})
		}

		if _,ok := tempMappa[pos{n.currentPosition.x+1,n.currentPosition.y}]; !ok && n.depth+1 < 20{
			coda = append(coda,nodo{pos{n.currentPosition.x+1,n.currentPosition.y},tempMappa,n.depth+1})
		}

		if _,ok := tempMappa[pos{n.currentPosition.x-1,n.currentPosition.y}]; !ok && n.depth+1 < 20{
			coda = append(coda,nodo{pos{n.currentPosition.x-1,n.currentPosition.y},tempMappa,n.depth+1})
		}

		if _,ok := tempMappa[pos{n.currentPosition.x,n.currentPosition.y-1}]; !ok  && n.depth+1 < 20{
			coda = append(coda,nodo{pos{n.currentPosition.x,n.currentPosition.y-1},tempMappa,n.depth+1})
		}

		if _,ok := tempMappa[pos{n.currentPosition.x,n.currentPosition.y+1}]; !ok && n.depth+1 < 20{
			coda = append(coda,nodo{pos{n.currentPosition.x,n.currentPosition.y+1},tempMappa,n.depth+1})
		}

		if len(coda) == 1{
			coda = nil
		}else{
			coda = coda[1:]
		}
	}
}

func stateAlreadyPresent(temp state) bool{

	for _,v := range states{
		for i := 0; i < maxY; i++{
				k := temp.blizzards[i]
				for d := 0; d <= maxX; d++{
					if k[d] != v.blizzards[i][d]{
						return false
					}
				}
			}
	}
	return true
}

func aggiornaMappa(mappa map[pos][]string) (map[pos][]string){
	res := make(map[pos][]string)
	for k,v := range mappa{
		for i := 0; i < len(v); i++{
			if v[i] == ">"{
				x := k.x+1
				if x == maxX+1{
					x = 1
				}
				if _,ok := res[pos{x,k.y}]; !ok{
					res[pos{x,k.y}] = make([]string,0)
				}
				res[pos{x,k.y}] = append(res[pos{x,k.y}],v[i])
			}else if v[i] == "^"{
				y := k.y-1
				if y == minY-1{
					y = maxY
				}
				if _,ok := res[pos{k.x,y}]; !ok{
					res[pos{k.x,y}] = make([]string,0)
				}
				res[pos{k.x,y}] = append(res[pos{k.x,y}],v[i])
			}else if v[i] == "<"{
				x := k.x-1
				if x == minX-1{
					x = maxX
				}
				if _,ok := res[pos{x,k.y}]; !ok{
					res[pos{x,k.y}] = make([]string,0)
				}
				res[pos{x,k.y}] = append(res[pos{x,k.y}],v[i])
			}else if v[i] == "v"{
				y := k.y+1
				if y == maxY+1{
					y = 1
				}
				if _,ok := res[pos{k.x,y}]; !ok{
					res[pos{k.x,y}] = make([]string,0)
				}
				res[pos{k.x,y}] = append(res[pos{k.x,y}],v[i])
			}else if v[i] == "#"{
				res[pos{k.x,k.y}] = make([]string,1)
				res[pos{k.x,k.y}][0] = "#"
			}
		}
	}


	return res
}