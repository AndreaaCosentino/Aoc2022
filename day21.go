package main 

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type monkey struct{
	resAvailable bool
	res int
	operands []string
	operation string
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	mappaScimmie := make(map[string]monkey)
	for scanner.Scan(){
		line := scanner.Text()

		myInput := strings.Split(line," ")
		myInput[0] = myInput[0][:len(myInput[0])-1]

		if len(myInput) == 2{
			var temp monkey 
			temp.resAvailable = true 
			temp.res,_ = strconv.Atoi(myInput[1])
			mappaScimmie[myInput[0]] = temp
		}else{
			var temp monkey 
			temp.resAvailable = false 
			temp.operands = make([]string,2)
			temp.operands[0] = myInput[1]
			temp.operands[1] = myInput[3]
			temp.operation = myInput[2]

			mappaScimmie[myInput[0]] = temp
		}
	}

	for i:= 0; i < 10000; i++{
		for k,v := range mappaScimmie{
			if v.resAvailable{
				continue
			}
			if k == "root"{continue}
			op1 := mappaScimmie[v.operands[0]]
			op2 := mappaScimmie[v.operands[1]]
			if v.operands[0] == "humn" || v.operands[1] == "humn"{continue}
			if op1.resAvailable && op2.resAvailable{
					switch(v.operation){
					case "+": v.res = op1.res + op2.res
					case "*": v.res = op1.res * op2.res
					case "/": v.res = op1.res / op2.res
					case "-": v.res = op1.res - op2.res
					}
					v.resAvailable = true
					mappaScimmie[k] = v
			}
		}
	}
	for i:= 3373767893000; ;i++{
			tempMap := make(map[string]int)
		for ; isTrue(tempMap,mappaScimmie) ;{
			for k,v := range mappaScimmie{
			if v.resAvailable{
				continue
			}
			if k == "root"{continue}
			op1 := mappaScimmie[v.operands[0]]
			op2 := mappaScimmie[v.operands[1]]
			if v.operands[0] == "humn"{
				op1.res = i
			}else if v.operands[1] == "humn"{
				op2.res = i
			}
			v1,ok := tempMap[v.operands[0]]
			v2,ok1 := tempMap[v.operands[1]]

			if (ok || op1.resAvailable) && (op2.resAvailable || ok1){
				if ok {op1.res = v1}
				if ok1{op2.res = v2}
					switch(v.operation){
					case "+": v.res = op1.res + op2.res
					case "*": v.res = op1.res * op2.res
					case "/": v.res = op1.res / op2.res
					case "-": v.res = op1.res - op2.res
					}
					tempMap[k] = v.res
				}
			}
		}
		v,ok :=tempMap[mappaScimmie["root"].operands[0]]
	    v1,ok1 := tempMap[mappaScimmie["root"].operands[1]]

	    valore := 0
	    valore1 := 0
	    if ok{valore = v}else{valore = mappaScimmie[mappaScimmie["root"].operands[0]].res}
	    if ok1{valore1 = v1}else{valore1 = mappaScimmie[mappaScimmie["root"].operands[1]].res}
	    fmt.Println(valore-valore1)
	    if valore == valore1{
	    	fmt.Println(i)
	    	break
	    }
	   fmt.Println(i)
	}
	//fmt.Println(mappaScimmie["root"].res)
}

func isTrue(tempMap map[string]int, mappaScimmie map[string]monkey)bool{
	_,ok :=tempMap[mappaScimmie["root"].operands[0]]
	_,ok1 := tempMap[mappaScimmie["root"].operands[1]]
	if (mappaScimmie[mappaScimmie["root"].operands[0]].resAvailable || ok) && (mappaScimmie[mappaScimmie["root"].operands[1]].resAvailable || ok1){
		return false
	}else{return true}
}

func createOperand(mappaScimmie map[string]monkey, monkey string) string{
	if monkey == "humn"{
		return "x"
	}

	if mappaScimmie[monkey].resAvailable{
		return strconv.Itoa(mappaScimmie[monkey].res)
	}
	return "(" + createOperand(mappaScimmie,mappaScimmie[monkey].operands[0]) +" "+ mappaScimmie[monkey].operation  +" "+ createOperand(mappaScimmie,mappaScimmie[monkey].operands[1]) + ")"
}