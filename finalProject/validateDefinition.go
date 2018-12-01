package main

import (
	"regexp"
	"strings"
)

func validVarDig(s string) bool {
	r1, err := regexp.Match(`[a-zA-Z][a-zA-Z0-9]*|\d{1,}`, []byte(s))
	if err != nil {
		panic("validateVarDig regex match compiled failed: ")
	}
	return r1
}

func isDigit(s string) bool {
	r1, err := regexp.Match(`[+-]?\d{1,}`, []byte(s))
	if err != nil {
		panic("validateVarDig regex match compiled failed: ")
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
		//		if v != "+" || v != "-" || v != "*" || v != "/" || v != "(" || v != ")" || v != "=" {
		if isNotOperator(v) {
			if !validVarDig(v) {
				if i == len(toBeValidated)-1 {
					if v != ";" {
						return i, "Missing a semicolon"
					}
				} else {
					return i, "Variable or Digit Error"
				}
			} else {
				passToMathChecker[i] = "a" // change all variables and digits into the letter 'a'
				if !isDigit(v) {
					if !variableHasBeenDeclared(v) {
						return i, "Variable has not been Declared"
					}
				}
			}
		}
	}
	mathCode := mathrhs(strings.Join(passToMathChecker, "")) // pass the whole expression into our parser, no spaces
	// if no error, mathCode will == -1, otherwise mathCode == subStringIndexWhereErrorOccurred
	if mathCode == -1 {
		return -1, ""
	} else {
		return mathCode, "Invalid Mathematical Expression" // this will return the location of error,
		// which will be the index of the string array in the line map.
	}

}
