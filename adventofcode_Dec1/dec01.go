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
var textNumbers = [10]string{"z0o","o1e","t2o","t3e","f4r","f5e","s6x","s7n","e8t","n9e"}

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStr := ""
		convString:=scanner.Text();
		fmt.Println("Old String: ", convString)
		for i, number := range textNumbers {	
			//fmt.Println("Replacing "+ number + "with ", i)
			convString = strings.Replace(convString, number,strconv.Itoa(i),-1)
		}
		fmt.Println("New String: ", convString)
		cleanStr := clearString(convString)
		fmt.Println("Clean String: ", cleanStr)
		for i := 0; i < len(cleanStr); i++ {
			if i == 0 || i == len(cleanStr)-1 {
				numStr = numStr + string(cleanStr[i])
			}
		}
		if len(numStr) == 1 {
			numStr = numStr + numStr
			fmt.Println("only one digit, duping: " + numStr)
		}
		hold, _ := strconv.Atoi(numStr)
		sum += hold
		fmt.Println("Number this time: ", hold)
		fmt.Println("Sum so far:", sum)
		count++
	}
	fmt.Println("total: ", sum)
	fmt.Println("lines processed: ", count)

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
