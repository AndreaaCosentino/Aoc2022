package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type point struct{
	x,y int
}

const TOSEARCH = 2000000

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	manDist := make(map[point]int)
	for scanner.Scan(){

		line := scanner.Text()
		line = strings.Replace(line," ","",-1)
		inp := strings.Split(line,"=")
		inp = inp[1:]
		inp[0] = inp[0][:strings.Index(inp[0],",")]
		inp[1] = inp[1][:strings.Index(inp[1],":")]
		inp[2] = inp[2][:strings.Index(inp[2],",")]

		stPoint := createPoint(inp[0],inp[1])
		beacon := createPoint(inp[2],inp[3])
		manDist[stPoint] = manhattanDistance(stPoint,beacon)
	}
	for   y := 0;y < 4000000; y++{
		for x := 0; x  < 4000000; x++{
			if confront(manDist,point{x,y},&x){
				fmt.Println(x*4000000 + y)
				break
			}
		} 
	}

}


func confront(manDist map[point]int, p point,i *int) bool{
	for k,v := range manDist{
		if !(v < manhattanDistance(k,p)){
			(*i) +=  v - manhattanDistance(k,p)
			return false
		}
	}
	return true
}
/*func fillMap(noBeacon map[point]int, startingPoint point, currPoint point,dist int,i int,dispX int, dispY int){

	if manhattanDistance(startingPoint,currPoint) <= dist{
		noBeacon[currPoint] = 1
		if(noBeacon[point{currPoint.x-dispX,currPoint.y-dispY}] != -1){noBeacon[point{currPoint.x-dispX,currPoint.y-dispY}] = 1}
		if(noBeacon[point{currPoint.x-dispX,currPoint.y}] != -1){noBeacon[point{currPoint.x-dispX,currPoint.y}] = 1}
		if(noBeacon[point{currPoint.x,currPoint.y-dispY}] != -1){noBeacon[point{currPoint.x,currPoint.y-dispY}] = 1}
	}else{return}
		tx := dispX+1
		fillMap(noBeacon,startingPoint,point{currPoint.x+1,currPoint.y},dist,1,tx,dispY)
		ty := dispY+1
		fillMap(noBeacon,startingPoint,point{currPoint.x,currPoint.y-1},dist,1,dispX,ty)
}
*/

func createPoint(x string, y string ) point{
		x1,_ := strconv.Atoi(x)
		y1,_ := strconv.Atoi(y)

		return point{x1,y1};
}

func manhattanDistance(p1 point, p2 point)int{
	return int(math.Abs(float64(p1.x - p2.x)) + math.Abs(float64(p1.y - p2.y)));
}