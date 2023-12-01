package main

//get all numbers from line, only take first and last number, make int, add.
//part 2, get text representation of numbers and consider them as well.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)
var textNumbers = [10]string{"zero","one","two","three","four","five","six","seven","eight","nine"}

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStr := ""
		convString:=scanner.Text();
		for i, number := range textNumbers {
			convString = strings.Replace(convString, number,strconv.Itoa(i),-1)
		}
		cleanStr := clearString(convString)
		for i := 0; i < len(cleanStr); i++ {
			if i == 0 || i == len(cleanStr)-1 {
				numStr = numStr + string(cleanStr[i])
			}
		}
		if len(numStr) == 1 {
			numStr = numStr + numStr
		}
		hold, _ := strconv.Atoi(numStr)
		sum += hold
		fmt.Println("Number this time: ", hold)
		fmt.Println("Sum so far:", sum)
	}
	fmt.Println("total: ", sum)

	//test change

	// str := clearString("znggdvvkjthreethree79eight4")
	// numStr := ""

	// for i := 0; i < len(str); i++ {
	// 	if i == 0 || i == len(str)-1 {
	// 		numStr = numStr + string(str[i])
	// 	}
	// }
	// fmt.Println(numStr)

	// if err != nil {
	// 	//okay what now
	// }
	// fmt.Println(sum)
}
