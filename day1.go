package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main(){

	s := bufio.NewScanner(os.Stdin)

	count := 0
	max := 0
	max2 := 0
	max3 := 0

	for s.Scan(){
		line := s.Text()
		if line == ""{
			if count > max{
				max3 = max2
				max2 = max
				max = count
			}else if count > max2{
				max3 = max2
				max2 = count
			}else if count > max3{
				max3 =  countcd doc
			}
			count = 0
		}else{
			temp,_ := strconv.Atoi(line)
			count += temp
		}
	}

	fmt.Println(max+max2+max3)
}
