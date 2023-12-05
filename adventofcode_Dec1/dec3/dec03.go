package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//disclaimer: heavily used anothers solution to learn how to use maps, and structs in go. not my best day of advent of code.
//walked through and rewrote each line myself to try and understand it all, got quite confused on where to init the checkforSymbolsHere map
//as well as the one liner for how to check if a value exists in a map, if _, exists := symbols[p]; exists {
//learned alot though!
//source that I followed along with: https://github.com/bsadia/aoc_goLang/blob/master/day03/main.go

type Point struct {
	X, Y int
}

func (p Point) add_point(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func main() {

	file, _ := os.ReadFile("input03.txt")

	//get symbols
	symbols := map[Point]string{}
	for y, line := range strings.Fields(string(file)) {
		for x, character := range line {
			if character != '.' && !unicode.IsDigit(character) {
				symbols[Point{x, y}] = string(character)
			}
		}
	}

	//get nums (parts)
	numParts := map[Point][]int{}
	getNums := regexp.MustCompile(`\d+`)
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for y, line := range strings.Fields(string(file)) {
		//get line
		for _, regexNum := range getNums.FindAllStringIndex(line, -1) {
			//get each number from line
			checkforSymbolsHere := map[Point]bool{}
			//a place to store where to check for symbols based on each number
			for x := regexNum[0]; x < regexNum[1]; x++ {
				//for each character that is part of a number generate a point to check for symbols
				for _, direction := range directions {
					checkforSymbolsHere[Point{x, y}.add_point(direction)] = true
				}
			}
			//get the number
			n, _ := strconv.Atoi(line[regexNum[0]:regexNum[1]])
			//for each place that we check for symbols
			for p := range checkforSymbolsHere {
				//if a symbol exists
				if _, exists := symbols[p]; exists {
					//add that number to our part list
					numParts[p] = append(numParts[p], n)
				}
			}
		}
	}

	part_numbers := 0
	for _, values := range numParts {
		fmt.Println(values)
		for _, value := range values {
			part_numbers += value
		}
	}
	fmt.Println(part_numbers)
	gearValue := 0
	for symbolLoc, values := range numParts {
		if symbols[symbolLoc] == "*" && len(values) == 2 {
			fmt.Println(symbols[symbolLoc], values)
			gearValue += values[0] * values[1]
		}

	}
	fmt.Println(gearValue)

}

// ...88
// *....
// 88...
// ..*..
