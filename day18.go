package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type position struct{
	x,y,z int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	positions := make([]position,0)

	surface := 0

	for scanner.Scan(){
		line := scanner.Text()

		input := strings.Split(line,",")

		x,_ := strconv.Atoi(input[0])
		y,_ := strconv.Atoi(input[1])
		z,_ := strconv.Atoi(input[2])

		positions = append(positions,position{x,y,z})

		surface += 6
		for _,v := range positions{

			if v == (position{x,y,z}){
				continue
			}

			if  (math.Abs(float64(x-v.x))) + math.Abs(float64(y-v.y)) + math.Abs(float64(z-v.z)) == 1{
				surface -= 2
			}
		}
	}
	fmt.Println(surface)

}
