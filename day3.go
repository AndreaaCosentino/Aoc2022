package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan(){
		line := scanner.Text()
		tempMap := make(map[rune]int)
		rucksack := strings.Split(line,"")
		var duplicate rune
		
		// first half of the rucksack
		for i := 0; i < len(rucksack)/2; i++{
			tempMap[rune(rucksack[i][0])]++
		}

		// second half of the rucksack
		for i := len(rucksack)/2 ; i < len(rucksack); i++{
			if _,ok := tempMap[rune(rucksack[i][0])]; ok{
				duplicate = rune(rucksack[i][0])
				break
			}
		}

		var value int 
		if duplicate <= 'Z'{
			value = int(duplicate%'A') + 27
		}else{
			value = int(duplicate%'a') + 1
		}

		sum += value
	}
	fmt.Println(sum)
}