package main

import (
	"fmt"
	"regexp"
)

var ourVariables []string

func validVariable(s string) bool {
	r1, err := regexp.Match(`[a-zA-Z][a-zA-Z0-9]*`, []byte(s))
	if err != nil {
		panic("vailidateVariable regex match failed: ")
	}
	return r1
}

// we are returning int, int, int, string for error handling
// int => errorCode, int => lineNumber, int => start of error, int => end of error, string => error statement
func checkFormat(lines map[int][]string) (int, int, int, int, string) {

	//check for program declaration
	if len(lines) >= 1 {

		if len(lines[1]) > 0 && lines[1][0] != "program" {
			return 1, 1, 0, len(lines[1]) - 1, "No Program Declaration"
		}

		if len(lines[1]) > 1 && !validVariable(lines[1][1]) {
			return 1, 1, 0, len(lines[1]) - 1, "Program Declaration Variable Not Valid"
		}

		if len(lines[1]) > 2 && lines[1][2] != ";" {
			return 1, 2, 0, 1, "Missing semicolon"
		}

	}

	// check for var declaration
	if len(lines) >= 2 {
		if len(lines[2]) > 0 {
			if lines[2][0] != "var" {
				return 1, 2, 0, len(lines[2]) - 1, "No Variable Declaration"
			}
		}
	}

	if len(lines) >= 3 {

		for i, word := range lines[3] {
			fmt.Println("CHECKING: ", word, " at location: ", i)
			if i%2 == 0 {
				if !validVariable(word) && i < len(lines[3])-3 {
					// fix math for the fucked up substring area
					return 1, 3, len(lines[3][i]), len(lines[3][i]), " Not a valid variable"
				} else if validVariable(word) && i < len(lines[3])-3 {
					ourVariables = append(ourVariables, word)
				} else if i == len(lines[3])-2 && word != "integer" {
					return 1, 3, len(lines[3][i]), len(lines[3][i]), " Must declare integer type"
				}

			} else {
				fmt.Println("     math:", len(lines[3])-3)
				if word != "," && i < len(lines[3])-3 {
					// fix math for the fucked up substring area
					return 1, 3, len(lines[3][i]), len(lines[3][i]), " missing a comma"
				} else if i == len(lines[3])-3 && word != ":" {
					return 1, 3, len(lines[3][i]), len(lines[3][i]), " Your variable declaration is missing a data type delimiter"
				}
			}

		}
	}

	// check for end declaration
	var lengthLines = len(lines)
	if lines[lengthLines][0] != "end" {
		return 1, lengthLines, 0, len(lines[lengthLines][0]) - 1, "No End Declaration"
	}

	return 0, 0, 0, 0, ""
}
