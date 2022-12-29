package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
	//"math"
)

type position struct{
	x,y,z int
}

var mapPosition map[position]bool
var visited map[position]int

func main(){
	start := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	hX,lX,hY,lY,hZ,lZ := 0,0,0,0,0,0

	mapPosition = make(map[position]bool)
	for scanner.Scan(){
		line := scanner.Text()

		input := strings.Split(line,",")

		x,_ := strconv.Atoi(input[0])
		y,_ := strconv.Atoi(input[1])
		z,_ := strconv.Atoi(input[2])

		mapPosition[position{x,y,z}] = true
	}


	for k,_ := range  mapPosition{
		if k.x > hX {
			hX = k.x
		}else if k.x < lX{lX = k.x}

		if k.y > hY {
			hY = k.y
		}else if k.y < lY{lY = k.y}

		if k.z > hZ {
			hZ = k.z
		}else if k.z < lZ{lZ = k.z}
	}
	faces := 0
	for k,_ := range  mapPosition{
		d := 0
		visited = make(map[position]int)
		d+= find(position{k.x+1,k.y,k.z},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		visited = make(map[position]int)
		d+= find(position{k.x-1,k.y,k.z},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		visited = make(map[position]int)
		d+= find(position{k.x,k.y+1,k.z},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		visited = make(map[position]int)
		d+= find(position{k.x,k.y-1,k.z},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		visited = make(map[position]int)
		d+= find(position{k.x,k.y,k.z+1},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		visited = make(map[position]int)
		d+= find(position{k.x,k.y,k.z-1},lX-5,lY-5,lZ-5,hX+5,hY+5,hZ+5)
		faces += d
	}

	fmt.Println(faces)
	fmt.Println(time.Since(start))
}
func find(pos position,lX int,lY int,lZ int,hX int, hY int,hZ int) int{
	visited[pos] = 1

	if _,ok := mapPosition[pos]; ok{
		return 0
	}

	if pos.x < lX{
		return 1
	}else if pos.x > hX {return 1}

	if pos.y < lY{
		return 1
	}else if pos.y > hY{return 1}

	if pos.z < lZ{
		return 1
	}else if pos.z > hZ{return 1}

	d:= 0


	if _,ok := visited[position{pos.x-1,pos.y,pos.z}]; !ok{
		d += find(position{pos.x-1,pos.y,pos.z},lX,lY,lZ,hX,hY,hZ)
	}
	if _,ok := visited[position{pos.x,pos.y-1,pos.z}]; !ok {
		d += find(position{pos.x,pos.y-1,pos.z},lX,lY,lZ,hX,hY,hZ)
	}
	if _,ok := visited[position{pos.x,pos.y,pos.z-1}]; !ok{
		d += find(position{pos.x,pos.y,pos.z-1},lX,lY,lZ,hX,hY,hZ)
	}
	if _,ok := visited[position{pos.x+1,pos.y,pos.z}]; !ok{
		d += find(position{pos.x+1,pos.y,pos.z},lX,lY,lZ,hX,hY,hZ)
	}
	if _,ok := visited[position{pos.x,pos.y+1,pos.z}]; !ok{
		d += find(position{pos.x,pos.y+1,pos.z},lX,lY,lZ,hX,hY,hZ)
	}
	if _,ok := visited[position{pos.x,pos.y,pos.z+1}]; !ok{
		d += find(position{pos.x,pos.y,pos.z+1},lX,lY,lZ,hX,hY,hZ)
	}

	if d == 0 {return 0}
	return 1
}