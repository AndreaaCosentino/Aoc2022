package main

import(
	"fmt"
	"bufio"
	"os"
	//"strings"
	//"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	c := 0
	for scanner.Scan(){
		line := scanner.Text()

		var v string
		for i := 0; i < len(line); i++{
			v += line[i:i+1]
			c++
			for j := 0; j < c-1; j++{
				if v[j:j+1] == v[c-1:c]{
					v = v[j+1:]
					c -= j+1
					break
				}
			}
			if(c == 14){
				fmt.Println(i)
				break
			}
		}
	}
}