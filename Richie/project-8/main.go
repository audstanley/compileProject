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
const cBlue = "\033[34m"
const cMagenta = "\033[35m"
const cDefault = "\033[0m"

//   i    +    -    *    /    (    )    $    E    T    F
var matrix = [][]int{
	{405, 700, 700, 700, 700, 404, 700, 700, 501, 502, 503}, // 0
	{700, 406, 407, 700, 700, 700, 700, 600, 700, 700, 700}, // 1
	{700, 303, 303, 408, 409, 700, 303, 303, 700, 700, 700}, // 2
	{700, 306, 306, 306, 306, 700, 306, 306, 700, 700, 700}, // 3
	{405, 700, 700, 700, 700, 404, 700, 700, 510, 502, 503}, // 4
	{700, 308, 308, 308, 308, 700, 308, 308, 700, 700, 700}, // 5
	{405, 700, 700, 700, 700, 404, 700, 700, 700, 511, 503}, // 6
	{405, 700, 700, 700, 700, 404, 700, 700, 700, 512, 503}, // 7
	{405, 700, 700, 700, 700, 404, 700, 700, 700, 700, 514}, // 8
	{405, 700, 700, 700, 700, 404, 700, 700, 700, 700, 700}, // 9
	{700, 406, 407, 700, 700, 700, 415, 700, 700, 700, 700}, // 10
	{700, 301, 301, 408, 409, 700, 301, 301, 700, 700, 700}, // 11
	{700, 302, 302, 408, 409, 700, 302, 302, 700, 700, 700}, // 12
	{700, 304, 304, 304, 304, 700, 304, 304, 700, 700, 700}, // 13
	{700, 305, 305, 305, 305, 700, 305, 305, 700, 700, 700}, // 14
	{700, 307, 307, 307, 307, 700, 307, 307, 700, 700, 700}} // 15

func asciiCode(k byte) int {
	return int([]rune(string(k))[0])
}

func getRow(i int) int {
	return i / 100
}

func popTimes(nTimes int, intSt *stack.Stack, stString *string) {
	for i := 0; i < nTimes; i++ {
		fmt.Println(cRed, "    POP!, ", (*stString)[len(*stString)-4:len(*stString)-1], cDefault)
		intSt.Pop()
		if len(*stString) > 3 {
			*stString = (*stString)[0 : len(*stString)-4]
		} else {
			*stString = ""
		}
	}
}

func printStateNum(i int) string {
	if i/100 == 3 {
		return "R" + strconv.Itoa(i%300) + "  "
	} else if i/100 == 4 {
		return "S" + strconv.Itoa(i%400) + "  "
	} else if i/100 == 5 {
		if len(strconv.Itoa(i%500)) == 1 {
			return "N" + strconv.Itoa(i%500) + "  "
		} else {
			return "N" + strconv.Itoa(i%500) + " "
		}

	} else if i > 32 && i < 127 {
		return string(asciiCode(byte(i))) + "   "
	}
	if len(strconv.Itoa(i)) == 1 {
		return strconv.Itoa(i) + "   "
	} else if len(strconv.Itoa(i)) == 2 {
		return strconv.Itoa(i) + "  "
	} else if len(strconv.Itoa(i)) == 3 {
		return strconv.Itoa(i) + " "
	}
	return strconv.Itoa(i) + " "

}

func getColumn(i int) int {
	if i == 'i' {
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
	} else if i == 'E' {
		return 8
	} else if i == 'T' {
		return 9
	} else if i == 'F' {
		return 10
	}
	return -1
}

func lineNumFormula(lineNum int) int {
	switch lineNum {
	case 1:
		return asciiCode('E')
	case 2:
		return asciiCode('E')
	case 3:
		return asciiCode('E')
	case 4:
		return asciiCode('T')
	case 5:
		return asciiCode('T')
	case 6:
		return asciiCode('T')
	case 7:
		return asciiCode('F')
	case 8:
		return asciiCode('F')
	}
	return -1
}

func rLengthFormula(rNum int) int {
	switch rNum {
	case 1:
		return 2 * len("E+T")
	case 2:
		return 2 * len("E-T")
	case 3:
		return 2 * len("T")
	case 4:
		return 2 * len("T*F")
	case 5:
		return 2 * len("T/F")
	case 6:
		return 2 * len("F")
	case 7:
		return 2 * len("(E)")
	case 8:
		return 2 * len("i")
	}
	return 0
}

func getRHSoFS(rhs int, color string) int {
	fmt.Println(color, "Before RHS: ", rhs, cDefault)
	if getRow(rhs) == 3 {
		return rhs % 300
	} else if getRow(rhs) == 4 {
		return rhs % 400
	} else if getRow(rhs) == 5 {
		return rhs % 500
	} else if getRow(rhs) == 6 {
		return rhs % 600
	} else if getRow(rhs) == 7 {
		return rhs % 700
	}
	return rhs % 200
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

	stackOfBytes := stack.New()
	stackString := ""
	dat, _ := ioutil.ReadFile("./grammer.txt")
	fmt.Println("\nChecking the word: ", cGreen, string(dat), cDefault, "\n\n")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\n\n        Lines in:")
	fmt.Println(cYellow, "           YELLOW:  S-STATES ", cDefault)
	fmt.Println(cRed, "              RED:  R-STATES ", cDefault)
	fmt.Println(cBlue, "      NonTerminal: NT-STATES \n\n", cDefault)
	fmt.Println("--------------------------------------------------------------")
	stackOfBytes.Push(0) // push 0
	stateNum := 0
	stackString += "....    0   "

	for _, k := range dat {

	switchJump:

		fmt.Println("    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString)

		if getRow(stateNum) == 0 { // ascii code state
			fmt.Println("IN THE ASCII-STATE")
			switch stateNum {
			case 0:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 1:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 2:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 3:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 4:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 5:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 6:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 7:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			case 8:
				lastStateNum := stateNum
				stateNum = matrix[stackOfBytes.Pop().(int)][getColumn(asciiCode(k))]
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				fmt.Println("    Going to\t\t", printStateNum(stateNum))
				stackOfBytes.Push(lastStateNum % 200)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, ""))
				stackString += printStateNum(lastStateNum / 200)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, ""))

			}
		} else if getRow(stateNum) == 3 { // R states ///////////////////////////////////////////////////////////////////////////////////
			fmt.Println(cRed, "IN THE R-STATE", cDefault)
			switch getRow(stateNum) % 300 {
			case 1:
				fmt.Println(cRed, "    BEFORE R1   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(getRow(stateNum)%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 2:
				fmt.Println(cRed, "    BEFORE R2   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(stateNum%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 3:
				fmt.Println(cRed, "    BEFORE R3   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(stateNum%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				stackString += printStateNum(row)
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(stateNum)

				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 4:
				fmt.Println(cRed, "    BEFORE R4   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(stateNum%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 5:
				fmt.Println(cRed, "    BEFORE R5   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(stateNum%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 6:
				fmt.Println(cRed, "    BEFORE R6   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(stateNum%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 7:
				fmt.Println(cRed, "    BEFORE R7   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(getRow(stateNum)%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			case 8:
				fmt.Println(cRed, "    BEFORE R8   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := lineNumFormula(stateNum % 300)
				fmt.Println(cRed, "   Last StateNum:", stateNumToString(lastStateNum), cDefault)
				popTimes(rLengthFormula(getRow(stateNum)%300), stackOfBytes, &stackString)
				row := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[row][getColumn(lastStateNum)]
				stackOfBytes.Push(row)
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(stateNum)
				fmt.Println(cRed, "    PUSH!, ", row, "\n    PUSH!, ", printStateNum(lastStateNum), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cRed, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				goto switchJump
			}

		} else if getRow(stateNum) == 4 { // S states //////////////////////////////////////////////////////////////////////////////////
			fmt.Println(cYellow, "IN THE S-STATE", cDefault)
			switch stateNum % 400 {
			case 4:
				fmt.Println(cYellow, "    BEFORE S4   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 { // THIS MIGHT NEED TO BE DUPLICATED IN THE S STATES!!
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 5:
				fmt.Println(cYellow, "    BEFORE S5   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 { // THIS MIGHT NEED TO BE DUPLICATED IN THE S STATES!!
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 6:
				fmt.Println(cYellow, "    BEFORE S6   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 { // THIS MIGHT NEED TO BE DUPLICATED IN THE S STATES!!
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 7:
				fmt.Println(cYellow, "    BEFORE S7   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 8:
				fmt.Println(cYellow, "    BEFORE S8  \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 9:
				fmt.Println(cYellow, "    BEFORE S9   \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 15:
				fmt.Println(cYellow, "    BEFORE S15  \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				popped := stackOfBytes.Pop().(int)
				fmt.Println(cYellow, "JUST POPPED: ", popped)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[popped][getColumn(asciiCode(k))]
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(popped)
					stackString += printStateNum(popped)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				fmt.Println("    stNum\t\t", stateNum)
				stackOfBytes.Push(popped)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cYellow))
				stackString += printStateNum(popped)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cYellow))
				fmt.Println(cYellow, "    PUSH!, ", (popped), "\n    PUSH!, ", printStateNum(asciiCode(k)), "\n    PUSH!, ", stateNum, cDefault)
				fmt.Println(cYellow, "    No Match Yet\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			}

		} else if getRow(stateNum) == 5 { // Non-Terminal States
			fmt.Println(cBlue, "IN THE NonTerminal-STATE", cDefault)
			switch stateNum % 500 {
			case 1:
				fmt.Println(cBlue, "    BEFORE NT-1 \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}

				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)

				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}

				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))

				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 2:
				fmt.Println(cBlue, "    BEFORE NT-2 \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}

				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)

				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}

				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))

				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 3:
				fmt.Println(cBlue, "    BEFORE NT-3 \t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}

				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))

				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 10: ////////////////////////////////////////// BOOP
				fmt.Println(cBlue, "    BEFORE NT-10\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))

				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 11:
				fmt.Println(cBlue, "    BEFORE NT-11\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))
				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 12:
				fmt.Println(cBlue, "    BEFORE NT-12\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))
				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 13:
				fmt.Println(cBlue, "    BEFORE NT-13\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))
				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			case 14:
				fmt.Println(cBlue, "    BEFORE NT-14\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				lastStateNum := stateNum
				popped := stackOfBytes.Pop().(int)
				if len(stackString) > 3 {
					stackString = stackString[0 : len(stackString)-4]
				} else {
					stackString = ""
				}
				stateNum = matrix[getRHSoFS(popped, cBlue)][getColumn(asciiCode(k))]
				fmt.Println(cBlue, "   Last StateNum: [", getRHSoFS(popped, cBlue), ",", string(k), "] = ", stateNum, cDefault)
				fmt.Println(cBlue, "   New StateNum:", stateNumToString(lastStateNum), cDefault)
				fmt.Println(cBlue, "    BEFORE POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)
				if stateNum < 400 && stateNum > 300 {
					stackOfBytes.Push(lastStateNum)
					stackString += printStateNum(lastStateNum)
					popTimes(rLengthFormula(stateNum), stackOfBytes, &stackString)
					goto switchJump
				} else if stateNum == 600 {
					goto switchJump
				}
				stackOfBytes.Push(lastStateNum)
				stackOfBytes.Push(asciiCode(k))
				stackOfBytes.Push(getRHSoFS(stateNum, cBlue))
				stackString += printStateNum(lastStateNum)
				stackString += printStateNum(asciiCode(k))
				stackString += printStateNum(getRHSoFS(stateNum, cBlue))
				fmt.Println(cBlue, "    AFTER POPS\t\t\t\t\tr:", getRow(stateNum), "\t\tcol: ", string(k), "\t\t StateNumStr: ", printStateNum(stateNum), "\t\t StateNumRaw: ", stateNum, "\tStack: ", stackString, cDefault)

			}
		} else if getRow(stateNum) == 6 { // accept state
			fmt.Println(cGreen, "IN THE ACCEPT-STATE", cDefault)
			switch stateNum % 6 {
			case 0:
				fmt.Println(cGreen, "    WORD ACCEPTED!", cDefault)
			}
		} else if getRow(stateNum) == 7 { // reject state
			fmt.Println(cMagenta, "IN THE REJECT-STATE", cDefault)
			switch stateNum % 7 {
			case 0:
				fmt.Println(cMagenta, "    THE WORD WAS REJECTED", cDefault)
			}
		}
	}
}
