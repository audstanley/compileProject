//Author: Kenny Chao
//Editor: Visual Studio Code
//Language: Go
//Project No.8 LR Parsing Table Method 2
//Input strings to test (1) (i+i)*$		(2) (i*)$
//Each input string is tested separately using the same program

package main

import (
	"fmt"
	"os"
)

//Stack ...
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

//New ...
func New() *Stack {
	return &Stack{nil, 0}
}

//Len ...
func (receiver *Stack) Len() int {
	return receiver.length
}

//Peek ...
func (receiver *Stack) Peek() interface{} {
	if receiver.length == 0 {
		return nil
	}
	return receiver.top.value
}

//Pop ...
func (receiver *Stack) Pop() interface{} {
	if receiver.length == 0 {
		return nil
	}

	n := receiver.top
	receiver.top = n.prev
	receiver.length--
	return n.value
}

//Push ...
func (receiver *Stack) Push(value interface{}) {
	n := &node{value, receiver.top}
	receiver.top = n
	receiver.length++
}

var input1 = "(i+i)*i$"

//var input1 = "(i*)$"

func main() {

	var multiD [16][11]string
	multiD = [16][11]string{
		//		    i   +   -   *    /   (    )    $   E    T    F
		/* 0 */ {"S5", "", "", "", "", "S4", "", "", "1", "2", "3"},
		/* 1 */ {"", "S6", "S7", "", "", "", "", "acc", "", "", ""},
		/* 2 */ {"", "R3", "R3", "S8", "S9", "", "R3", "R3", "", "", ""},
		/* 3 */ {"", "R6", "R6", "R6", "R6", "", "R6", "R6", "", "", ""},
		/* 4 */ {"S5", "", "", "", "", "S4", "", "", "10", "2", "3"},
		/* 5 */ {"", "R8", "R8", "R8", "R8", "", "R8", "R8", "", "", ""},
		/* 6 */ {"S5", "", "", "", "", "S4", "", "", "", "11", "3"},
		/* 7 */ {"S5", "", "", "", "", "S4", "", "", "", "12", "3"},
		/* 8 */ {"S5", "", "", "", "", "S4", "", "", "", "", "13"},
		/* 9 */ {"S5", "", "", "", "", "S4", "", "", "", "", "14"},
		/* 10 */ {"", "S6", "S7", "", "", "", "S15", "", "", "", ""},
		/* 11 */ {"", "R1", "R1", "S8", "S9", "", "R1", "R2", "", "", ""},
		/* 12 */ {"", "R2", "R2", "S8", "S9", "", "R2", "R2", "", "", ""},
		/* 13 */ {"", "R4", "R4", "R4", "R4", "", "R4", "R4", "", "", ""},
		/* 14 */ {"", "R5", "R5", "R5", "R5", "", "R5", "R5", "", "", ""},
		/* 15 */ {"", "R7", "R7", "R7", "R7", "", "R7", "R7", "", "", ""},
	}

	s := New()
	s.Push("0")
	for s.Len() != 0 {
		var counter int
		var mdValue string
		var readState string
		goto STATE_SWITCHCASE
	INCREMENTER:
		counter = counter + 1
		goto STATE_SWITCHCASE
	STATE_SWITCHCASE:
		var index = string(input1[counter])
		p := s.Pop()
		fmt.Printf("Popped %s off the stack.\n", p)
		switch p {
		case "0":
			if index == "i" {
				mdValue = multiD[0][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[0][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "1":
			if index == "+" {
				mdValue = multiD[1][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "6")
				s.Push(p)
				s.Push(index)
				s.Push("6")
				goto INCREMENTER
			} else if index == "-" {
				mdValue = multiD[1][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "7")
				s.Push(p)
				s.Push(index)
				s.Push("7")
				goto INCREMENTER
			} else if index == "$" {
				mdValue = multiD[1][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				if mdValue == "acc" {
					fmt.Printf("The string %s is accepted.\n", input1) //When index is $ and popped value is 1 then accept string
					os.Exit(0)
				}
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "2":
			if index == "+" {
				mdValue = multiD[2][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R3"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[2][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R3"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[2][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "8")
				s.Push(p)
				s.Push(index)
				s.Push("8")
				goto INCREMENTER
			} else if index == "/" {
				mdValue = multiD[2][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "9")
				s.Push(p)
				s.Push(index)
				s.Push("9")
				goto INCREMENTER
			} else if index == ")" {
				mdValue = multiD[2][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R3"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[2][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R3"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "3":
			if index == "+" {
				mdValue = multiD[3][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[3][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[3][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "/" {
				mdValue = multiD[3][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == ")" {
				mdValue = multiD[3][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[3][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R6"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "4":
			if index == "i" {
				mdValue = multiD[4][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[4][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "5":
			if index == "+" {
				mdValue = multiD[5][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[5][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[5][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "/" {
				mdValue = multiD[5][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == ")" {
				mdValue = multiD[5][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[5][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R8"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "6":
			if index == "i" {
				mdValue = multiD[6][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[6][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "7":
			if index == "i" {
				mdValue = multiD[7][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[7][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "8":
			if index == "i" {
				mdValue = multiD[8][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[8][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "9":
			if index == "i" {
				mdValue = multiD[9][0]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "5")
				s.Push(p)
				s.Push(index)
				s.Push("5")
				goto INCREMENTER
			} else if index == "(" {
				mdValue = multiD[9][5]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "4")
				s.Push(p)
				s.Push(index)
				s.Push("4")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "10":
			if index == "+" {
				mdValue = multiD[10][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "6")
				s.Push(p)
				s.Push(index)
				s.Push("6")
				goto INCREMENTER
			} else if index == "-" {
				mdValue = multiD[10][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "7")
				s.Push(p)
				s.Push(index)
				s.Push("7")
				goto INCREMENTER
			} else if index == ")" {
				mdValue = multiD[10][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "15")
				s.Push(p)
				s.Push(index)
				s.Push("15")
				goto INCREMENTER
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "11":
			if index == "+" {
				mdValue = multiD[11][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R1"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[11][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R1"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[11][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "8")
				s.Push(p)
				s.Push(index)
				s.Push("8")
				goto INCREMENTER
			} else if index == "/" {
				mdValue = multiD[11][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "9")
				s.Push(p)
				s.Push(index)
				s.Push("9")
				goto INCREMENTER
			} else if index == ")" {
				mdValue = multiD[11][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R1"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[11][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R1"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "12":
			if index == "+" {
				mdValue = multiD[12][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R2"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[12][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R2"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[12][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "8")
				s.Push(p)
				s.Push(index)
				s.Push("8")
				goto INCREMENTER
			} else if index == "/" {
				mdValue = multiD[12][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s, %s, and %s onto  the stack.\n", p, index, "9")
				s.Push(p)
				s.Push(index)
				s.Push("9")
				goto INCREMENTER
			} else if index == ")" {
				mdValue = multiD[12][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R2"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[12][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R2"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "13":
			if index == "+" {
				mdValue = multiD[13][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[13][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[13][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "/" {
				mdValue = multiD[13][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == ")" {
				mdValue = multiD[13][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[13][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R4"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "14":
			if index == "+" {
				mdValue = multiD[14][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[14][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[14][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "/" {
				mdValue = multiD[14][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == ")" {
				mdValue = multiD[14][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[14][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R5"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		case "15":
			if index == "+" {
				mdValue = multiD[15][1]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "-" {
				mdValue = multiD[15][2]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "*" {
				mdValue = multiD[15][3]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "/" {
				mdValue = multiD[15][4]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == ")" {
				mdValue = multiD[15][6]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else if index == "$" {
				mdValue = multiD[15][7]
				fmt.Printf("Read in: %s\n", index)
				fmt.Printf("[%s,%s] = %s\n", p, index, mdValue)
				fmt.Printf("Pushed %s onto  the stack.\n", p)
				s.Push(p)
				readState = "R7"
				READSTATE(s, readState)
				goto STATE_SWITCHCASE
			} else {
				fmt.Println("Error: empty value in cell.")
				goto EXIT
			}
		default:
			goto EXIT
		}
	EXIT:
		fmt.Printf("The string %s is not accepted.", input1)
		os.Exit(0)
	}
}

//READSTATE function is for the read state
func READSTATE(s *Stack, readState string) {
	var leftState string
	switch readState {
	case "R1":
		fmt.Println("(1) E -> E + T")
		fmt.Println("Pop 2 * |E + T| = 6")
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		leftState = "E"
		LEFTSTATE(s, leftState)
	case "R2":
		fmt.Println("(2) E -> E - T")
		fmt.Println("Pop 2 * |E - T| = 6")
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		leftState = "E"
		LEFTSTATE(s, leftState)
	case "R3":
		fmt.Println("(3) E -> T")
		fmt.Println("Pop 2 * |T| = 2")
		s.Pop()
		s.Pop()
		leftState = "E"
		LEFTSTATE(s, leftState)
	case "R4":
		fmt.Println("(4) T -> T * F")
		fmt.Println("Pop 2 * |T * F| = 6")
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		leftState = "T"
		LEFTSTATE(s, leftState)
	case "R5":
		fmt.Println("(5) T -> T / F")
		fmt.Println("Pop 2 * |T / F| = 6")
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		leftState = "T"
		LEFTSTATE(s, leftState)
	case "R6":
		fmt.Println("(6) T -> F")
		fmt.Println("Pop 2 * |F| = 2")
		s.Pop()
		s.Pop()
		leftState = "T"
		LEFTSTATE(s, leftState)
	case "R7":
		fmt.Println("(7) F -> ( E )")
		fmt.Println("Pop 2 * |( E )| = 6")
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		leftState = "F"
		LEFTSTATE(s, leftState)
	case "R8":
		fmt.Println("(7) F -> i")
		fmt.Println("Pop 2 * |i| = 2")
		s.Pop()
		s.Pop()
		leftState = "F"
		LEFTSTATE(s, leftState)
	default:
		fmt.Printf("The string %s is not accepted.", input1)
		os.Exit(0)
	}
}

//LEFTSTATE function is for the left most state
func LEFTSTATE(s *Stack, leftState string) {
	p := s.Pop()
	fmt.Printf("Popped %s off the stack.\n", p)
	switch leftState {
	case "E":
		if p == "0" {
			fmt.Printf("[%s,%s] = 1\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "1")
			s.Push(p)
			s.Push(leftState)
			s.Push("1")
		} else if p == "4" {
			fmt.Printf("[%s,%s] = 10\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "10")
			s.Push(p)
			s.Push(leftState)
			s.Push("10")
		} else {
			fmt.Printf("The string %s is not accepted.", input1)
			os.Exit(0)
		}
	case "T":
		if p == "0" {
			fmt.Printf("[%s,%s] = 2\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "2")
			s.Push(p)
			s.Push(leftState)
			s.Push("2")
		} else if p == "4" {
			fmt.Printf("[%s,%s] = 2\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "2")
			s.Push(p)
			s.Push(leftState)
			s.Push("2")
		} else if p == "6" {
			fmt.Printf("[%s,%s] = 11\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "11")
			s.Push(p)
			s.Push(leftState)
			s.Push("11")
		} else if p == "7" {
			fmt.Printf("[%s,%s] = 12\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "12")
			s.Push(p)
			s.Push(leftState)
			s.Push("12")
		} else {
			fmt.Printf("The string %s is not accepted.", input1)
			os.Exit(0)
		}
	case "F":
		if p == "0" {
			fmt.Printf("[%s,%s] = 3\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "3")
			s.Push(p)
			s.Push(leftState)
			s.Push("3")
		} else if p == "4" {
			fmt.Printf("[%s,%s] = 3\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "3")
			s.Push(p)
			s.Push(leftState)
			s.Push("3")
		} else if p == "6" {
			fmt.Printf("[%s,%s] = 3\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "3")
			s.Push(p)
			s.Push(leftState)
			s.Push("3")
		} else if p == "7" {
			fmt.Printf("[%s,%s] = 3\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "3")
			s.Push(p)
			s.Push(leftState)
			s.Push("3")
		} else if p == "8" {
			fmt.Printf("[%s,%s] = 13\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "13")
			s.Push(p)
			s.Push(leftState)
			s.Push("13")
		} else if p == "9" {
			fmt.Printf("[%s,%s] = 14\n", p, leftState)
			fmt.Printf("Pushed %s, %s, and %s onto the stack.\n", p, leftState, "14")
			s.Push(p)
			s.Push(leftState)
			s.Push("14")
		} else {
			fmt.Printf("The string %s is not accepted.", input1)
			os.Exit(0)
		}
	default:
		fmt.Printf("The string %s is not accepted.", input1)
		os.Exit(0)
	}
}
