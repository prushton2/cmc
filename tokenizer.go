package main

import "fmt"
import "regexp"

const (
	keyword
	identifier
	constant
	stringliteral
	punctuator
)


type Token struct {
	value string
	tokenType int
}

func parseNextToken() {
	fmt.Println("Tokenizing")

	exp, err := regexp.Compile("(int)([^a-zA-Z0-9])")
	var s = "int a = 2;"

	fmt.Println(exp)
	fmt.Println(err)

	var find = exp.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		fmt.Println("Matched!!")
		fmt.Println(s[find[0][0]:find[0][1]])
	}

	// fmt.Println([0:2])

}