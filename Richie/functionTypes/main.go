package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type operation struct {
	s    string
	x, y int
}

// generic interface for operation type
//type T interface {
//	runOperation()
//}

// The run function is implecitely called from
// an operation struct, runs the operation
// and returns a struct.
func (o operation) runOperation() int {
	if o.s == "+" {
		return o.x + o.y
	} else if o.s == "-" {
		return o.x - o.y
	} else if o.s == "*" {
		return o.x * o.y
	} else if o.s == "/" {
		if o.y == 0 {
			panic("You tried to divide by zero.")
		}
		return o.x / o.y
	}
	return 0
}

type conditionalOperation struct {
	leftSide  string
	rightSide int
	operator  string
}

type forLoop struct {
	counterName        string
	countValue         int
	condition          conditionalOperation
	forIncDecOperation incOrDec
	body               string
}

type incOrDec struct {
	leftside  string
	rightSide int
	operator  string
}

func makeAForLoop(params, body string) *forLoop {
	sliceOfForLoopParameters := strings.Split(params, ";")
	sliceOfFirstCondition := strings.Split(sliceOfForLoopParameters[0], " ")
	sliceOfConditionalOperation := strings.Split(sliceOfForLoopParameters[1], " ")
	sliceOfIncOrDecOperation := strings.Split(sliceOfForLoopParameters[1], " ")
	// slice up the parameters of the for loop
	newForLoop := new(forLoop)
	newForLoop.counterName = sliceOfFirstCondition[0]
	newForLoop.countValue, _ = strconv.Atoi(sliceOfFirstCondition[2]) // ignoring errors for now.
	newForLoop.condition.leftSide = sliceOfConditionalOperation[0]
	newForLoop.condition.operator = sliceOfConditionalOperation[1]
	newForLoop.condition.rightSide, _ = strconv.Atoi(sliceOfConditionalOperation[2])
	newForLoop.forIncDecOperation.leftside = sliceOfIncOrDecOperation[0]
	newForLoop.forIncDecOperation.rightSide, _ = strconv.Atoi(sliceOfIncOrDecOperation[2])
	newForLoop.forIncDecOperation.operator = sliceOfIncOrDecOperation[1]
	newForLoop.body = body
	return newForLoop

}

type function struct {
	fType           string
	hashName        string
	mapOfParameters map[string]fParameters
	mapOfOperations map[string]operation
}

type fParameters struct {
	location int
	pName    string
	pType    string
}

var functionReturnTypes = make(map[string]function)

func isAValidFunction(s string) bool {
	return true
}

func main() {
	//test := function{
	//	"void",
	//	"awdbwaidubw",
	//	map[string]fParameters{"awdbwaidubw": {0, "str", "string"}},
	//	map[string]operation{"awdbwaidubw": {"+", 3, 4}}}
	multiLineStr := "for(int i = 0; i < 10; i++){\n  someVar = someVal;\n  moreOperations++;\n  someFunctionCall(2);\n}"
	forLoopRegex := regexp.MustCompile(`for\s*\(\s*int\s+([\w]+)\s*(=)\s*([\w]+)\s*;\s*([\w]+)\s*([<>]|[=]|[<=>=])\s*(\w+);\s*([\w+-=<>]+)\s*\)\s*\{\s*[\r\n]+([^}]+)}`)
	fl := forLoopRegex.FindStringSubmatch(multiLineStr)
	countVal, _ := strconv.Atoi(fl[3])
	Cond, _ := strconv.Atoi(fl[6])
	aCoolNewForLoopStruct := forLoop{fl[1], countVal, conditionalOperation{fl[4], upperCond, fl[5]}, incOrDec{"i", 10, "+="}, fl[8]}

	fmt.Println(fl)
	fmt.Printf("%+v\n", aCoolNewForLoopStruct)

}
