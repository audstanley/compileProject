/*
	Richard	Stanley
	Alex
	Kenny

	This is the homework from week 01

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
	This recursive function will populate the map of variables.
	So we can query the user to populate the values of the map
	to the associated key.
*/
func populateMap(s string) {

	charRegex := regexp.MustCompile(`[a-z+-\\*\\/]`)
	alphaRegex := regexp.MustCompile(`[a-z]`)
	if len(s) != 0 {
		if len(s) > 0 {
			if charRegex.MatchString(s[:1]) {
				if alphaRegex.MatchString(s[:1]) {
					_, keyExists := mapOfVariables[s[:1]]
					if !keyExists {
						mapOfVariables[s[:1]] = 0
					}
				}
				populateMap(string(s[1:]))
			}

		}

		if len(s) == 1 && string(s[0]) != "$" {
			println("Please end your post fix expression with: $")
		} else if string(s[0]) == "$" {
			for k := range mapOfVariables {
				var i int
				// error checking loop
				for {
					print("Enter the value of ", k, ": ")
					_, err := fmt.Scanf("%d", &i)
					if err == nil {
						break
					}
				}

				mapOfVariables[k] = i
				println()
			}
		}
	}
}

/*
	Once the variables are mapped (from the previous function),
	we can push the variables to a stack, and when we see an operator,
	we will immediately use the operator to calculate the variables in the stack
	whose values are called from the mapOfVariables map

*/

func calculate(s string) {

	alphaRegex := regexp.MustCompile(`[a-z]`)
	operatorRegex := regexp.MustCompile(`[+-\\*\\/]`)

	if len(s) != 0 && s[0] != 117 {
		if alphaRegex.MatchString(s[:1]) {
			variableStack.Push(s[:1])
			calculate(s[1:])
		} else if operatorRegex.MatchString(s[:1]) {
			if variableStack.Len() >= 1 {

				if calulationStack.Len() == 0 {
					xStr := variableStack.Pop().(string) // assert that the variable is a string
					yStr := variableStack.Pop().(string) // assert that the variable is a string
					x, _ := mapOfVariables[xStr]         // assign the value fot that variable for calculations
					y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
					calulationStack.Push(operatorToFunction(s[:1], x, y))
					calculate(s[1:])

				} else if variableStack.Len() == 2 {
					xStr := variableStack.Pop().(string) // assert that the variable is a string
					yStr := variableStack.Pop().(string) // assert that the variable is a string
					x, _ := mapOfVariables[xStr]         // assign the value fot that variable for calculations
					y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
					calulationStack.Push(operatorToFunction(s[:1], x, y))
					calculate(s[1:])

				} else {
					x, _ := calulationStack.Pop().(int)  // assert that the variable is an int
					yStr := variableStack.Pop().(string) // assert that the variable is a string
					y, _ := mapOfVariables[yStr]         // assign the value fot that variable for calculations
					calulationStack.Push(operatorToFunction(s[:1], x, y))
					calculate(s[1:])
				}

			} else if calulationStack.Len() == 2 {
				x, _ := calulationStack.Pop().(int) // assert that the variable is an int
				y, _ := calulationStack.Pop().(int) // assert that the variable is an int
				calulationStack.Push(operatorToFunction(s[:1], x, y))

			}
		}
	}
}

func runWithASingleStringNoSpaces() {
	println()
	println("*******************************************")
	println("*          The postfix calculator         *")

	for {
		println("*******************************************")
		println()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a postfix expression with $ at the end: ")
		postfixExpression, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		postfixExpressionWithoutNewLine := strings.Replace(postfixExpression, "\n", "", -1)

		populateMap(postfixExpressionWithoutNewLine)

		// ab*ac+*$  			where: a = 2, b = 3, c = 4 	resultShouldBe: 36
		// beef*++$				where: b = 2, e = 3, f = 4	resultShouldBe: 17
		// aaaabbbb***+++*$		where: a = 2, b = 3			resultShouldBe: 340

		calculate(postfixExpressionWithoutNewLine)

		if calulationStack.Len() == 1 && variableStack.Len() == 0 {
			finalResult = calulationStack.Pop().(int)
			println("Final Result:", finalResult)
		} else {
			println("You expression is malformed")
		}

		print("Continue(y/n)? ")
		contBuffer := bufio.NewReader(os.Stdin)
		cont, _ := contBuffer.ReadString('\n')
		if cont == "n\n" {
			break
		}

		// clear out the stacks, in case user continues
		for i := 0; i < variableStack.Len()-1; i++ {
			variableStack.Pop()
		}
		for i := 0; i < calulationStack.Len()-1; i++ {
			calulationStack.Pop()
		}
		mapOfVariables = make(map[string]int)
		println()
	}
	println()
	println("**************** Done *********************")
	println()

}

func shree() {
	fmt.Printf("fucking shree")
}
