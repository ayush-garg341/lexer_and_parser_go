package scanner

import (
	"bufio"
	"bytes"
	"io"
	tokens "rule_engine_parsers/token"
	"strings"
)

// Rune:- stores codes that represent Unicode characters (collection of all possible characters present in the whole world.).
// In Unicode, each of these characters is assigned a unique number called the Unicode code point (hex). This code point is what we store in a rune data type.
// rune is also known as an alias for int32, as each rune can store an integer value of at most 32-bits.
// Go does not have a char data type, so all variables initialized with a character would automatically be typecasted into int32.

// var str string = "ABCD"
// r_array := []rune(str)
// fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)

// var eof = rune(0)
// fmt.Printf("%c", eof)
// var a = rune(97)
// fmt.Printf("%c\n", a)
// var emoji = rune(128512)
// fmt.Printf("%c\n", emoji)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok tokens.Token, lit string) {

	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}
	if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return tokens.EOF, ""
	case '*':
		return tokens.ASTERISK, string(ch)
	case ',':
		return tokens.COMMA, string(ch)
	}
	return tokens.ILLEGAL, ""
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok tokens.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == rune(tokens.EOF) {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tokens.WS, buf.String()
}

func (s *Scanner) scanIdent() (tok tokens.Token, lit string) {
	// Create a buffer and read the current character into it
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident char into the buffer
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return tokens.SELECT, buf.String()
	case "FROM":
		return tokens.FROM, buf.String()
	}

	// Otherwise return as a regular identifier.
	return tokens.IDENT, buf.String()
}

// read reads the next rune from the bufferred reader.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

var eof = rune(0)
