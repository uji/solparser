package lexer

import (
	"io"

	scanner "github.com/uji/solparser/scanner2"
	"github.com/uji/solparser/token"
)

type Lexer struct {
	scanner *scanner.Scanner

	// scan result
	token token.Token
	err   error

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
