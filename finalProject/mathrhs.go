package main

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

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
	} else if i == ';' {
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
}

func mathrhs(expression string) int {
	// a   +   -   *   /   (    )   ;  =    b
	var matrix = [][]int{
		{10, -1, -1, -1, -1, 10, -1, -1, 10, 10},     // E
		{-1, 11, 12, -1, -1, -1, ';', ';', -1, -1},   // Q
		{13, -1, -1, -1, -1, 13, -1, -1, -1, 13},     // T
		{-1, ';', ';', 14, 15, -1, ';', ';', -1, -1}, // R
		{'a', -1, -1, -1, -1, 16, -1, -1, -1, 'b'},   // F
		{17, -1, -1, -1, -1, -1, -1, -1, -1, 17}}     // S

	stateNum := -1
	stackOfBytes := stack.New()
	stackString := ""
	finalLocation := 0
	dat := []byte(expression)
	stackOfBytes.Push(59) // push ;
	stackOfBytes.Push(5)  // push S
	stackString += ";"
	stackString += "S"
	for i, k := range dat {
		finalLocation = i
		if stackOfBytes.Len() != 0 {
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		}
		goto done

	switchPlace:
		switch stateNum {
		case 0: // State E
			stateNum = matrix[0][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 1: // State Q
			stateNum = matrix[1][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 2: // State T
			stateNum = matrix[2][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 3: // State R
			stateNum = matrix[3][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 4: // State F
			stateNum = matrix[4][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 5: // State S
			stateNum = matrix[5][getColumn(asciiCode(k))]
			if stateNum == -1 {
				reject(i, dat)
				return i
			}
			goto switchPlace
		case 'a': // StateNumber: 97
			if k != 97 {
				reject(i, dat)
				return i
			}
		case 'b': // StateNumber: 98
			if k != 98 {
				reject(i, dat)
				return i
			}
		case '+': // StateNumber: 43
			if k != 43 {
				reject(i, dat)
				return i
			}
		case '-': // StateNumber: 45
			if k != 45 {
				return i
			}
		case '*': // StateNumber: 42
			if k != 42 {
				reject(i, dat)
				return i
			}
		case '/': // StateNumber: 47
			if k != 47 {
				reject(i, dat)
				return i
			}
		case '(': // StateNumber: 40
			if k != 40 {
				reject(i, dat)
				return i
			}
		case ')': // StateNumber: 41
			if k != 41 {
				reject(i, dat)
				return i
			}
		case '=': // StateNumber: 61
			if k != 61 {
				reject(i, dat)
				return i
			}
		case ';': // StateNumber: 59 (lambda)
			if k != 59 && stackOfBytes.Len() != 0 {
				stateNum = stackOfBytes.Pop().(int)
				stackString = stackString[0 : len(stackString)-1]
				goto switchPlace
			} else {
				if stateNum == ';' && i == len(dat)-1 && stackOfBytes.Len() == 0 {
					goto done
				} else if stackOfBytes.Len() == 0 && k != 59 {
					reject(i, dat)
					return i
				} else {
					stateNum = stackOfBytes.Pop().(int)
					stackString = stackString[0 : len(stackString)-1]
					goto switchPlace
				}
			}
		case 10: //  TQ
			stackOfBytes.Push(1) // push Q
			stackOfBytes.Push(2) // push T
			stackString += "Q"
			stackString += "T"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 11: // +TQ
			stackOfBytes.Push(1)  // push Q
			stackOfBytes.Push(2)  // push T
			stackOfBytes.Push(43) // push +
			stackString += "Q"
			stackString += "T"
			stackString += "+"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 12: // -TQ
			stackOfBytes.Push(1)  // push Q
			stackOfBytes.Push(2)  // push T
			stackOfBytes.Push(45) // push -
			stackString += "Q"
			stackString += "T"
			stackString += "-"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 13: //  FR
			stackOfBytes.Push(3) // push R
			stackOfBytes.Push(4) // push F
			stackString += "R"
			stackString += "F"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 14: // *FR
			stackOfBytes.Push(3)  // push R
			stackOfBytes.Push(4)  // push F
			stackOfBytes.Push(42) // push *
			stackString += "R"
			stackString += "F"
			stackString += "*"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 15: // /FR
			stackOfBytes.Push(3)  // push R
			stackOfBytes.Push(4)  // push F
			stackOfBytes.Push(47) // push /
			stackString += "R"
			stackString += "F"
			stackString += "/"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 16: // (E)
			stackOfBytes.Push(41) // push )
			stackOfBytes.Push(0)  // push E
			stackOfBytes.Push(40) // push (
			stackString += ")"
			stackString += "E"
			stackString += "("
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		case 17: // a=U
			stackOfBytes.Push(0)  // push E
			stackOfBytes.Push(61) // push =
			stackOfBytes.Push(97) // push a
			stackString += "E"
			stackString += "="
			stackString += "a"
			stateNum = stackOfBytes.Pop().(int)
			stackString = stackString[0 : len(stackString)-1]
			goto switchPlace
		}
	}

done:
	if stackOfBytes.Len() == 0 {
		return -1
	} else {
		reject(finalLocation, dat)
		return finalLocation
	}

}
