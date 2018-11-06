package main

/*
	Richard Stanley
	Row 5

*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

const cRed = "\033[31m"
const cGreen = "\033[1;32m"
const cYellow = "\033[33m"
const cDefault = "\033[0m"

func asciiCode(k byte) int {
	return int([]rune(string(k))[0])
}

func getColumn(i int) int {
	if i == 'a' {
		return 0
	} else if i == '+' {
		return 1
	} else if i == '-' {
		return 2
	} else if i == '*' {
		return 3
	} else if i == '/' {
		return 4
	} else if i == '(' {
		return 5
	} else if i == ')' {
		return 6
	} else if i == '$' {
		return 7
	} else if i == '=' {
		return 8
	} else if i == 'b' {
		return 9
	}
	return -1
}

func stateNumToString(i int) string {
	if i > 32 && i < 127 {
		return string(asciiCode(byte(i)))
	}
	return strconv.Itoa(i)
}

func reject(i int, dat []byte) {
	fmt.Println()
	fmt.Println(cRed, "-----------------------------------------------------", cDefault)
	fmt.Println(cRed, "\n\n\tThe expression was rejected at location:", i, cDefault)
	fmt.Print("\t\t")
	for n, k := range dat {
		if n == i {
			fmt.Print(cRed, string(k), cDefault)
		} else {
			fmt.Print(string(k))
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(cRed, "-----------------------------------------------------", cDefault)
	fmt.Println()
	os.Exit(1)
}

func main() {
	// a   +   -   *   /   (    )   $  =    b
	var matrix = [][]int{
		{10, -1, -1, -1, -1, 10, -1, -1, 10, 10},     // E
		{-1, 11, 12, -1, -1, -1, '$', '$', -1, -1},   // Q
		{13, -1, -1, -1, -1, 13, -1, -1, -1, 13},     // T
		{-1, '$', '$', 14, 15, -1, '$', '$', -1, -1}, // R
		{'a', -1, -1, -1, -1, 16, -1, -1, -1, 'b'},   // F
		{17, -1, -1, -1, -1, -1, -1, -1, -1, 17}}     // S

	stateNum := -1
	stackOfBytes := stack.New()
	stackString := ""
	finalLocation := 0
	tabs := "\t\t\t"
	dat, _ := ioutil.ReadFile("./grammer.txt")
	fmt.Println("\nChecking the word: ", cGreen, string(dat), cDefault, "\n\n")
	stackOfBytes.Push(36) // push $
	stackOfBytes.Push(5)  // push S
	stackString += "$"
	stackString += "S"
	for i, k := range dat {
		//time.Sleep(time.Millisecond * 10)
		if len(stackString) > 3 {
			tabs = "\t\t"
		} else {
			tabs = "\t\t\t"
		}
		fmt.Print("StackString:", stackString)
		fmt.Print(tabs, "LastStateN:", stateNum)
		fmt.Println("\t\tk:", string(k))
		finalLocation = i
		if stackOfBytes.Len() != 0 {
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		}
		goto done

	skipIteration:
		fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString) // so we don't proceed to the next rune in the buffer
		//time.Sleep(time.Millisecond * 10)

	switchPlace:
		switch stateNum {
		case 0: // State E
			stateNum = matrix[0][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 1: // State Q
			stateNum = matrix[1][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 2: // State T
			stateNum = matrix[2][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 3: // State R
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = matrix[3][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 4: // State F
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = matrix[4][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 5: // State S
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = matrix[5][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
			}
			goto skipIteration
		case 'a': // StateNumber: 97
			if k != 97 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case 'b': // StateNumber: 98
			if k != 98 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '+': // StateNumber: 43
			if k != 43 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '-': // StateNumber: 45
			if k != 45 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '*': // StateNumber: 42
			if k != 42 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '/': // StateNumber: 47
			if k != 47 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '(': // StateNumber: 40
			if k != 40 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case ')': // StateNumber: 41
			if k != 41 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '=': // StateNumber: 61
			if k != 61 {
				reject(i, dat)
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case '$': // StateNumber: 36 (lambda)
			if k != 36 && stackOfBytes.Len() != 0 {
				stateNum = stackOfBytes.Pop().(int)
				stackString = stackString[0 : len(stackString)-1]
				goto skipIteration
			} else {

				if stateNum == '$' && i == len(dat)-1 && stackOfBytes.Len() == 0 {
					fmt.Println(cGreen, "\t\tWORD ACCEPTED", cDefault)
					goto done
				} else if stackOfBytes.Len() == 0 && k != 36 {
					reject(i, dat)

				} else {
					//fmt.Println("\t\tNO MATCH    k:", string(k), "stateNum: ", stateNum)
					stateNum = stackOfBytes.Pop().(int)
					stackString = stackString[0 : len(stackString)-1]
					fmt.Println("\t\tNO MATCH    k:", string(k), "stateNum: ", stateNum)

					goto skipIteration
				}
			}
			fmt.Println(cGreen, "\t\tMATCH:", string(k), " ✓ ", cDefault)
		case 10: //  TQ
			stackOfBytes.Push(1) // push Q
			stackOfBytes.Push(2) // push T
			stackString += "Q"
			stackString += "T"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 11: // +TQ
			stackOfBytes.Push(1)  // push Q
			stackOfBytes.Push(2)  // push T
			stackOfBytes.Push(43) // push +
			stackString += "Q"
			stackString += "T"
			stackString += "+"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 12: // -TQ
			stackOfBytes.Push(1)  // push Q
			stackOfBytes.Push(2)  // push T
			stackOfBytes.Push(45) // push -
			stackString += "Q"
			stackString += "T"
			stackString += "-"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 13: //  FR
			stackOfBytes.Push(3) // push R
			stackOfBytes.Push(4) // push F
			stackString += "R"
			stackString += "F"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 14: // *FR
			stackOfBytes.Push(3)  // push R
			stackOfBytes.Push(4)  // push F
			stackOfBytes.Push(42) // push *
			stackString += "R"
			stackString += "F"
			stackString += "*"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 15: // /FR
			stackOfBytes.Push(3)  // push R
			stackOfBytes.Push(4)  // push F
			stackOfBytes.Push(47) // push /
			stackString += "R"
			stackString += "F"
			stackString += "/"
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNum, "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 16: // (E)
			stackOfBytes.Push(41) // push )
			stackOfBytes.Push(0)  // push E
			stackOfBytes.Push(40) // push (
			stackString += ")"
			stackString += "E"
			stackString += "("
			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNumToString(stateNum), "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration
		case 17: // a=U
			stackOfBytes.Push(0)  // push E
			stackOfBytes.Push(61) // push =
			stackOfBytes.Push(97) // push a
			stackString += "E"
			stackString += "="
			stackString += "a"

			fmt.Println("    No Match Yet\t\t\t\t\tk:", string(k), "\t\t StateNum: ", stateNum, "\tStack: ", stackString)
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto skipIteration

		}
	}

done:
	if stackOfBytes.Len() == 0 {
		fmt.Println(cGreen, "This was a valid Expression", cDefault)
	} else {
		//fmt.Println(cRed, "NOT A VALID EXPRESSION", cDefault)
		reject(finalLocation, dat)
	}

}
