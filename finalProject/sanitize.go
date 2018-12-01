package main

import (
	"regexp"
)

func sanitize(theString *string) {
	r0 := regexp.MustCompile(`[\s]*([\W\S]+)`)
	r1 := regexp.MustCompile(`[\s]?\/\*([^*]*)\*\/|[\s]?//([^\n]*)`)
	r2 := regexp.MustCompile(` ([ ]+)`)
	r3 := regexp.MustCompile(`[,][\s\t]*\n`)
	twoOrMoreNewLines := regexp.MustCompile(`([\s]?\n[\t\n\s]*)`)
	o0 := r0.FindStringSubmatch(*theString)
	o1 := r1.ReplaceAllString((o0[1]), "")
	o2 := r2.ReplaceAllString(o1, " ")
	o3 := r3.ReplaceAllString(o2, ", ")
	o4 := twoOrMoreNewLines.ReplaceAllString(o3, "\n")
	*theString = o4
}
