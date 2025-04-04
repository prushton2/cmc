package main

import "fmt"

func main() {
	s := "int main(int i) { int offset = 2; return i + offset; }"

	token, index := parseNextToken(s, 0)

	fmt.Println(token)

	for token.value != "" {
		token, index = parseNextToken(s, index)
		
		fmt.Println(token)
	}

}