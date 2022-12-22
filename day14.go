package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type point struct{
	x int 
	y int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	matrix := make([][]int,0)

	j := 0
	for scanner.Scan(){
		line := scanner.Text()
		matrix = append(matrix,make([]int,0))
		line = strings.Replace(line,"-","",-1)
		line = strings.Replace(line," ","",-1)
		line = strings.Replace(line,">",",",-1)
		inputs := strings.Split(line,",")

		rocksPosition := make([]point,0)

		for i := 0;i < len(inputs);{
			x1,_ := strconv.Atoi(inputs[i])
			y1,_ := strconv.Atoi(inputs[(i+1)])
			newPoint := point{
				x : x1, 
				y : y1,
			}
			rocksPosition = append(rocksPosition,newPoint)
			if(i != 0){
				fmt.Println(rocksPosition)
				formerPoint := rocksPosition[len(rocksPosition)-1]
				if formerPoint.x != newPoint.x{
					if formerPoint.x > newPoint.x{
						for ;newPoint.x < formerPoint.x-1;{
							newPoint.x++
							rocksPosition = append(rocksPosition,newPoint)
						}
					}else{
						for ;newPoint.x-1 >  formerPoint.x;{
							newPoint.x--
							rocksPosition = append(rocksPosition,newPoint)
						}
					}
				}else{
					if formerPoint.y > newPoint.y{
						for ;newPoint.y -1 > formerPoint.y;{
							newPoint.y++
							rocksPosition = append(rocksPosition,newPoint)
						}
					}else{
						for ;newPoint.y -1 >formerPoint.y;{
							newPoint.y--
							rocksPosition = append(rocksPosition,newPoint)
						}
					}
				}
			}
			i += 2
		}
		fmt.Println(rocksPosition)
		j++
	}
	
}
