package main

import (
	"regexp"
)

func sanitize(theString *string) {

	//r0 := regexp.MustCompile(`([^;]+;)([\W\S]+)`)
	r0 := regexp.MustCompile(`[\s]*([\W\S]+)`)
	r1 := regexp.MustCompile(`\/\*([^*]*)\*\/|//([^\n]*)`)
	r2 := regexp.MustCompile(` ([ ]+)`)
	twoOrMoreNewLines := regexp.MustCompile(`(\n[ \t\n\s]*)`)

	o0 := r0.FindStringSubmatch(*theString)
	o1 := r1.ReplaceAllString((o0[1]), "")

	o2 := r2.ReplaceAllString(o1, " ")
	o3 := twoOrMoreNewLines.ReplaceAllString(o2, "\n")

	*theString = o3
	//fmt.Println(o3)
}
