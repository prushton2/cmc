package main

import "fmt"

func main() {
	s := "int main() { 0.5 500 5.21 }"

	token, index := parseNextToken(s, 0)

	fmt.Println(token)
	// fmt.Println(index)

	for token.value != "" {
		token, index = parseNextToken(s, index)
		
		fmt.Println(token)
	}

}