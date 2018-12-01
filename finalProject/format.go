package main

import (
	"regexp"
	"sort"
	"strings"
)

var ourVariables []string
var ourOutput []string

func validVariable(s string) bool {
	r1, _ := regexp.Match(`^[a-zA-Z][a-zA-Z0-9]*$`, []byte(s))
	return r1
}

func showTest(s string) []string {
	r1 := regexp.MustCompile(`show\s?\(\s?(\w+)\s?\)\s?(;)?`)
	match := r1.FindStringSubmatch(s)
	return match
}

func mightBeTheWordInteger(s string) bool {
	// the word needs to at least start with int, otherwise we spellcheck,
	// and throw an error.
	r1, _ := regexp.Match(`[i][n][t][aeiouwsdr43]?[gftyhbv]?[aeiouwsdr43]?[redft54]?`, []byte(s))
	return r1
}

func mightBeTheWordShow(s string) bool {
	// if the first word starts with an s,
	// and the last two characters are any keypad letter near where they should be
	// throw this spelling error
	r1, _ := regexp.Match(`[s][hgtyujnb]?[aeiouiklp09]?[wqase321]?`, []byte(s))
	return r1
}

// we are returning int, int, int, string for error handling
// int => errorCode { 0: no error, 1: error }, int => lineNumber, int => start of error, int => end of error,
// string => error statement
func checkFormat(lines map[int][]string) (int, int, int, int, string) {
	//check for program declaration
	// <prog> -> program <id> ;
	if len(lines) >= 1 {

		if len(lines[1]) > 0 && lines[1][0] != "program" {
			return 1, 1, 0, len(lines[1]) - 1, " No Program Declaration"
		}

		if len(lines[1]) > 1 && !validVariable(lines[1][1]) {
			return 1, 1, 0, len(lines[1]) - 1, " Program Declaration Variable Not Valid"
		}

		if len(lines[1]) > 2 && lines[1][2] != ";" {
			return 1, 2, 0, 1, " Missing semicolon"
		}

	}

	// check for var declaration
	// <prog> -> ... var <dec-list> ...
	if len(lines) >= 2 {
		if len(lines[2]) > 0 {
			if lines[2][0] != "var" {
				return 1, 2, 0, len(lines[2]) - 1, " No Variable Declaration"
			}
		}
	}

	// run through each variable declaration, and
	// <dec-list> -> <dec> : <type> ;
	if len(lines) >= 3 {
		for i, word := range lines[3] {
			if i%2 == 0 { //
				if !validVariable(word) && i < len(lines[3])-3 {
					// fix math for the fucked up substring area
					return 1, 3, i, i + 1, " Not a valid variable"
				} else if validVariable(word) && i < len(lines[3])-3 {
					ourVariables = append(ourVariables, word)
				} else if i == len(lines[3])-2 && !mightBeTheWordInteger(word) && word != "integer" {
					return 1, 3, i, i + 1, " Must declare integer type, might be a misspelling"
				} else if i == len(lines[3])-2 && mightBeTheWordInteger(word) && word != "integer" {
					return 1, 3, i, i + 1, " the word integer is misspelled"
				} else if i == len(lines[3])-2 && word == "integer" {
					continue
				} else {
					return 1, 3, i + 1, len(lines[3]), "improper format in variable declaration"
				}

			} else if i%2 == 1 {
				if word != "," && i < len(lines[3])-3 {
					// fix math for the fucked up substring area
					return 1, 3, i, i + 1, " missing a comma"
				} else if i == len(lines[3])-1 {
					continue
				} else if i == len(lines[3])-3 && word != ":" {
					return 1, 3, i, i + 1, " Your variable declaration is missing a data type delimiter"
				}
			}
		}
	}

	// <prog> -> ... begin <stat-list> end
	if len(lines) >= 4 && lines[4][0] != "begin" {
		return 1, 4, 0, 1, " You need to have a begin statement before defining variables"
	}

	// populate the body of the code to be checked for mathematical expressions,
	// or for functions (such as the show func):
	// <stat-list> -> <stat> | <stat-list>  => the body map list(int {line integers} -> []string{ words })
	body := make(map[int][]string) // we are going to populate just the body of the data (between begin and end)
	var beginPosition, endPosition = 0, len(lines)
	for lNum, lArr := range lines {
		if len(lArr) > 0 {
			if lArr[0] == "begin" {
				if beginPosition == 0 {
					beginPosition = lNum
				} else {
					return 1, lNum, 0, 0, " You cannot have more than one begin statement."
				}
			} else if lArr[0] == "end" && lNum != len(lines) {
				return 1, lNum, 0, 0, " End statement needs to be at the end of the document"
			}
		}
	}
	if beginPosition == 0 {
		return 1, len(lines), 0, 0, " You are missing a begin statement in your document"
	} else if lines[len(lines)][0] != "end" {
		return 1, len(lines), 0, 0, " You are missing an end statement at the end of your document"
	} // populated the map list based on positions of begin line, and end line
	for i, k := range lines {
		if i < endPosition && i > beginPosition {
			body[i] = k
		}
	}
	// populate the main.go file, and verify the body of the function below
	var keys []int
	for k := range lines {
		keys = append(keys, k)
		ourOutput = append(ourOutput, "")
	}
	sort.Ints(keys)
	ourOutput[0] = "package main"
	ourOutput[1] = "import (\"fmt\")"
	// we should use a byte array to be o(1) efficient when appending
	variableDeclaration := []byte("var ")
	for i, k := range ourVariables {
		if i != len(ourVariables)-1 {
			variableDeclaration = append(variableDeclaration, []byte(k+", ")...)
		} else {
			variableDeclaration = append(variableDeclaration, []byte(k+" ")...)
		}
	}
	variableDeclaration = append(variableDeclaration, []byte("int")...)
	ourOutput[2] += string(variableDeclaration)
	ourOutput[3] = "func main() {"
	ourOutput[len(ourOutput)-1] = "}"
	// We're gonna parse the map of string arrays (slice) to see if it is a validated string format
	// If it is validated, then turn the mapped string into an expression
	// Move to next line

	// check for end declaration
	var lengthLines = len(lines)
	if lines[lengthLines][0] != "end" && len(lines[lengthLines]) == 1 { //last line of the file
		return 1, lengthLines, 0, len(lines[lengthLines][0]) - 1, " End Declaration Error"
	}

	//Validating left-hand variables of the body map
	// <stat> -> <write> | <assign>
	for lineNum, lineArr := range body {
		if lineNum > beginPosition && lineNum < endPosition {
			var valVar string
			valVar = body[lineNum][0]
			if valVar != "show" && len(lineArr) >= 4 {
				// will be checked if it has the right format
				if mightBeTheWordShow(valVar) {
					return 1, lineNum, 0, 1, "Incorrect spelling of the word 'show'"
				}
				if validVariable(valVar) { // <assign> -> <id> = <expr> ;
					if body[lineNum][1] == "=" {
						eCode, eStr := validateDefinition(body[lineNum]) // function defined in validateDefinition.go
						if eCode != -1 {                                 // check if there is NO error from mathrhs
							return 1, lineNum, eCode, eCode + 1, eStr
						} else {
							// the expression is goodbeginPosition
							linePopulated := []byte("\t")
							for i, k := range lineArr {
								if i != len(lineArr)-1 {
									linePopulated = append(linePopulated, []byte(k+" ")...)
								}
							}
							ourOutput[lineNum-1] = string(linePopulated)
						}
					} else {
						return 1, lineNum, 0, 0, " Invalid lefthand syntax+"
					}
				} else {
					if validVariable(valVar) != true {
						return 1, lineNum, 0, 0, " Invalid expression"
					}
				}
			} else if valVar == "show" { // <write> -> show( <id> ) ;
				showArr := showTest(strings.Join(body[lineNum], " "))
				if len(showArr) == 3 {
					// for loop to check the var
					if !variableHasBeenDeclared(showArr[1]) {
						// return error return lineNum, start = 2, end = 3
						return 1, lineNum, 2, 3, " Variable not declared"
					}
					if showArr[2] != ";" {
						return 1, lineNum, 4, 5, " Missing semicolon or invalid element"
					} else {
						// the show function is good:
						ourOutput[lineNum-1] = "\tfmt.Println(" + showArr[1] + ")"
					}

				} else {
					return 1, lineNum, 0, 0, " The show function is improperly formatted"
				}
			} else {
				return 1, lineNum, 0, 0, " Illegal expression"
			}
		}
	}
	return 0, 0, 0, 0, ""
}
