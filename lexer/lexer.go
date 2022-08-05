package lexer

import (
	"io"
	"unicode/utf8"

	"github.com/uji/solparser/scanner"
	"github.com/uji/solparser/token"
)

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
// TODO: remove function because scanner package has same function.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

type Lexer struct {
	scanner *scanner.Scanner

	// scan result
	token token.Token
	err   error

	// position state
	offset     int
	lineOffset int

	// peek state
	peeked    bool
	peekToken token.Token
	peekErr   error
}

func New(input io.Reader) *Lexer {
	s := scanner.New(input)

	return &Lexer{
		scanner: s,
	}
}

func (l *Lexer) scan() (result bool, tkn token.Token, err error) {
	offset := l.offset
	lineOffset := l.lineOffset

	// Scan until next token.
	for {
		if !l.scanner.Scan() {
			return false, token.Token{}, nil
		}
		if err := l.scanner.Err(); err != nil {
			return false, token.Token{}, err
		}
		txt := l.scanner.Text()
		r, _ := utf8.DecodeRune([]byte(txt))
		if r == '\n' {
			lineOffset++
			offset = 0
			continue
		}
		if isSpace(r) {
			offset += len([]rune(txt))
			continue
		}
		break
	}
	txt := l.scanner.Text()
	pos := token.Pos{
		Column: offset + 1,
		Line:   lineOffset + 1,
	}

	return true, token.NewToken(txt, pos), nil
}

func (l *Lexer) Scan() (result bool) {
	if l.peeked {
		l.token = l.peekToken
		l.err = l.peekErr
		l.peeked = false
		return true
	}

	rslt, tkn, err := l.scan()
	if err != nil {
		l.err = err
		return false
	}
	if !rslt {
		return false
	}
	l.token = tkn
	l.offset = tkn.Pos.Column - 1 + len([]rune(tkn.Text))
	l.lineOffset = tkn.Pos.Line - 1

	return true
}

func (l Lexer) Token() token.Token {
	return l.token
}

func (l Lexer) Error() error {
	return l.err
}

func (l *Lexer) Peek() bool {
	if l.peeked {
		return true
	}

	rslt, tkn, err := l.scan()
	if !rslt {
		return false
	}

	l.peeked = true
	l.peekToken = tkn
	l.peekErr = err
	return true
}

func (l Lexer) PeekToken() token.Token {
	return l.peekToken
}

func (l Lexer) PeekError() error {
	return l.peekErr
}
