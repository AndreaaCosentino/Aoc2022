package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)



func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	cycle := 1
	reg := 1
	sum := 0
	v := 0
	d := 0
	n := 0
	monitor := make([]string,6)
	pos := 0
	for i := 0; i < 6; i++{
		monitor[i] = ""
	}
	for scanner.Scan(){
		line := scanner.Text()
		instruction := strings.Split(line," ")
		if(instruction[0] == "noop"){
			if(pos >= reg-1 && pos <= reg+1){
				monitor[n] +="@"
					fmt.Println(pos," ",reg)
			}else{
				monitor[n] += ","
				fmt.Println(",", pos," ",reg," ",monitor[n])
			}
			if cycle%40 == 0 {
				sum += cycle*reg
				n++
				n = n%6
				pos = 0
			}
			cycle++
			pos++
			d = -1
		}else{
			q,_ := strconv.Atoi(instruction[1])
			v = q;
			d = 2;
		}
		for ; d > 0;{
			d--
			if(pos >= reg-1 && pos <= reg+1){
				monitor[n] += "#"
			}else{monitor[n] += "."}
			if(d == 0){
				reg += v 
				v = 0
			}
			if cycle%40 == 0{
				sum += cycle*reg
				n++
				n = n%6
				pos = 0
			}
			cycle++
			pos++
		}
	}
	for k,_ := range monitor{
		fmt.Println(monitor[k])
	}
}
