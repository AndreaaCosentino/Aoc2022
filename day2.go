package main

import(
	"fmt"
	"bufio"
	//"strconv"
	"os"
)

func main(){

	s := bufio.NewScanner(os.Stdin)
	/*A = rock (1) , B = paper (2) , C = scissor (3)*/
	score := 0
	for s.Scan(){
		line := s.Text()

		move := line[0]

		if line[2] == 'Y'{
			score += 3
			if move == 'A'{
				score += 1
			}else if move == 'B'{
				score += 2
			}else{
				score += 3
			}
		}else if line[2] == 'Z'{
			score += 6
			if move == 'A'{
				score += 2
			}else if move == 'B'{
				score += 3
			}else{
				score += 1
			}
		}else if line[2] == 'X'{
			if move == 'A'{
				score += 3
			}else if move == 'B'{
				score += 1
			}else{
				score += 2
			}
		}
		/*Soluzione parte 1*/
		/*if line[2] == 'Y'{
			score += 2
			if move == 'B'{
				score += 3
			}else if move == 'A'{
				score += 6
			}
		}else if line[2] == 'Z'{
			score += 3
			if move == 'C'{
				score +=3
			}else if move == 'B'{
				score += 6
			}
		}else if line[2] == 'X'{
			score += 1
			if move == 'C'{
				score += 6
			}else if move == 'A'{
				score += 3
			}
		}*/
	}
	fmt.Println(score)
}
