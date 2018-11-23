package main

import (
	"fmt"
	"regexp"
	"strings"
)

func validVarDig(s string) bool {
	r1, err := regexp.Match(`[a-zA-Z][a-zA-Z0-9]*|\d{1,}`, []byte(s))
	if err != nil {
		panic("vailidateVarDig regex match compiled failed: ")
	}
	return r1
}

func isDigit(s string) bool {
	r1, err := regexp.Match(`\d{1,}`, []byte(s))
	if err != nil {
		panic("vailidateVarDig regex match compiled failed: ")
	}
	return r1
}

func isNotOperator(s string) bool {
	r1, err := regexp.Match(`[^*+-/()=]`, []byte(s))
	if err != nil {
		panic("not operator regex match compiled failed")
	}
	return r1
}

func variableHasBeenDeclared(s string) bool {
	match := false
	if validVariable(s) {
		for _, y := range ourVariables {
			if s == y {
				match = true
				break
			}
		}
	}
	return match

}

func validateDefinition(toBeValidated []string) (int, string) {
	// toBeValidated is a string array of an individual line that is in a defined style format
	// return -1 to accept
	// anything else is location of error
	var passToMathChecker []string

	for _, k := range toBeValidated {
		passToMathChecker = append(passToMathChecker, k)
	}

	for i, v := range toBeValidated {
		fmt.Println("toBeValidated length: ", len(toBeValidated), " value: ", toBeValidated)
		fmt.Println("String Index:", i, " Value:", v)

		//		if v != "+" || v != "-" || v != "*" || v != "/" || v != "(" || v != ")" || v != "=" {
		if isNotOperator(v) {
			if !validVarDig(v) {
				if i == len(toBeValidated)-1 {
					if v != ";" {
						return i, "Missing a semicolon"
					}
				} else {
					fmt.Println("TOBEVALIDATED ERROR: ", v)
					return i, "Variable or Digit Error"
				}
			} else {
				passToMathChecker[i] = "a"
				fmt.Println("EXPRESSION: ", passToMathChecker)
				if !isDigit(v) {
					if !variableHasBeenDeclared(v) {
						return i, "Variable has not been Declared"
					}
				}
			}
		}
	}
	mathCode := mathrhs(strings.Join(passToMathChecker, ""))
	if mathCode == -1 {
		return -1, ""
	} else {
		return mathCode, "Invalid Mathematical Expression"
	}

}
