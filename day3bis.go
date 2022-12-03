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
		
		// Go through the first elf's rucksack
		for i := 0; i < len(rucksack); i++{
				tempMap[rune(rucksack[i][0])]++
		}


		tempMap2 := make(map[rune]int)
		// Go through the other two elves' rucksacks
		for i := 0; i < 2; i++{

			scanner.Scan()
			line = scanner.Text()

			rucksackbis := strings.Split(line,"")

			// For the second elf
			if i == 0{
				for j := 0; j < len(rucksackbis); j++{
					if _,ok := tempMap[rune(rucksackbis[j][0])]; ok{
						tempMap2[rune(rucksackbis[j][0])]++
					}	
				}
			}
			// For the third elf
			if i == 1{
				for j := 0; j < len(rucksackbis); j++{
					if _,ok := tempMap2[rune(rucksackbis[j][0])]; ok{
						duplicate = rune(rucksackbis[j][0])
						break
					}	
				}
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