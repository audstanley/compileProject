package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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

		// inputs:
		// ab 22 * c + ab + $
		// tom jerry 123 + tom * - $

		populateMapWithOneLongString(postfixExpressionWithoutNewLine)

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
