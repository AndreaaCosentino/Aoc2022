package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type coordinates struct{
	x,y int
}

type stato struct{
	position coordinates
	facing int // 0 right,1 down, 2 left, 3 up
	currentInstruction int
}

type faccia1 struct{
	ul,ur,dl,dr coordinates
}

var maxX,maxY int
var faccia map[coordinates]int
var facce []faccia1

func main(){
	maxX,maxY = 0,0
	scanner := bufio.NewScanner(os.Stdin)
	path := ""
	mappa := make(map[coordinates]string)
	faccia = make(map[coordinates]int)

	facce = make([]faccia1,0)
	facce = append(facce,faccia1{coordinates{51,1},coordinates{100,1},coordinates{51,50},coordinates{100,50}})
	facce = append(facce,faccia1{coordinates{101,1},coordinates{150,1},coordinates{101,50},coordinates{150,50}})
	facce = append(facce,faccia1{coordinates{51,51},coordinates{100,51},coordinates{51,100},coordinates{100,100}})
	facce = append(facce,faccia1{coordinates{51,101},coordinates{100,101},coordinates{51,150},coordinates{100,150}})
	facce = append(facce,faccia1{coordinates{1,101},coordinates{50,101},coordinates{1,150},coordinates{50,150}})
	facce = append(facce,faccia1{coordinates{1,151},coordinates{50,151},coordinates{1,200},coordinates{50,200}})
	// IF EMPTY, NO ENTRY IN THE MAP.
	// IF WALKABLE, ENTRY = .
	// IF NOT WALKABLE, ENTRY = #
	row := 1
	startingPoint := false
	var startPosition coordinates
	for scanner.Scan(){
		line := scanner.Text()
		if len(line) != 0{
			for k,v := range line {
				if v == '.'{
					if !startingPoint{
						startPosition = coordinates{k+1,row}
						mappa[startPosition] = "S"
						startingPoint = true
					}else{mappa[coordinates{k+1,row}] = "."}
				}else if v == '#'{
					mappa[coordinates{k+1,row}] = "#"
				}else{continue}
				if (k+1) > maxX{
					maxX = k+1 
				}

				if row <= 50{
					if k+1 <= 100{
						faccia[coordinates{k+1,row}] = 1
					}else{faccia[coordinates{k+1,row}] = 2}
				}else if row <= 100{
					faccia[coordinates{k+1,row}] = 3
				}else if row <= 150{
					if k+1 <= 50{
						faccia[coordinates{k+1,row}] = 5
					}else{faccia[coordinates{k+1,row}] = 4}
				}else{
					faccia[coordinates{k+1,row}] = 6
				}
			}
		}else{
			scanner.Scan()
			path = scanner.Text()
			break
		}
		row++
	}
	maxY = row
	instructions := make([]string,0)
	mypath := path
	for ; len(mypath) > 0;{
		nextR := strings.Index(mypath,"R")
		nextL := strings.Index(mypath,"L")
		closer := 0
		if nextR < nextL && nextR != -1{
			closer = nextR
		}else{closer = nextL}
		if closer != -1{
			instructions = append(instructions,mypath[:closer])
			instructions = append(instructions,mypath[closer:closer+1])
			mypath = mypath[closer+1:]
		}else{
			instructions = append(instructions,mypath[0:])
			break
		}
	}
	statoCorrente := stato{startPosition,0,0}

	for ;statoCorrente.currentInstruction < len(instructions); statoCorrente.currentInstruction++{
		n,err := strconv.Atoi(instructions[statoCorrente.currentInstruction])
		if err == nil{
			movx := 0
			movy := 0
			switch(statoCorrente.facing){
			case 0: movx = 1 
			case 1: movy = 1
		    case 2: movx = -1
		    case 3: movy = -1
			}
			for i := 0; i < n; i++{
				if !move(statoCorrente.position,mappa,movx,movy,&statoCorrente){
					break
				}
				movx = 0
				movy = 0
				switch(statoCorrente.facing){
				case 0: movx = 1 
				case 1: movy = 1
		    	case 2: movx = -1
		   		 case 3: movy = -1
				}
			}
		}else{
			statoCorrente.facing = calcolaDirezione(statoCorrente.facing,instructions[statoCorrente.currentInstruction])
		}
	}
	fmt.Println(statoCorrente)
	fmt.Println(statoCorrente.position.y*1000 + statoCorrente.position.x*4 + statoCorrente.facing)
}

func move(posizione coordinates, mappa map[coordinates]string, movX int, movY int, statoCorrente *stato)bool{
	nuovaPosizione := coordinates{posizione.x+movX,posizione.y+movY}

	if v,ok := mappa[nuovaPosizione]; ok{
		if v == "#"{return false}
		statoCorrente.position = nuovaPosizione
		return true
	}else{
		tempPos,tempFace := destinazione(statoCorrente.position, statoCorrente.facing)
		if mappa[tempPos] == "#"{return false}
		statoCorrente.position = tempPos
		statoCorrente.facing = tempFace
		return true
	}
	panic("NOT FOUND")
}


func calcolaDirezione(dirCorrente int, RoL string) int {
	if RoL == "L"{
		dirCorrente--
		if dirCorrente == -1{
			dirCorrente = 3
		}
	}else{
		dirCorrente++
		if dirCorrente == 4{
			dirCorrente = 0
		}
	}
	return dirCorrente
}

func destinazione(currPosition coordinates, facing int) (coordinates,int){
	var newPosition coordinates
	newFacing := -1
	if faccia[currPosition] == 1{
			if facing == 2{
				newFacing = 0
				newPosition.x = facce[4].ul.x
				newPosition.y = (currPosition.y-facce[0].ul.y)+ facce[4].dl.y
			}else if facing == 3{
				newFacing = 0
				newPosition.y = (currPosition.y-facce[0].ul.y)+ facce[5].ul.y
				newPosition.x = facce[5].ul.x
			}
	}

	if faccia[currPosition] == 2{
		if facing == 0 {
				newFacing = 2
				newPosition.x = facce[3].ur.x
				newPosition.y =	(currPosition.y-facce[1].ul.y)+ facce[3].ul.y
			}else if facing == 3{
				newFacing = 3
				newPosition.y = facce[5].dl.y
				newPosition.x =	(currPosition.x-facce[1].ul.x)+ facce[5].ul.x
			}else if facing == 1{
				newFacing = 2
				newPosition.x = facce[2].ur.x
				newPosition.y =	(currPosition.x-facce[1].ul.x)+ facce[2].ur.y
			}
	}

	if faccia[currPosition] == 3{
		if facing == 0{
			newFacing = 3
			newPosition.y = facce[1].dl.y
			newPosition.x =	(currPosition.y-facce[2].ul.y) + facce[1].ul.x
		}else if facing == 2{
			newFacing = 1
			newPosition.y = facce[4].ul.y
			newPosition.x =	(currPosition.y-facce[2].ul.y) + facce[4].ul.x
		}
	}

	if faccia[currPosition] == 4{
		if facing == 0{
			newFacing = 2
			newPosition.x = facce[1].ur.x
			newPosition.y =	(currPosition.y-facce[3].ul.y) + facce[1].ur.y
		}else if facing == 1{
			newFacing = 2
			newPosition.x = facce[5].ur.x
			newPosition.y =	(currPosition.y-facce[3].ul.y) + facce[5].ur.y
		}
	}

	if faccia[currPosition] == 5{
		if facing == 3{
			newFacing = 0
			newPosition.x = facce[2].ul.x
			newPosition.y =	(currPosition.x-facce[4].ul.x) + facce[2].ur.y
		}else if facing == 2 {
			newFacing = 0
			newPosition.x = facce[0].ul.x
			newPosition.y =	(currPosition.x-facce[4].ul.x) + facce[0].ur.y
		}
	}

	if faccia[currPosition] == 6{
		if facing == 2{
			newFacing = 1
			newPosition.y = facce[0].ul.y
			newPosition.x = (currPosition.y-facce[5].ul.y) + facce[0].ul.x
		}else if facing == 1{
			newFacing = 1
			newPosition.y = facce[1].ul.y
			newPosition.x = (currPosition.x-facce[5].ul.x) + facce[1].ul.x
		}else if facing == 0{
			newFacing = 3
			newPosition.y = facce[3].dl.y 
			newPosition.x = (currPosition.y-facce[5].ul.y) + facce[3].ul.x
		}
	}

	if newFacing == -1{
		fmt.Println("ERROR at position : ",currPosition," and facing: ",facing)
		fmt.Println(faccia[currPosition])
		panic("SHUT DOWN")
	}

	return newPosition,newFacing
}