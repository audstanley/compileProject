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
// int => errorCode { 0: no error, 1: error }, int => lineNumber, int => start of error, int => end of error,
// string => error statement
func checkFormat(lines map[int][]string) (int, int, int, int, string) {
	fmt.Println("CHECKING FORMAT")
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
					return 1, 3, len(lines[3][i]), len(lines[3][i]),
						" Your variable declaration is missing a data type delimiter"
				}
			}
		}
	}

	// populate the body of the code to be checked for mathematical expressions, or for functions (such as the show func):
	body := make(map[int][]string) // we are going to popluate just the body of the data (between begin and end)
	var beginPosistion, endPosition = 0, len(lines)
	for lNum, lArr := range lines {
		if len(lArr) > 0 {
			if lArr[0] == "begin" {
				if beginPosistion == 0 {
					beginPosistion = lNum
				} else {
					return 1, lNum, 0, 0, "You cannot have more than one begin statement."
				}
			} else if lArr[0] == "end" && lNum != len(lines) {
				return 1, lNum, 0, 0, "Inproper end statement"
			}
		}
	}
	if beginPosistion == -1 {
		return 1, len(lines), 0, 0, "You are missing a begin statement"
	} else if lines[len(lines)][0] != "end" {
		return 1, len(lines), 0, 0, "You are missing an end statement"
	}
	for i, k := range lines {
		if i < endPosition && i > beginPosistion {
			body[i] = k
		}
	}

	fmt.Println("THE BODY OF THE SHIT IS: ", body)

	// Youre gonna pass in the map of string arrays
	// We're gonna parse the map of string arrays (slice) to see if it is a validated string format
	// If it is validated, then turn the mapped string into an expression
	// Move to next line

	// check for end declaration
	var lengthLines = len(lines)
	if lines[lengthLines][0] != "end" {
		return 1, lengthLines, 0, len(lines[lengthLines][0]) - 1, " No End Declaration"
	}

	//Validating left-hand variables of the body map
	for lineNum, lineArr := range body {
		if lineNum > beginPosistion && lineNum < endPosition {
			var valVar string
			valVar = body[lineNum][0]
			if valVar != "show" && len(lineArr) >= 4 { //Every string that is not "show" will be checked if it has the right format
				if validVariable(valVar) {
					fmt.Println("VALIDATE A MATH ESPRESSION")
					if body[lineNum][1] == "=" {
						// Here is where we will parse through the (call the function Alex was working on)
						eCode, eStr := validateDefinition(body[lineNum]) // function defined in validateDefinition.go
						if eCode != -1 {
							return 1, lineNum, 0, 0, eStr
						}
					} else {
						return 1, lineNum, 0, 0, " Invalid lefthand syntax+"
					}
				} else {
					if validVariable(valVar) != true {
						return 1, lineNum, 0, 0, " Invalid variable"
					}
				}
			} else if valVar == "show" {
				if body[lineNum][1] == "(" && body[lineNum][len(lineArr)-2] == ")" && body[lineNum][len(lineArr)-1] == ";" {
					// if content inside is an integer or an existing variable then accept
					// we can append content to the goOutputStruct
					fmt.Println("The show function is Golden")

				} else {
					fmt.Println("ERROR: ", body[lineNum])
					return 1, lineNum, 0, 0, " Invalid id inside 'show'"
				}
			} else {
				fmt.Println(body[lineNum])
				return 1, lineNum, 0, 0, "The show function is improperly formatted"
			}
		}
	}

	return 0, 0, 0, 0, ""
}
