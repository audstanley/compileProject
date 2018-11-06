package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	multiLineStr := "for(int i = 0; i < 10; i++){\n  someVar = someVal;\n  moreOperations++;\n  someFunctionCall(2);\n}"
	forLoopRegex := regexp.MustCompile(`for\s*\(\s*int\s+([\w]+)\s*(=)\s*([\w]+)\s*;\s*([\w]+)\s*([<>]|[=]|[<=>=])\s*(\w+);\s*([\w+-=<>]+)\s*\)\s*\{\s*[\r\n]+([^}]+)}`)
	sliceOfMultiLineRegexCaptureGroups := forLoopRegex.FindStringSubmatch(multiLineStr)

	readablePrint := strings.Join(sliceOfMultiLineRegexCaptureGroups, ",\n")
	fmt.Println("[\n", readablePrint, "]\n")
}
