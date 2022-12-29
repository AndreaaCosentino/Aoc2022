package main 

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type num struct{
	n int 
	pos int
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	numbers := make([]num,0)
	copy_numbers := make([]int,0)
	j:=0
	itsZero := 0
	for scanner.Scan(){
		line := scanner.Text()

		temp,_ := strconv.Atoi(line)

		if temp == 0{
			itsZero = j
		}
		numbers = append(numbers,num{temp*811589153,j})
		copy_numbers = append(copy_numbers,temp*811589153)
		j++
	}
	for j := 0; j < 10;j++{
		for i := 0; i < len(copy_numbers); i++{
			temp := copy_numbers[i]
			pos := searchI(numbers,num{temp,i})
			if temp == 0{continue}
			posFin := temp+pos
			if posFin <= 0{
				posFin =  mod(posFin,len(copy_numbers)-1)
			}else if  (posFin >= len(numbers)-1){
				posFin =  mod(posFin,len(copy_numbers)-1)
			}

			if pos < posFin{
				for i:= pos; i < posFin; i++{
					numbers[i],numbers[i+1] = numbers[i+1],numbers[i]
				}
			}else{
			for i:= pos; i > posFin; i--{
						numbers[i],numbers[i-1] = numbers[i-1],numbers[i]
				}
			}
			//fmt.Println(numbers)
		}
	}
	pos := searchI(numbers,num{0,itsZero})
	one := numbers[mod(pos+1000,len(copy_numbers))].n
	two := numbers[mod(pos+2000,len(copy_numbers))].n
	three := numbers[mod(pos+3000,len(copy_numbers))].n
	fmt.Println(one+two+three)
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func searchI(numbers []num, n num ) int {
	for i := 0; i < len(numbers); i++{
		if numbers[i] == n{
			return i
		}
	}
	panic("WHAT THE HELL")
}