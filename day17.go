package main

import(
	"fmt"
	"bufio"
	"os"
	"time"
)

type position struct{
	x,y int 
}

type state struct{
	bitMask int
	j int 
	i int
}

type situation struct{
	i int 
	height int
}

func main(){
	deltaTime := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	rocksPositions := make(map[position]bool)
	var myInput string

	mapStates := make(map[state]situation)

	scanner.Scan()
	myInput = scanner.Text()

	//   1 <= X <= 7
	j := 0
	height := 0
	q := 1000000000000
	for i := 1; i < q; i++{
		rocks := createRock(i,height)

		for k := 0;;k++{

			if k%2 == 0{
				if j == len(myInput){
					j = 0
				}
				if myInput[j] == '>'{
					if checkRight(rocksPositions,rocks){
						sum(rocks,1,0)
					}
				}else{
					if checkLeft(rocksPositions,rocks){
						sum(rocks,-1,0)
					}
				}
				j++
			}else{
				if checkBottom(rocksPositions,rocks){
					sum(rocks,0,-1)
				}else{break}
			}
		}
		for _,v := range rocks{
			rocksPositions[v] = true
		}	

		tempHeight := getTop(rocks)

		if tempHeight > height{
			height = tempHeight	
		}
		tempState := createState(height,j,i,rocksPositions)

		if _,ok := mapStates[tempState]; !ok{
			mapStates[tempState] = situation{i,height}
		}else{
			if (q-mapStates[tempState].i) % ( i - mapStates[tempState].i) == 0{
				x := (q  - mapStates[tempState].i) /( i - mapStates[tempState].i)
				d := (x-1) * (height - mapStates[tempState].height)
				height += d
				break
			}
			continue
		}
	}
	fmt.Println(height)
	fmt.Println("Tempo: ",time.Since(deltaTime))
}

func createState(height int,j int, i int, rocksPositions map[position]bool) state{
	var tempBitMask int 

	for x := 1; x < 8; x++{
		if _, ok := rocksPositions[position{x,height}]; ok{
			tempBitMask += 1 << x
		}
	}

	return state{tempBitMask,j,i%5}
}

func sum(rocks []position,x int, y int){

	for k,_ := range rocks{
		rocks[k].x += x 
		rocks[k].y += y
	}
}

func getTop(rocks []position) int {
	var y int 
	y = rocks[0].y

	for _,v := range rocks{
		if v.y > y {
			y = v.y
		}
	}

	return y
}

// VERO SE NON COLPISCE NULLA
func checkBottom(rocksPositions map[position]bool, rocks []position )bool{
	for _,v := range rocks{
		v.y--
		if _,ok := rocksPositions[v]; ok || v.y < 1{
			return false
		}
	}
	return true
}

func checkLeft(rocksPositions map[position]bool, rocks []position)bool{
	for _,v := range rocks{
		v.x--
		if _,ok := rocksPositions[v]; ok || v.x < 1{
			return false
		}
	}
	return true
}

func checkRight(rocksPositions map[position]bool, rocks []position)bool{
	for _,v := range rocks{
		v.x++
		if _,ok := rocksPositions[v]; ok || v.x > 7{
			return false
		}
	}
	return true
}

func createRock(i int, height int) []position{

	rocks := make([]position,0)	
	i--
	switch(i%5){

			case 0:
					rocks = append(rocks,position{3,height+4})
					rocks = append(rocks,position{4,height+4})
					rocks = append(rocks,position{5,height+4})
					rocks = append(rocks,position{6,height+4})
			case 1:
					rocks = append(rocks,position{4,height+6})
					rocks = append(rocks,position{3,height+5})
					rocks = append(rocks,position{4,height+5})
					rocks = append(rocks,position{5,height+5})
					rocks = append(rocks,position{4,height+4})
			case 2: 
					rocks = append(rocks,position{3,height+4})
					rocks = append(rocks,position{4,height+4})
					rocks = append(rocks,position{5,height+4})
					rocks = append(rocks,position{5,height+5})
					rocks = append(rocks,position{5,height+6})
			case 3:
					rocks = append(rocks,position{3,height+4})
					rocks = append(rocks,position{3,height+5})
					rocks = append(rocks,position{3,height+6})
					rocks = append(rocks,position{3,height+7})
		    case 4:
					rocks = append(rocks,position{3,height+4})
					rocks = append(rocks,position{3,height+5})
					rocks = append(rocks,position{4,height+4})
					rocks = append(rocks,position{4,height+5})
		}
	return rocks
}
func UNUSED(x ...interface{}) {}