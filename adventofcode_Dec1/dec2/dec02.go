package main

// sample input
// Game 1: 20 green, 3 red, 2 blue; 9 red, 16 blue, 18 green; 6 blue, 19 red, 10 green; 12 red, 19 green, 11 blue
// Game 2: 12 green, 3 blue, 16 red; 6 red, 4 blue, 12 green; 11 green, 4 red, 3 blue; 8 green, 15 red, 5 blue
// Game 3: 13 blue, 4 red, 8 green; 2 green, 4 red, 19 blue; 5 blue; 10 blue, 6 green, 2 red; 19 blue; 8 blue, 6 red

//which games would have been possible if there are only 12 red, 13 green and 14 blue.
//for each game if a single reveal revealed more than the set amount, dont include it

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)


var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}



func main() {
	file, _ := os.Open("input02.txt")
	defer file.Close()
	idSum := 0
	maxGreen := 13
	maxRed := 12
	maxBlue := 13

	//iterate through each line, iterate through each game, if no color pulled is higher than a max, add to idSum.



	scanner := bufio.NewScanner(file)
	for scanner.Scan (){
		//split to get game id
		currline := scanner.Text()
		gameStr := strings.Split(currline, ":")
		holdId := 0
		err := *new(error)
		fmt.Println(err)
		for _, part := range gameStr {
			// if part contains game, get id for sum, split second part
			//set bool
			validGame := true
			if(strings.Contains(part, "Game")) {
				//hold id as int, just in case
				holdId, err = strconv.Atoi(clearString(part))
			}else {
				//must be start of draw listings, split into draws
				draws := strings.Split(part, ";")
				for _, drawPart := range draws {
					//split into colors 
					colors := strings.Split(drawPart, ",")
					for _, colorPart := range colors {
						if(strings.Contains(drawPart,"green")){
							colorCount, _ := strconv.Atoi(clearString(colorPart))
							if(colorCount > maxGreen){
								validGame = false
							}
						}
						if(strings.Contains(drawPart,"red")){
							colorCount, _ := strconv.Atoi(clearString(colorPart))
							if(colorCount > maxRed){
								validGame = false
							}
						}
						if(strings.Contains(drawPart,"blue")){
							colorCount, _ := strconv.Atoi(clearString(colorPart))
							if(colorCount > maxBlue){
								validGame = false
							}
						}
					}
				}
			}
			if(validGame){
				idSum += holdId
			}
		}
		fmt.Println("Valid Game Sum: ", idSum);


	}
	

}



