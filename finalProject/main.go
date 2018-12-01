package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
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
	readFileTime := time.Now()
	var timeSinceTranspile int
	// Read from file and sanitize the string:
	data, err := ioutil.ReadFile("./originalCode.txt")
	if err != nil { // Error ioutilexists (Alex)
		panic("Could not read originalCode.txt")
	}
	afterReadFileTime := time.Since(readFileTime).Nanoseconds()
	dataString := string(data)
	sanitize(&dataString)                                                // function that will remove comments, as well as extra spaces
	ioutil.WriteFile("./output/sanitized.txt", []byte(dataString), 0644) // save the "cleaned up" data.
	dataStringSlice := strings.Split(dataString, "\n")                   // slice up by lines

	for indx := 0; indx < 1; indx++ { // let's get an average microseconds for file parsing
		startTime := time.Now() // start timer for the start of the compile time.
		lines := make(map[int][]string)
		for i, k := range dataStringSlice {
			lines[i+1] = strings.Split(k, " ")
		}
		errorCode, lineErr, ssBegin, ssEnd, errorStr := checkFormat(lines)
		if errorCode == 1 {
			consolePrintErrorSubString(lines, lineErr, ssBegin, ssEnd, errorStr)
		} else {
			timeSinceTranspile = int(time.Since(startTime).Nanoseconds() / 1000)
			fmt.Println("Reading from file ran in: ",
				afterReadFileTime/1000, "\tmicroseconds")
			fmt.Println("ProjectTranspiling ran in:", timeSinceTranspile, "\tmicroseconds")

			writeAndCompileTime := time.Now()
			ioutil.WriteFile("./output/main.go", []byte(strings.Join(ourOutput, "\n")), 0644)
			// once we write to a file, we can compile using
			// a local shell script that will use the go compiler to
			// output binaries for multiple operating systems

			// Compile to multiple platform binaries:
			cmd := exec.Command("./go-executable-build.sh", "main.go")
			_, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(cRed, err, cDefault)
			}
			fmt.Println("CompiledToBinary ran in:  ",
				time.Since(writeAndCompileTime).Nanoseconds()/1000000, "\tmilliseconds")
		}

		// reset all the variables
		ourOutput = []string{}
		ourVariables = []string{}
		f, fileErr := os.OpenFile("./compileTime.csv", os.O_APPEND|os.O_WRONLY, 0644)
		if fileErr != nil {
			fmt.Println(fileErr)
		}
		f.WriteString(strconv.Itoa(timeSinceTranspile) + "\n")
		f.Close()
		time.Sleep(time.Millisecond * 100)

	}

}
