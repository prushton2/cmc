package main

import "fmt"
import "regexp"

func parseNextToken() {
	fmt.Println("Tokenizing")

	exp, err := regexp.Compile("int(?!([a-zA-Z0-9]))");

	fmt.Println(exp)
	fmt.Println(err)

	fmt.Println(exp.Find([]byte(`int a = 2`)))

}