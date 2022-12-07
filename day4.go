package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan(){
		line := scanner.Text()
		pair := strings.Split(line,",")

		firstElf := strings.Split(pair[0],"-")
		secondElf := strings.Split(pair[1],"-")

		f1,_ := strconv.Atoi(firstElf[0])
		f2,_ := strconv.Atoi(firstElf[1])
		s1,_ := strconv.Atoi(secondElf[0]) 
		s2,_ := strconv.Atoi(secondElf[1])

		if(  (f1 <= s1 && f2 >= s1) || (s1 <= f1 && s2 >= f1))  {
			count++
		}
	}

	fmt.Println(count)
}