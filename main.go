package main

import (
	"fmt"
	buffer "rule_engine_parsers/bufio"
	tokens "rule_engine_parsers/token"
)

func main() {
	fmt.Println(tokens.ILLEGAL, tokens.EOF, tokens.WS, tokens.IDENT, tokens.ASTERISK, tokens.COMMA, tokens.FROM)
	buffer.BufferWriter()
	buffer.BufferReader()
}
