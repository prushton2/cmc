package main

import "fmt"
import "regexp"

const (
	KEYWORD        = 0
	IDENTIFIER     = 1

	FLOATCONSTANT  = 2
	INTCONSTANT    = 3
	ENUMCONSTANT   = 4
	CHARCONSTANT   = 5
	PREDEFCONSTANT = 6

	STRINGLITERAL  = 7
	PUNCTUATOR     = 8
)


type Token struct {
	value string
	tokenType int
}

func parseNextToken(s string, startIndex int) (Token, int) {
	s = s[startIndex:]

	fmt.Println("Parsing", s)

	var keyword = regexp.MustCompile("(alignas|alignof|auto|bool|break|case|char|const|constexpr|continue|default|do|double|else|enum|extern|false|float|for|goto|if|inline|int|long|nullptr|register|restrict|return|short|signed|sizeof|static|static_assert|struct|switch|thread_local|true|typedef|typeof|typeof_unqual|union|unsigned|void|volatile|while|_Atomic|_BitInt|_Complex|_Decimal128|_Decimal32|_Decimal64|_Generic|_Imaginary|_Noreturn)([^a-zA-Z0-9])")
	var identifier = regexp.MustCompile("([a-zA-Z]|_)([a-zA-Z0-9]|_)*")
	var floatconstant = regexp.MustCompile("[0-9]*\\.[0-9]*")
	var intconstant = regexp.MustCompile("[1-9][0-9]*")

	for s[0] == ' ' || s[0] == '\n' { //destroy any whitespace between tokens
		s = s[1:]
		startIndex++;
	}
	
	var token = Token{tokenType: 0, value: ""}
	
	var find = keyword.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		token = Token{tokenType: KEYWORD, value: s[find[0][0]:find[0][1]-1]}
		return token, startIndex + find[0][1] - find[0][0] - 1;
	}

	find = identifier.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		token = Token{tokenType: IDENTIFIER, value: s[find[0][0]:find[0][1]]}
		return token, startIndex + find[0][1] - find[0][0];
	}

	find = floatconstant.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		token = Token{tokenType: FLOATCONSTANT, value: s[find[0][0]:find[0][1]]}
		return token, startIndex + find[0][1] - find[0][0];
	}

	find = intconstant.FindAllIndex([]byte(s), 1)

	if(find != nil && find[0][0] == 0) {
		token = Token{tokenType: INTCONSTANT, value: s[find[0][0]:find[0][1]]}
		return token, startIndex + find[0][1] - find[0][0];
	}

	// fmt.Println(token)

	return token, startIndex;
}