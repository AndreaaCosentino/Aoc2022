package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"unicode"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	sum :=0
	i := 1
	cont := make([]string,0)
	for scanner.Scan(){
		line := scanner.Text()
		if len(line) == 0{
			continue
		}
		cont = append(cont,line)
		scanner.Scan()
		line2 := scanner.Text()
		cont = append(cont,line2)
		c := compare(&line,&line2)
		if c == 1{
			sum += i
		}
		i++
	}
	stop := false
	for ;!stop;{
		stop = true
		for k:=0; k < len(cont)-1;k++{
			copy1 := cont[k]
			copy2 := cont[k+1] 
			s := compare(&copy1,&copy2)
			if(s == -1){
				cont[k],cont[k+1] = cont[k+1],cont[k]
				stop = false
			}
		}
	}
	d := 1
	for k,_ := range cont{
		if cont[k] == "[[2]]"{
			d *= (k+1)
			fmt.Println(k)
		}
		if cont[k] == "[[6]]"{
			fmt.Println(k)
			d *= (k+1)
			break
		}
	}
	fmt.Println(d)
}

/**
 * 1 SE  i1 < i2 
 * -1 SE  i1 > i2
 * 0 altrimenti
 * */
func compare(i1 *string, i2 *string) int{

	//Se c'è una virgola, la rimuovo. Non è importante
	if (*i1)[0:1] == ","{
		*i1 = (*i1)[1:]
	}
	if (*i2)[0:1] == ","{
		*i2 = (*i2)[1:]
	}


	if (*i1)[0:1] == "]"{
		*i1 = (*i1)[1:]
		if (*i2)[0:1] != "]"{
			return 1
		}
		(*i2) = (*i2)[1:]
		return 0
	}
	if (*i2)[0:1] == "]"{
		return -1
	}

	if unicode.IsNumber(rune((*i1)[0])) && unicode.IsNumber(rune((*i2)[0])){
		j:=0
		for ;unicode.IsNumber(rune((*i1)[j]));j++{}
		n, _ := strconv.Atoi((*i1)[0:j])
		*i1 = (*i1)[j:]
		for j = 0;unicode.IsNumber(rune((*i2)[j]));j++{}
		m, _ := strconv.Atoi((*i2)[0:j])
		*i2 = (*i2)[j:]
		if n < m {
			return 1
		}else if n > m{
			return -1
		}else{return 0}

	}

	if unicode.IsNumber(rune((*i1)[0])){
		j:=0
		for ;unicode.IsNumber(rune((*i1)[j]));j++{}
		*i1 = (*i1)[0:j] + "]"+ (*i1)[j:]
		*i1 = "[" + (*i1)
	}
	if unicode.IsNumber(rune((*i2)[0])){
		j:=0
		for ;unicode.IsNumber(rune((*i2)[j]));j++{}
		*i2 = (*i2)[0:j] + "]"+ (*i2)[j:]
		*i2 = "[" + (*i2)
	}


	for ;len((*i1)) != 0 && len((*i2)) != 0;{
		*i1 = (*i1)[1:]
		*i2 = (*i2)[1:]
		
		c := compare(i1,i2)
		if c != 0{
			return c
		}

		if (*i1)[0:1] == "]"{
			*i1 = (*i1)[1:]
				if (*i2)[0:1] != "]"{
					return 1
				}else{
					(*i2) = (*i2)[1:]
				}
			return 0
		}
		if (*i2)[0:1] == "]"{
			return -1
		}

	}

	if len((*i1)) > len((*i2)){
		return -1
	}else if len((*i1)) < len((*i2)){
		return 1
	}else{return 0}
}
