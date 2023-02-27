package parser

import (
	"fmt"
	"io"
	scanner "rule_engine_parsers/scanner"
	tokens "rule_engine_parsers/token"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct {
	Fields    []string
	TableName string
}

// Parser represents a parser.
type Parser struct {
	s   *scanner.Scanner
	buf struct {
		tok tokens.Token // last read token
		lit string       // last read literal
		n   int          // buffer size (max=1)
	}
}

// NewParser returns a new instance of parser
func NewParser(r io.Reader) *Parser {
	return &Parser{s: scanner.NewScanner(r)}
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (*SelectStatement, error) {

	stmt := &SelectStatement{}

	// First token should be a "SELECT" keyword.
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tokens.SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	// Next we should loop over all our comma-delimited fields.
	for {
		// Read a field
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tokens.IDENT && tok != tokens.ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)
		tok, _ = p.scanIgnoreWhitespace()
		if tok != tokens.COMMA {
			p.unscan()
			break
		}
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != tokens.FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	// Finally we should read the table name.
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tokens.IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	// Return the successfully parsed statement.
	return stmt, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok tokens.Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok tokens.Token, lit string) {
	tok, lit = p.scan()
	if tok == tokens.WS {
		tok, lit = p.scan()
	}
	return tok, lit
}
