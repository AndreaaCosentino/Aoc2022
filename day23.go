package main 

import(
	"fmt"
	"bufio"
	"strings"
	"os"
)

type pos struct{
	x,y int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	posizioni := make(map[pos]int,0)

	j := 0

	for scanner.Scan(){
		line := scanner.Text()

		input := strings.Split(line,"")

		for k,v := range input{
			if v == "#"{
				posizioni[pos{k,j}] = 1
			}
		}

		j++
	}
	fmt.Println(len(posizioni))
	list := "NSWE"
	f := 0
	for i := 0;; i++{
		mappaTemp := make(map[pos]int)
		mappaElfi := make(map[pos]pos)

		for k,_ := range posizioni{
			if !checkAround(k,posizioni){
				d := f
				for n := 0; n < 4; n++{
					if list[d] == 'N' && checkNorth(k,posizioni) {
						mappaTemp[pos{k.x,k.y-1}]++
						mappaElfi[k] = pos{k.x,k.y-1}
						break
					}else if list[d] == 'S' && checkSouth(k,posizioni){
						mappaTemp[pos{k.x,k.y+1}]++
						mappaElfi[k] = pos{k.x,k.y+1}
						break
					}else if list[d] == 'W' && checkWest(k,posizioni){
						mappaTemp[pos{k.x-1,k.y}]++
						mappaElfi[k] = pos{k.x-1,k.y}
						break
					}else if list[d] == 'E' && checkEast(k,posizioni){
						mappaTemp[pos{k.x+1,k.y}]++
						mappaElfi[k] = pos{k.x+1,k.y}
						break
					}
					d++
					d = d%4
				}
			}
		}
		f++
		f = f%4
		moved := false
		for k,v := range mappaElfi{
			if mappaTemp[v] == 1{
				moved = true
				delete(posizioni,k)
				posizioni[v] = 1
			}
		}
		if !moved{
			fmt.Println(i+1)
			panic("end")
			break
		}
	}
	toSet := true
	maxX,maxY,minX,minY := 0,0,0,0

	for k,_ := range posizioni{
		if toSet{
			toSet = false
			maxX,maxY,minX,minY =	k.x,k.y,k.x,k.y
		}

		if k.x > maxX{
			maxX = k.x
		}

		if k.x < minX{
			minX = k.x
		}

		if k.y > maxY{
			maxY = k.y
		}

		if k.y < minY{
			minY = k.y
		}
	}
	count := 0
	for d := minY; d <= maxY; d++{
		for i := minX; i <= maxX; i++{
			if _,ok := posizioni[pos{i,d}]; !ok{
				count++
				fmt.Print(".")
			}else{fmt.Print("#")}
		}
		fmt.Print("\n")
	}
	fmt.Println(len(posizioni))
	fmt.Println(count)
}

func checkNorthWest(k pos, posizioni map[pos]int)bool{
	for i := k.x-1; i <= k.x;i++{
		if _,ok := posizioni[pos{i,k.y-1}]; ok{
			return false
		}
	}
	return true
}
func checkNorthEast(k pos, posizioni map[pos]int)bool{
	for i := k.x; i <= k.x+1;i++{
		if _,ok := posizioni[pos{i,k.y-1}]; ok{
			return false
		}
	}
	return true
}
func checkSouthWest(k pos, posizioni map[pos]int)bool{
	for i := k.x-1; i <= k.x ;i++{
		if _,ok := posizioni[pos{i,k.y+1}]; ok{
			return false
		}
	}
	return true
}
func checkSouthEast(k pos, posizioni map[pos]int)bool{
	for i := k.x; i <= k.x+1;i++{
		if _,ok := posizioni[pos{i,k.y+1}]; ok{
			return false
		}
	}
	return true
}
func checkWestNorth(k pos, posizioni map[pos]int)bool{
	for i := k.y-1; i <= k.y;i++{
		if _,ok := posizioni[pos{k.x-1,i}]; ok{
			return false
		}
	}
	return true
}
func checkWestSouth(k pos, posizioni map[pos]int)bool{
	for i := k.y; i <= k.y+1;i++{
		if _,ok := posizioni[pos{k.x-1,i}]; ok{
			return false
		}
	}
	return true
}
func checkEastNorth(k pos, posizioni map[pos]int)bool{
	for i := k.y-1; i <= k.y;i++{
		if _,ok := posizioni[pos{k.x+1,i}]; ok{
			return false
		}
	}
	return true
}
func checkEastSouth(k pos, posizioni map[pos]int)bool{
	for i := k.y; i <= k.y+1;i++{
		if _,ok := posizioni[pos{k.x+1,i}]; ok{
			return false
		}
	}
	return true
}


func checkNorth(k pos, posizioni map[pos]int)bool{
	for i := k.x-1; i <= k.x+1;i++{
		if _,ok := posizioni[pos{i,k.y-1}]; ok{
			return false
		}
	}
	return true
}

func checkSouth(k pos,posizioni map[pos]int)bool{
	for i := k.x-1; i <= k.x+1;i++{
		if _,ok := posizioni[pos{i,k.y+1}]; ok{
			return false
		}
	}
	return true
}

func checkWest(k pos, posizioni map[pos]int)bool{
	for i := k.y-1; i <= k.y+1;i++{
		if _,ok := posizioni[pos{k.x-1,i}]; ok{
			return false
		}
	}
	return true
}

func checkEast(k pos, posizioni map[pos]int)bool{
	for i := k.y-1; i <= k.y+1;i++{
		if _,ok := posizioni[pos{k.x+1,i}]; ok{
			return false
		}
	}
	return true
}

func checkAround(k pos, posizioni map[pos]int)bool{
	if !checkNorth(k,posizioni){return false}

	if !checkSouth(k,posizioni){return false}

	if !checkEast(k,posizioni){return false}

	if !checkWest(k,posizioni){return false}

	return true
}