package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numbersOnly = regexp.MustCompile(`\d+`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner1 := bufio.NewScanner(file)
	scanner2 := bufio.NewScanner(file)
	ptSum := 0
	playTracker := []int{}
	lineCounter := 0
	for scanner1.Scan() {
		lineCounter++
	}
	//appending slices to fit
	for i := 0; i < lineCounter; i++ {
		playTracker = append(playTracker, 1)
	}
	file.Seek(0, 0)
	for scanner2.Scan() {
		split1 := strings.Split(scanner2.Text(), ":")
		split2 := strings.Split(split1[1], "|")
		cardNum, _ := strconv.Atoi(numbersOnly.FindAllString(split1[0], -1)[0])
		ptSum += getPointsForLine(split2, split1)

		for i := 0; i < playTracker[cardNum-1]; i++ {
			//fmt.Println(playTracker)
			matches := getMatchesForLine(split2, split1)
			//increase play counts
			for k := 0; k < matches; k++ {
				playTracker[cardNum+k] = playTracker[cardNum+k] + 1
			}
		}
	}
	sumOfCards := 0
	for _, num := range playTracker {
		sumOfCards += num
	}
	fmt.Println("Total Points: ", ptSum)     //pt1
	fmt.Println("Total Cards: ", sumOfCards) //pt2
}
func getPointsForLine(playPart []string, gameId []string) int {
	// cardNum := numbersOnly.FindAllString(gameId[0], -1)
	totalPoints := 0.0
	winningNums := numbersOnly.FindAllString(playPart[0], -1)
	playNums := numbersOnly.FindAllString(playPart[1], -1)
	matches := 0 //also known as total cards in pt2
	for _, playNum := range playNums {
		for _, winNum := range winningNums {
			if strings.EqualFold(playNum, winNum) {
				matches = matches + 1
			}
		}
	}
	//fmt.Println(cardNum, ": Matches:", matches)
	if matches > 0 {
		totalPoints += math.Pow(float64(2), float64(matches-1))
	}
	return int(totalPoints)
}
func getMatchesForLine(playPart []string, gameId []string) int {
	winningNums := numbersOnly.FindAllString(playPart[0], -1)
	playNums := numbersOnly.FindAllString(playPart[1], -1)
	matches := 0 //also known as total cards in pt2
	for _, playNum := range playNums {
		for _, winNum := range winningNums {
			if strings.EqualFold(playNum, winNum) {
				matches = matches + 1
			}
		}
	}
	//fmt.Println(cardNum, ": Matches:", matches)
	return int(matches)
}
