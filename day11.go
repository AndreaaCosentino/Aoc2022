/*DOESNT WORK ON GIVEN INPUT IF COMMAS NOT REMOVED MANUALLY.*/
package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Monkey struct{
	n int
	list []int
	test int
	operation string 
	amount int
	to [2]int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	monkeys := make([]Monkey,8)

	for scanner.Scan(){
		line := scanner.Text()
		d := strings.Split(line," ")
		if len(d) != 1{
			d[1] = d[1][0:1]
			i,_ := strconv.Atoi(d[1])
			var temp Monkey
			temp.n = i
			temp.list = make([]int,0)
			scanner.Scan()
			line = scanner.Text()
			d = strings.Split(line," ")
			for j := 4; j < len(d); j++{
				k,_ := strconv.Atoi(d[j])
				temp.list =  append(temp.list,k)
			}
			scanner.Scan()
			line = scanner.Text()
			d = strings.Split(line," ")
			temp.operation = d[6]
			temp.amount, _ = strconv.Atoi(d[7])

			scanner.Scan()
			line = scanner.Text()
			d = strings.Split(line," ")
			temp.test, _ = strconv.Atoi(d[5])
			scanner.Scan()
			line = scanner.Text()
			d = strings.Split(line," ")
			temp1, _ := strconv.Atoi(d[9])
			temp.to[0] = temp1
			scanner.Scan()
			line = scanner.Text()
			d = strings.Split(line," ")
			temp1,_ = strconv.Atoi(d[9])
			temp.to[1] = temp1
			monkeys[i] = temp
		}
	}
	s := make([]int,8)

	for i := 0; i < 8 ; i++ {
		s[i] = 0
	}
	for i := 0; i < 10000; i++{
		for j:=0; j < len(monkeys); j++{
			m := len(monkeys[j].list)
			s[j] += m
			for d := 0; d < m ; d++{
				temp := monkeys[j].list[0]
				monkeys[j].list = monkeys[j].list[1:]
				if monkeys[j].operation == "*"{
					if monkeys[j].amount == 0{
						temp = temp*temp
					}else{
						temp = temp*monkeys[j].amount
					}
				}else{
					if monkeys[j].amount == 0{
						temp = temp+temp
					}else{
						temp = temp+monkeys[j].amount
					}
				}
				//temp = temp/3
				temp = temp%9699690
				if temp % monkeys[j].test == 0{
					monkeys[monkeys[j].to[0]].list =append(monkeys[monkeys[j].to[0]].list,temp)
				}else{monkeys[monkeys[j].to[1]].list =append(monkeys[monkeys[j].to[1]].list,temp)}
			}
		}
	}
	for i := 0; i < len(monkeys); i++{
		fmt.Println(s[i])
	}
	fmt.Println(monkeys)
}