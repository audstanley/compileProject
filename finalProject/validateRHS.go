package main

import (
	"fmt"
	"strings"
)

func validVarDig(s string) bool {
	r1, err := regexp.Match(`[a-zA-Z][a-zA-Z0-9]*|\d{1,}`, []byte(s))
	if err != nil {
		panic("vailidateVariable regex match failed: ")
	}
	return r1
}

func validateDefinition(toBeValidated []string) (int, string ) {
// -1 to accept
// anything else is location of error

	passToMathChecker := []string
	match := false

	for i, v := range toBeValidated {
		fmt.Println("String Index:", i, " Value:", v)

		if v != "+" || v != "-" || v != "*" || v != "/" || v != "(" || v != ")" || v != "=" {
			if !validVarDig(v) {
				return i, "not a valid expression, Variable or Digit Error"
			} else {
				passToMathChecker[i] = v

				if validVariable(v) {
					passToMathChecker[i] = "a"
					for _, y := range ourVariables {
						if v == y {
							match == true
							break
						}
					}
				}
				if match == false  {
					return i, "not a valid variable"
				}
			}
		}
	}
	match = false	
}
