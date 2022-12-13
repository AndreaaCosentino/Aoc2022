package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	lines := make([]string,0)
	dlines := make([][]int,100)
	h:= 0

	for scanner.Scan(){
		line := scanner.Text()
		
		lines = append(lines,line)
		dlines[len(lines)-1] = make([]int,100)
		for i:= 0; i < 100; i++{
			dlines[len(lines)-1][i] = 1 
		}
	}

	HEIGHT := len(lines)
	WIDTH := len(lines[0])
	for i := 0; i < HEIGHT; i++{
		for j := 0; j < WIDTH; j++{


			//RIGHT
			r:= j+1
			for r = j+1; r < WIDTH-1; r++{
				if lines[i][j] <= lines[i][r]{
					break
				}
			} 
			fmt.Println(r-j)
			dlines[i][j] *= r-j
			//LEFT
			for r = j-1; r > 0; r--{
				if lines[i][j] <= lines[i][r]{
					break
				}
			}
			dlines[i][j] *= j-r
			fmt.Println(j-r)
			//UP
			for r = i-1; r > 0; r--{
				if lines[i][j] <= lines[r][j]{
					break
				}
			} 
			dlines[i][j] *= i-r
			fmt.Println(i-r)
			//BOTTOM
			for r = i+1; r < HEIGHT-1; r++{
				if lines[i][j] <= lines[r][j]{
					break
				}
			}
			dlines[i][j] *= r-i
			fmt.Println(r-i)

			if(dlines[i][j] > h){
				h = dlines[i][j]
			}
		}		
	}
	fmt.Println(h)
}