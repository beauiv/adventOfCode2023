package main

//sample
// ...317..........214.....................................751.................................630...479..205....41.993............416.........
// ...*....813........%....572........%...629.154......518....*....365..................-.......*.......#.....................422...........661
// 269.......*...58...........=......264.....*..........*......937.-...........235...303.........848..............195.....154*.........144.-...

//traverse each line, get start and end pos (+1 for adj) of each number
//check prev line in range for symbols, same for following line, and currline for adjacents

type number struct {
	value int
	startLoc int
	endLoc int
	counted bool
	line int
}

type symbol struct {
	loc int
	line int
}

func main() {
	var validSymbols  = [14]String{"*","%","#","!","@","$","^","&","(",")","/","=","-","_"}
	numberSlice := []number
	symbolSlice := []symbol
	lineCount := 0

	file, _ := os.Open("input03.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan (){

		for i := 0; i < 
		
	}

}

