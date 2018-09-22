/*
	Richard Stanley
	Alex Truong
	Kenny Chao
	Row 5
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const cRed = "\033[31m"
const cGreen = "\033[1;32m"
const cYellow = "\033[33m"
const cDefault = "\033[0m"

// This function takes a pointer to a regular expression that will attempt to match with the language,
// and print the results of a match, as well as a mismatch of the language.
func printMatchesOrMismatches(reg *regexp.Regexp, s string, location int, structureType string, name string) {
	// the regexp variable 'expressionNotInLanguage' is for extracting the input of what was attempted to be matched:
	// as long as the input is wrapped in something like: 'w1= "someExpression"$' than
	// we can at least inform the user if their word does not match the correct language whatsoever.
	expressionNotInLanguage := regexp.MustCompile(`^[wW]\d{1,}=\s?(\w+)\$$`)
	if reg.MatchString(s) {
		m := reg.FindStringSubmatch(s)
		if len(m) > 1 {
			subMatchIndexSlice := reg.FindSubmatchIndex([]byte(s)) // FindASubmatchIndex returns an int array of indexes
			var subStringStart, subStringStop int
			for i, v := range subMatchIndexSlice {
				if i%2 == 0 && v != 0 && v != -1 {
					subStringStart = v
				} else if i%2 == 1 && v != 0 && v != -1 {
					subStringStop = v
				}
			}
			// (FOR FUN) print in RED or GREEN with the ANSI/VT100 terminal color codes (this is linux/mac OS specific)
			fmt.Print(cGreen, "✓   Match   ", cDefault, ":\t", s[subStringStart:subStringStop]) // print the submatch that is NOT empty
		}
	} else {
		m := expressionNotInLanguage.FindStringSubmatch(s)
		if len(m) == 2 {
			// W#= "someWord"$ is in the correct format, but the word itself does NOT match the Language we are looking for.
			matchForTheExpressionNotInLanguage := m[1]
			fmt.Print(cRed, "⌧  No Match ", cDefault, ":\t", matchForTheExpressionNotInLanguage)
		} else {
			// This will gracefully print the Error when an expression is completely malformed.
			lineOrLocation := func() string {
				if structureType == "file" {
					return " line "
				} else {
					return " location "
				}
			}()
			// print a custom error if completely malformed
			fmt.Print(cRed, "ERROR -", " in ",
				structureType, ": \"", name, "\" at", lineOrLocation, location, ", your input: \"", s, "\" is malformed.", cDefault)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("\nProject 3:")
	fname := "./data.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("  Problem 1:\n    Reading from file ", cYellow, fname, cDefault, " and testing the lines in that file:")
	lineScanner := bufio.NewScanner(file)
	// write an equivalent regex for L=aa*b+bb*
	// This will match:
	/*
		STARTING WITH an uppercase or lower case 'w', one or more digit, equal sign, space(zero or one time),
		the first  expression, a literal '$', END, OR
		STARTING WITH an uppercase or lower case 'w', one or more digit, equal sign, space(zero or one time),
		the second expression, a literal '$', END
	*/
	// This regex language will be store in the langRegexForWordFromFile variable
	langRegexForWordFromFile := regexp.MustCompile(`^[wW]\d{1,}=\s?(aa*b|bb*)\$$`)
	lineNumber := 0
	for lineScanner.Scan() {
		lineNumber++
		line := lineScanner.Text()
		// Attempt to match aaab, and bcbbca
		// print results whether or not there is a match, print a custom error if completely malformed
		fmt.Print("      Line ", lineNumber, ": ")
		printMatchesOrMismatches(langRegexForWordFromFile, line, lineNumber, "file", fname)
	}
	fmt.Println()
	file.Close()

	/*
		Problem 2, part c:
		Now that we have converted the Non-Deterministic Finite Automota:
		 	S -> aA | aB | bB | λ
			B -> bB | λ
			A -> aA | aB
		to L=a*b*, write a regex to match this statement:
		STARTING WITH an uppercase or lower case 'w', one or more digit, equal sign, space(zero or one time),
		'a'(zero or more times)'b'(zero or more times) , a literal '$', END
	*/
	aStarBStar := regexp.MustCompile(`^[wW]\d{1,}=\s?(a*b*)\$$`)
	wordsToTest := []string{"W1= aabba$", "W2= bbbb$"}
	fmt.Print("  Problem 2:\n    Testing the words in array \"wordsToTest\": ", cYellow, "[")
	for idx, word := range wordsToTest {
		if idx != 0 {
			fmt.Print(", \"", word, "\"")
		} else {
			fmt.Print("\"", word, "\"")
		}
	}
	fmt.Print("]", cDefault, " against the language L=a*b* :\n")
	for idx, word := range wordsToTest {
		fmt.Print("      Location ", idx, ": ")
		printMatchesOrMismatches(aStarBStar, word, idx, "array", "wordsToTest")
	}
	fmt.Println()
}
