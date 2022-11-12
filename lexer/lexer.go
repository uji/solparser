package lexer

import (
	"io"

	"github.com/uji/solparser/scanner"
	"github.com/uji/solparser/token"
)

type Lexer struct {
	scanner *scanner.Scanner

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
	pos, lit, err := l.scanner.Scan()
	if err != nil {
		return false, token.Token{}, err
	}

	return true, token.NewToken(lit, pos), nil
}

func (l *Lexer) Scan() (token.Token, error) {
	if l.peeked {
		l.peeked = false
		return l.peekToken, l.peekErr
	}

	_, tkn, err := l.scan()
	return tkn, err
}

func (l *Lexer) Peek() (token.Token, error) {
	if l.peeked {
		return l.peekToken, l.peekErr
	}

	_, tkn, err := l.scan()
	l.peeked = true
	l.peekToken = tkn
	l.peekErr = err
	return l.peekToken, l.peekErr
}
