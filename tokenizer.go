package main

import "fmt"
import "regexp"

const (
	keyword       = 0
	identifier    = 1
	constant      = 2
	stringliteral = 3
	punctuator    = 4
)


type Token struct {
	value string
	tokenType int
}

func parseNextToken() {
	var s = "int a = 2;"

	fmt.Println("Tokenizing")

	var keyword = regexp.MustCompile("(alignas|alignof|auto|bool|break|case|char|const|constexpr|continue|default|do|double|else|enum|extern|false|float|for|goto|if|inline|int|long|nullptr|register|restrict|return|short|signed|sizeof|static|static_assert|struct|switch|thread_local|true|typedef|typeof|typeof_unqual|union|unsigned|void|volatile|while|_Atomic|_BitInt|_Complex|_Decimal128|_Decimal32|_Decimal64|_Generic|_Imaginary|_Noreturn)([^a-zA-Z0-9])")

	// fmt.Println(exp)

	var find = keyword.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		token := Token{tokenType: 1, value: ""}
	}
}