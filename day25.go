package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan(){
		line := scanner.Text()
		number := SNAFUtoDecimal(line)
		fmt.Println(line,number)
		sum += number
	}
	fmt.Println(decimalToSNAFU(sum),sum)
}

func SNAFUtoDecimal(SNAFU string) int{
	currentPow := 1
	res := 0

	for i := len(SNAFU)-1; i >=0;i--{
		temp := 0
		if SNAFU[i:i+1] == "="{
			temp = -2
		}else if SNAFU[i:i+1] == "-"{
			temp = -1
		}else{
			temp,_ = strconv.Atoi(SNAFU[i:i+1])
		}
		res += temp*currentPow
		currentPow *= 5
	}
	return res
}

func decimalToSNAFU(n int)string{
	res := ""
	for n != 0{
		temp := n/5
		digit := n-temp*5
		n = temp
		fmt.Println(digit)
		if digit == 3{
			res = "="+res
			n += 1
		}else if digit == 4{
			res =  "-" + res
			n += 1
		}else if digit == 2{
			res = "2" + res
		}else if digit == 1{
			res = "1"+ res
		}else{res =  "0"+res}
	}
	return res
}