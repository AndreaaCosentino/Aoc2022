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

//	count := 0

	for scanner.Scan(){
		line := scanner.Text()
		
		lines = append(lines,line)
		max := "/"
		j := len(lines)-1
		dlines[j] = make([]int,100)
		// From left to right
		for i := 0; i < len(line); i++{
			if(lines[j][i:i+1] > max){
			//	count++
				max = lines[j][i:i+1]
				dlines[j][i] = i
				if(dlines[j][i] > h){
					h = dlines[j][i]
				}
			}
		}
		// From right to left
		max = "/"
		for i := len(line)-1; i >= 0; i--{
			if(lines[j][i:i+1] > max){
				max = lines[j][i:i+1]
				/*if dlines[j][i:i+1] != "-"{
					count++
				}*/
				if(dlines[j][i] == 0){
					dlines[j][i] = 1
				}
				dlines[j][i] *= (len(line)-1 - i)
				if(dlines[j][i] > h){
					h = dlines[j][i]
				}
			}
		}
		fmt.Println(dlines[j])
	}

	// From top to bottom
	max := "/"
	for i := 0; i < len(lines[0]); i++{
		max = "/"
			for j := 0; j < len(lines); j++{
				if(lines[j][i:i+1] > max){
					max = lines[j][i:i+1]
				/*	if dlines[j][i:i+1] != "-"{
						count++
					}*/
					if(dlines[j][i] == 0){
						dlines[j][i] = 1
					}
					dlines[j][i] *= j
					if(dlines[j][i] > h){
						h = dlines[j][i]
					}
				}
			}
	} 

	// From bottom to up
	for i := len(lines[0])-1; i >= 0; i--{
		max = "/"
			for j := len(lines)-1; j >= 0; j--{
				if(lines[j][i:i+1] > max){
					max = lines[j][i:i+1]
					/*if dlines[j][i:i+1] != "-"{
						count++
					}*/
					if(dlines[j][i] == 0){
					dlines[j][i] = 1
				}
					dlines[j][i] *= (len(lines)-1 - j)
					if(dlines[j][i] > h){
						h = dlines[j][i]
					}
				}
			}
	} 
	fmt.Println(h)
}