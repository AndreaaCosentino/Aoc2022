package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type Position struct{
	x int
	y int 
}

const KNOTS = 10

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	points := make([]Position,KNOTS)

	for i := 0; i < KNOTS ; i++{
		points[i].x = 0
		points[i].y = 0
	}
	mapPos := make(map[Position]int)
	mapPos[points[KNOTS-1]] = 1

	for scanner.Scan(){
		line := scanner.Text()
		act := strings.Split(line," ")
		move, _ := strconv.Atoi(act[1])
		for j := 0; j < move ; j++{
			if act[0] == "R"{
				points[0].x++
			}else if act[0] == "L"{
				points[0].x--
			}else if act[0] == "U"{
				points[0].y++
			}else if act[0] == "D"{
				points[0].y--
			}
			for i := 1 ; i < KNOTS; i++{
				movement(points,mapPos,i)
			}
		}
	}
	count := 0
	for _,_ = range mapPos{
		count++
	}
	fmt.Println(count)
}

func movement(points []Position, mapPos map[Position]int, i int){
	x1 := points[i-1].x 
	x2 := points[i].x
	y1 := points[i-1].y 
	y2 := points[i].y 
	if(!(math.Abs(float64(x1-x2)) <= 1 && math.Abs(float64(y1-y2)) <= 1)){
		dirx:= 0
		diry := 0
		if(x1-x2 != 0){
			dirx = (x1-x2)/int(math.Abs(float64(x1-x2)))
		}
		if(y1-y2 != 0){
			diry = (y1-y2)/int(math.Abs(float64(y1-y2)))
		}
		points[i].x += dirx
		points[i].y += diry
	}
	if(i == KNOTS-1){
		mapPos[points[KNOTS-1]] = 1
	}
}