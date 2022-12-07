package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type directory struct{
	size int 
	directories []string
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	mapDirectories := make(map[string]directory)
	currentPath := make([]string,0)
	currentString := ""
	for scanner.Scan(){
		line := scanner.Text()

		words := strings.Split(line," ")
		
		if words[0] == "$"{
			if words[1] == "cd"{
				if words[2] == ".."{
					currentPath = currentPath[:len(currentPath)-1]
				}else if words[2] == "\\"{
					currentPath = currentPath[:1]
				}else{
					currentPath = append(currentPath,words[2])
				}
			}else if words[1] == "ls"{
				currentString = createPath(currentPath)
				if _,ok := mapDirectories[currentString]; !ok{
					var newDir directory
					newDir.size = 0
					newDir.directories =  make([]string,0)
					mapDirectories[currentString] = newDir
				}
			}
		}else{
			entry := mapDirectories[currentString]
			if words[0] == "dir"{
				entry.directories = append(entry.directories,currentString+words[1])
			}else{
				s,_ := strconv.Atoi(words[0])
				entry.size += s
			}
			mapDirectories[currentString] = entry
		}
	}
	for k,_ := range mapDirectories{
		_ = calculateSize(k,mapDirectories)
	}

	
	enough := 30000000 - (70000000-mapDirectories["/"].size)
	min := 70000000
	for k,_ := range mapDirectories{
		temp := mapDirectories[k].size 
		if temp < min && temp >= enough{
			min = temp
		}
	}
	fmt.Println(min)
}

func calculateSize(dir string, mapDir map[string]directory) int{
		sum := 0
		if len(mapDir[dir].directories) != 0{
			for k,_ := range mapDir[dir].directories{
				sum += calculateSize(mapDir[dir].directories[k],mapDir)
			}
		}
		entry := mapDir[dir]
		entry.directories = entry.directories[:0]
		entry.size += sum
		mapDir[dir] = entry
		return mapDir[dir].size
}


func createPath(currentPath []string) string{
	path := ""
	for i := 0; i < len(currentPath); i++{
		path += currentPath[i]
	}
	return path
}