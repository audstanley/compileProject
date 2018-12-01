//Author: Kenny Chao
//Editor: Visual Studio Code
//Language used: Go
//Project No.7 (Predictive Parsing Table) Program 2
//Input strings to test: (1) a=(a+a)*b$	(2) a=a*(b-a)$	(3) a=(a+a)b$
//Each string is tested separately using the same program

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

//Input strings are commented out for separate testing
//var string1 = "a=(a+a)*b$"
//var string1 = "a=a*(b-a)$"
var string1 = "a=(a+a)b"

func main() {
	var multiD [6][9]string
	multiD = [6][9]string{
		//    	   a     b    +    -    *    /    (    )    $
		/* S */ {"a=E", "z", "z", "z", "z", "z", "z", "z", "z"},
		/* E */ {"TQ", "TQ", "z", "z", "z", "z", "TQ", "z", "z"},
		/* Q */ {"z", "z", "+TQ", "-TQ", "z", "z", "z", "l", "l"},
		/* T */ {"FR", "FR", "z", "z", "z", "z", "FR", "z", "z"},
		/* R */ {"z", "z", "l", "l", "*FR", "/FR", "z", "l", "l"},
		/* F */ {"a", "b", "z", "z", "z", "z", "(E)", "z", "z"},
	}

	//Test to see values of multi-dimentional array
	test := 0
	if test == 1 {
		fmt.Println(multiD[0][0])
	}

	//Creation of a new stack
	s := New()
	s.Push("$")
	s.Push("S")

	for s.Len() != 0 {
		var counter int
		goto SWITCHCASE
		//Incrementer is meant to increment the index of the input string after a terminal is accepted
	INCREMENTER:
		counter = counter + 1
		goto SWITCHCASE
		//Start of switch case to determine what characters to push or pop from the stack
		//If the popped value is non-terminal continue to push values onto the stack (if not null values)
		//When the switch case reaches a popped value that is terminal the algorithm checks if the value matches
		//the value in the string index, if correct continue else exit out of program
		//When a null value is reached the algorithm automatically exits out of the program and rejects the string
		//Also determines whether string is accepted or not
	SWITCHCASE:
		var index = string(string1[counter])
		p := s.Pop()
		fmt.Printf("Popped Value: %s\n", p)
		switch p {
		case "S":
			if index == "a" {
				fmt.Printf("[%s,%s] = a=E\n", p, index)
				fmt.Println("Pushed: E, a= onto the stack.")
				s.Push("E")
				s.Push("a=")
				goto SWITCHCASE
			} else {
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "E":
			if index == "a" {
				fmt.Printf("[%s,%s] = TQ\n", p, index)
				fmt.Println("Pushed: Q, T onto  the stack.")
				s.Push("Q")
				s.Push("T")
				goto SWITCHCASE
			} else if index == "b" {
				fmt.Printf("[%s,%s] = TQ\n", p, index)
				fmt.Println("Pushed: Q, T onto the stack.")
				s.Push("Q")
				s.Push("T")
				goto SWITCHCASE
			} else if index == "(" {
				fmt.Printf("[%s,%s] = TQ\n", p, index)
				fmt.Println("Pushed: Q, T onto the stack.")
				s.Push("Q")
				s.Push("T")
				goto SWITCHCASE
			} else {
				fmt.Printf("[%s,%s] = null value\n", p, index)
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "Q":
			if index == "+" {
				fmt.Printf("[%s,%s] = +TQ\n", p, index)
				fmt.Println("Pushed: Q, T, + onto the stack.")
				s.Push("Q")
				s.Push("T")
				s.Push("+")
				goto SWITCHCASE
			} else if index == "-" {
				fmt.Printf("[%s,%s] = -TQ\n", p, index)
				fmt.Println("Pushed: Q, T, - onto the stack")
				s.Push("Q")
				s.Push("T")
				s.Push("-")
				goto SWITCHCASE
			} else if index == ")" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else if index == "$" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else {
				fmt.Printf("[%s,%s] = null value\n", p, index)
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "T":
			if index == "a" {
				fmt.Printf("[%s,%s] = FR\n", p, index)
				fmt.Println("Pushed: R, F onto the stack.")
				s.Push("R")
				s.Push("F")
				goto SWITCHCASE
			} else if index == "b" {
				fmt.Printf("[%s,%s] = FR\n", p, index)
				fmt.Println("Pushed: R, F onto the stack.")
				s.Push("R")
				s.Push("F")
				goto SWITCHCASE
			} else if index == "(" {
				fmt.Printf("[%s,%s] = FR\n", p, index)
				fmt.Println("Pushed: R, F onto the stack.")
				s.Push("R")
				s.Push("F")
				goto SWITCHCASE
			} else {
				fmt.Printf("[%s,%s] = null value\n", p, index)
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "R":
			if index == "+" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else if index == "-" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else if index == ")" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else if index == "$" {
				fmt.Printf("[%s,%s] = lambda\n", p, index)
				fmt.Println("Pushed: lambda onto the stack.")
				goto SWITCHCASE
			} else if index == "*" {
				fmt.Printf("[%s,%s] = *FR\n", p, index)
				fmt.Println("Pushed: R, F, * onto the stack.")
				s.Push("R")
				s.Push("F")
				s.Push("*")
				goto SWITCHCASE
			} else if index == "/" {
				fmt.Printf("[%s,%s] = /FR\n", p, index)
				fmt.Println("Pushed: R, F, / onto the stack.")
				s.Push("R")
				s.Push("F")
				s.Push("/")
				goto SWITCHCASE
			} else {
				fmt.Printf("[%s,%s] = null value\n", p, index)
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "F":
			if index == "a" {
				fmt.Printf("[%s,%s] = a\n", p, index)
				fmt.Println("Pushed: a onto the stack.")
				s.Push("a")
				goto SWITCHCASE
			} else if index == "b" {
				fmt.Printf("[%s,%s] = b\n", p, index)
				fmt.Println("Pushed: b onto the stack.")
				s.Push("b")
				goto SWITCHCASE
			} else if index == "(" {
				fmt.Printf("[%s,%s] = (E)\n", p, index)
				fmt.Println("Pushed: ), E, ( onto the stack.")
				s.Push(")")
				s.Push("E")
				s.Push("(")
				goto SWITCHCASE
			} else {
				fmt.Printf("[%s,%s] = null value\n", p, index)
				fmt.Println("Error: null value found!")
				goto EXIT
			}
		case "a=":
			if string(string1[counter]) == "a" && string(string1[counter+1]) == "=" {
				var index2 = "a="
				if p == index2 {
					var index3 = string(string1[counter])
					var index4 = string(string1[counter+1])
					fmt.Printf("%s matches with %s%s\n", p, index3, index4)
					counter = counter + 1
					goto INCREMENTER
				} else {
					fmt.Printf("%s does not match with %s\n", p, index)
					goto EXIT
				}
			}
		case "a":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "b":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "+":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "-":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "*":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "/":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "(":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case ")":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				goto INCREMENTER
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
		case "$":
			if p == index {
				fmt.Printf("%s matches with %s\n", p, index)
				fmt.Printf("The string '%s' is accepted.\n", string1)
			} else {
				fmt.Printf("%s does not match with %s\n", p, index)
				goto EXIT
			}
			os.Exit(0)
		default:
			goto EXIT
		}
	EXIT:
		fmt.Printf("The string '%s' is not accepted.\n", string1)
		os.Exit(0)
	}
}
