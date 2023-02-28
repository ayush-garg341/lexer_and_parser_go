package scanner_test

import (
	scanner "rule_engine_parsers/scanner"
	tokens "rule_engine_parsers/token"
	"strings"
	"testing"
)

type teststruct struct {
	s   string
	tok tokens.Token
	lit string
}

func TestScanner_Scan(t *testing.T) {
	var tests = []teststruct{

		// Special tokens (eof, ws, illegal)
		{s: ``, tok: tokens.EOF},
		{s: `#`, tok: tokens.ILLEGAL, lit: `#`},
		{s: ` `, tok: tokens.WS, lit: " "},
		{s: "\t", tok: tokens.WS, lit: "\t"},
		{s: "\n", tok: tokens.WS, lit: "\n"},

		// Misc chars
		{s: "*", tok: tokens.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: tokens.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: tokens.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: tokens.FROM, lit: "FROM"},
		{s: `SELECT`, tok: tokens.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := scanner.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tok != tt.tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
