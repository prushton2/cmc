package main

import "fmt"
import "bufio"
import "os"

func main() {
// s := "int main(int i) { int offset = 2; return i + offset; }"
	s := "";
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		s += scanner.Text()
	}
	// s := "int main(int i) { int offset = 2; return i + offset; }"

	token, index := parseNextToken(s, 0)

	fmt.Println(token)

	for token.value != "" {
		token, index = parseNextToken(s, index)
		
		fmt.Println(token)
	}

}