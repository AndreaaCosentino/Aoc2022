package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type bluePrint struct{
	id int 
	robotOre cost 
	robotClay cost 
	robotObs cost 
	robotGeode cost
	maxOre int 
	maxClay int 
	maxObs int
}

type cost struct{
	ore int 
	clay int 
	obs int 
}

type inventory struct{
	robotOre int 
	robotClay int 
	robotObs int 
	robotGeode int 
	ore int 
	clay int 
	obs int
	geo int 
}

var best map[int]int


func main(){
	scanner := bufio.NewScanner(os.Stdin)

	bluePrintList := make([]bluePrint,0)

	for scanner.Scan(){
		line := scanner.Text()

		input := strings.Split(line," ")
		input = input[1:]

		input[0] = input[0][:len(input[0])-1]

		id,_ := strconv.Atoi(input[0])
		robotOreCost,_ := strconv.Atoi(input[5])
		robotClayCost,_ := strconv.Atoi(input[11])
		robotOre := cost{robotOreCost,0,0}
		robotClay := cost{robotClayCost,0,0}

		robotObsCost_Ore,_ := strconv.Atoi(input[17])
		robotObsCost_Clay,_ := strconv.Atoi(input[20])
		robotObs := cost{robotObsCost_Ore,robotObsCost_Clay,0}

		robotGeodeCost_Ore,_ := strconv.Atoi(input[26])
		robotGeodeCost_Obs,_ := strconv.Atoi(input[29])
		robotGeode := cost{robotGeodeCost_Ore,0,robotGeodeCost_Obs}


		maxOre := max(max(robotClayCost,robotObsCost_Ore),robotGeodeCost_Ore)
		maxClay := robotObsCost_Clay
		maxObs := robotGeodeCost_Obs

		bluePrintList = append(bluePrintList,bluePrint{id,robotOre,robotClay,robotObs,robotGeode,maxOre,maxClay,maxObs})
	}

	res := 1
	j:=0
	for _,v := range bluePrintList{
		best = make(map[int]int)
		_ = searchBest(v,32,inventory{1,0,0,0,0,0,0,0})
		res *= best[1]
		fmt.Println(j,res,v.id)
		j++
	}
	fmt.Println(res)
}

func max(i int, j int) int{
	if i > j {
		return i
	}else{return j}
}
func searchBest(currentBluePrint bluePrint, time int, currentInventory inventory) int{
	if time == 0{
		//fmt.Println(currentInventory)
		return 0
	}

	max := 0


	tempInventoryOre := currentInventory
	tempInventoryClay := currentInventory
	tempInventoryObs := currentInventory
	tempInventoryGeo := currentInventory

	currentInventory.ore += currentInventory.robotOre
	currentInventory.clay += currentInventory.robotClay
	currentInventory.obs += currentInventory.robotObs
	currentInventory.geo += currentInventory.robotGeode

	if v,ok := best[time]; ok{
		if v > currentInventory.geo{
			return 0
		}else{
			best[time] = currentInventory.geo
		}
	}else{best[time] = currentInventory.geo}

	// IF YOU CAN, BUILD A GEODE ROBOT
	if currentBluePrint.robotGeode.ore <= tempInventoryGeo.ore && currentBluePrint.robotGeode.obs <= tempInventoryGeo.obs{
		tempInventoryGeo.ore += tempInventoryGeo.robotOre
		tempInventoryGeo.clay += tempInventoryGeo.robotClay
		tempInventoryGeo.obs += tempInventoryGeo.robotObs
		tempInventoryGeo.geo += tempInventoryGeo.robotGeode

		tempInventoryGeo.ore -= currentBluePrint.robotGeode.ore
		tempInventoryGeo.obs -= currentBluePrint.robotGeode.obs
		tempInventoryGeo.robotGeode++
		temp := searchBest(currentBluePrint,time-1,tempInventoryGeo)
		if temp > max{max = temp}
	}

	// IF YOU CAN, BUILD AN ORE ROBOT
	if currentBluePrint.robotOre.ore <= tempInventoryOre.ore && tempInventoryOre.robotOre < currentBluePrint.maxOre{
		tempInventoryOre.ore += tempInventoryOre.robotOre
		tempInventoryOre.clay += tempInventoryOre.robotClay
		tempInventoryOre.obs += tempInventoryOre.robotObs
		tempInventoryOre.geo += tempInventoryOre.robotGeode

		tempInventoryOre.ore -= currentBluePrint.robotOre.ore
		tempInventoryOre.robotOre++
		temp := searchBest(currentBluePrint,time-1,tempInventoryOre)
		if temp > max{max = temp}
	}
	// IF YOU CAN, BUILD A CLAY ROBOT
	
	if currentBluePrint.robotClay.ore <= tempInventoryClay.ore && tempInventoryClay.robotClay < currentBluePrint.maxClay{
		tempInventoryClay.ore += tempInventoryClay.robotOre
		tempInventoryClay.clay += tempInventoryClay.robotClay
		tempInventoryClay.obs += tempInventoryClay.robotObs
		tempInventoryClay.geo += tempInventoryClay.robotGeode

		tempInventoryClay.ore -= currentBluePrint.robotClay.ore
		tempInventoryClay.robotClay++
		temp := searchBest(currentBluePrint,time-1,tempInventoryClay)
		if temp > max{max = temp}
	}

	// IF YOU CAN, BUILD AN OBSIDIAN ROBOT
	if currentBluePrint.robotObs.ore <= tempInventoryObs.ore && currentBluePrint.robotObs.clay <= tempInventoryObs.clay && tempInventoryObs.robotObs < currentBluePrint.maxObs{
		tempInventoryObs.ore += tempInventoryObs.robotOre
		tempInventoryObs.clay += tempInventoryObs.robotClay
		tempInventoryObs.obs += tempInventoryObs.robotObs
		tempInventoryObs.geo += tempInventoryObs.robotGeode

		tempInventoryObs.ore -= currentBluePrint.robotObs.ore
		tempInventoryObs.clay -= currentBluePrint.robotObs.clay
		tempInventoryObs.robotObs++
		temp := searchBest(currentBluePrint,time-1,tempInventoryObs)
		if temp > max{max = temp}
	}

	// DONT SPEND ANYTHING, KEEP PRODUCING

	if currentInventory.robotOre < currentBluePrint.maxOre{
		geo := searchBest(currentBluePrint,time-1,currentInventory)
		if geo > max {max = geo}
	}	

	return max
}

