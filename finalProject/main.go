package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"sort"
	"strings"
	"time"
)

const cRed = "\033[31m"
const cBlue = "\033[96m"
const cGreen = "\033[1;32m"
const cYellow = "\033[33m"
const cMagenta = "\033[35m"
const cDefault = "\033[0m"

func consolePrintErrorSubString(lines map[int][]string, lErr int, ssBegin int, ssEnd int, eStr string) {
	fmt.Println(cRed, "ERROR on Line ", lErr, ": ", eStr, cDefault)
	fmt.Print("SS:", cYellow, ssBegin, " ", ssEnd, cDefault)
	var keys []int
	for k := range lines {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, i := range keys {
		if lErr == i {
			for j, k := range lines[i] {
				if j == 0 {
					fmt.Print("\t")
				}
				if j >= ssBegin && j < ssEnd {
					fmt.Print(cBlue, k, " ", cDefault)
				} else if ssEnd == ssBegin {
					fmt.Print(cRed, k, " ", cDefault)
				} else {
					fmt.Print(cRed, k, " ", cDefault)
				}
			}
			fmt.Println()
		} else {
			for j, k := range lines[i] {
				if j == 0 {
					fmt.Print("\t")
				}
				fmt.Print(k, " ")
			}
			fmt.Println()
		}
	}
}

func main() {
	startTime := time.Now()

	// Read from file and sanitize the string:
	data, err := ioutil.ReadFile("./originalCode.txt")
	if err != nil { // Error ioutilexists (Alex)
		panic("Could not read originalCode.txt")
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
		consolePrintErrorSubString(lines, lineErr, ssBegin, ssEnd, errorStr)
		//fmt.Print("Error on line: ", lineErr, ssBegin, ssEnd, errorStr, "\n", cDefault)
	} else {
		fmt.Println("Everything is good")
		for i, k := range ourOutput {
			fmt.Print(cBlue, "LINE: ", i, cMagenta, k, cDefault, "\n")
		}
		fmt.Println("FinalProject ran in: ", time.Since(startTime).Seconds())
		ioutil.WriteFile("./output/main.go", []byte(strings.Join(ourOutput, "\n")), 0644)
		// once we write to a file, we can compile using
		// a local shell script that will use the go compiler to
		// output binaries for multiple operating systems
		cmd := exec.Command("./go-executable-build.sh", "main.go")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(cRed, err, cDefault)
		} else {
			fmt.Println(string(out))
		}
	}
}
