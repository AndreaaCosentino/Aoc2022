package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanLines)

	cranes := make([]string,9)
	cranes[0] = "JHGMZNTF"
	cranes[1] = "VWJ"
	cranes[2] = "GVLJBTH"
	cranes[3] = "BPJNCDVL"
	cranes[4] = "FWSMPRG"
	cranes[5] = "GHCFBNVM"
	cranes[6] = "DHGMR"
	cranes[7] = "HNMVZD"
	cranes[8] = "GNFH"

	for scanner.Scan(){	
		line := scanner.Text()
		move := strings.Split(line," ")
		
		quantity,_ := strconv.Atoi(move[1])
		from,_ := strconv.Atoi(move[3])
		to,_ := strconv.Atoi(move[5])
		/*for i := 0; i < quantity; i++{
			temp := len(cranes[from-1])
			cranes[to-1] = cranes[to-1]+cranes[from-1][temp-1:temp]
			cranes[from-1] = cranes[from-1][:temp-1]
		}*/

		temp := len(cranes[from-1])
		cranes[to-1] = cranes[to-1] + cranes[from-1][temp-quantity:temp]
		cranes[from-1] = cranes[from-1][:temp-quantity]
	}	
	for i := 0; i < 9; i++{
		fmt.Print( cranes[i] [len(cranes[i])-1:len(cranes[i])])
	}
}