/*
	Richard	Stanley
	Alex
	Kenny

	This is the homework from week 02

*/
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

import (
	"github.com/golang-collections/collections/stack"
)

var variableStack = stack.New()
var calulationStack = stack.New()
var mapOfVariables = make(map[string]int)
var finalResult = 0

func operatorToFunction(s string, x, y int) int {
	if s == "+" {
		return x + y
	} else if s == "-" {
		return x - y
	} else if s == "*" {
		return x * y
	} else if s == "/" {
		if y == 0 {
			panic("You tried to divide by zero.")
		}
		return x / y
	}
	return 0
}

func populateMapWithStringGroup(s string, l int) {
	// this will take an individual string,
	// and populate a map with that string.

	wordRegex := regexp.MustCompile(`[a-zA-Z]+|[0-9]+`)
	digitRegex := regexp.MustCompile(`[0-9]+`)
	operatorRegex := regexp.MustCompile(`[+-\\*\\/]`)
	if wordRegex.MatchString(s) {
		_, keyExists := mapOfVariables[s]
		if !keyExists {
			if digitRegex.MatchString(s) {
				i, err := strconv.Atoi(s)
				if err == nil {
					mapOfVariables[s] = i
				} else {
					println(err)
					panic("There was an issue assigning variable: " + s + " to an integer \n")
				}
			} else {
				mapOfVariables[s] = 0
			}
		}
	} else if !operatorRegex.MatchString(s) {
		panic(s + " is not in the grammer")
	}

}

// ab 22 * c + ab + $
// [ab 22 * c + ab + $]
// ab 0 * c + ab + $

// tom jerry 123 + tom * - $
// tom: 2
// jerry: 3

func populateMapWithOneLongString(s string) {
	alphaRegex := regexp.MustCompile(`[a-zA-Z]+`)
	// this takes one big string and poplates a map with that sring
	strSliced := strings.Fields(s)
	for i, v := range strSliced {
		populateMapWithStringGroup(v, len(v))
		if v == "$" && i == len(strSliced)-1 { // once we see the $ sign
			for k := range mapOfVariables {
				// if the value of the variable has already been assigned:
				if alphaRegex.MatchString(k) {
					var j int
					// error checking
					print("Enter the value of ", k, ": ")
					_, err := fmt.Scanf("%d", &j)
					if err != nil {
						print(err)
						panic("That is not a valid number...exiting")
					}
					mapOfVariables[k] = j
				}

			}
		} else if v != "$" && i == len(strSliced)-1 { // if we don't see the $ sign as the last element
			panic("Please put a $ sign at the end of your expression...exiting")
		}
	}
	for k, v := range mapOfVariables {
		fmt.Println(k, v)
	}
	for _, v := range strSliced {
		calculateWithALongerString(v)
	}

}

func calculateWithALongerString(s string) {

	wordRegex := regexp.MustCompile(`[a-zA-Z]+|[0-9]+`)
	operatorRegex := regexp.MustCompile(`[+-\\*\\/]`)
	if wordRegex.MatchString(s) {
		variableStack.Push(s)
	} else if operatorRegex.MatchString(s) {
		if variableStack.Len() >= 1 {
			if calulationStack.Len() == 0 {
				xStr := variableStack.Pop().(string) // assert that the variable is a string
				yStr := variableStack.Pop().(string) // assert that the variable is a string
				x, _ := mapOfVariables[xStr]         // assign the value fot that variable for calculations
				y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
				calulationStack.Push(operatorToFunction(s, x, y))
				fmt.Println(x, s, y)

			} else if calulationStack.Len() == 1 {
				x, _ := calulationStack.Pop().(int)  // assert that the variable is an int
				yStr := variableStack.Pop().(string) // assert that the variable is a string
				y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
				calulationStack.Push(operatorToFunction(s, x, y))
				fmt.Println(x, s, y)

			} else if variableStack.Len() == 2 {
				xStr := variableStack.Pop().(string) // assert that the variable is a string
				yStr := variableStack.Pop().(string) // assert that the variable is a string
				x, _ := mapOfVariables[xStr]         // assign the value fot that variable for calculations
				y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
				calulationStack.Push(operatorToFunction(s, x, y))
				fmt.Println(x, s, y)

			} else {
				x, _ := calulationStack.Pop().(int)  // assert that the variable is an int
				yStr := variableStack.Pop().(string) // assert that the variable is a string
				y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
				calulationStack.Push(operatorToFunction(s, x, y))
				fmt.Println(x, s, y)

			}

		} else if calulationStack.Len() == 2 {
			x, _ := calulationStack.Pop().(int) // assert that the variable is an int
			y, _ := calulationStack.Pop().(int) // assert that the variable is an int
			calulationStack.Push(operatorToFunction(s, x, y))
			fmt.Println(x, s, y)

		}
	}
}
