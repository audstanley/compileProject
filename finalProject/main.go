package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Read from file and sanitize the string:
	fmt.Println("Starting ...")
	data, err := ioutil.ReadFile("./originalCode.txt")
	if err != nil { // Error exists (Alex)
		panic("Could not read originalCode.xt")
	}
	dataString := string(data)
	sanitize(&dataString) // function that will remove comments, as well as extra spaces
	dataStringSlice := strings.Split(dataString, "\n")
	fmt.Println(dataString)

	lines := make(map[int][]string)
	for i, k := range dataStringSlice {
		lines[i+1] = strings.Split(k, " ")
	}

	// double check out data structure:
	for lNum, lArr := range lines {
		fmt.Print("Line: ", lNum, ", ")
		for _, word := range lArr {
			fmt.Print(word, " ")
		}
		fmt.Println()
	}

	fmt.Println("checking format...")
	errorCode, lineErr, ssBegin, ssEnd, errorStr := checkFormat(lines)

	for i, k := range ourVariables {
		fmt.Println("ourVariables: ", k, " at location ", i)
	}

	if errorCode == 1 {
		// handle all the error stuff LATER!!!
		fmt.Print("Some fucking error: ", lineErr, ssBegin, ssEnd, errorStr)
	} else {
		fmt.Println("Everything is good")
	}
}
