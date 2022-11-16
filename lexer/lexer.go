package lexer

import (
	"errors"
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

func (l *Lexer) scan() (tkn token.Token, err error) {
	pos, lit, err := l.scanner.Scan()
	if err != nil {
		return token.Token{}, err
	}

	return token.NewToken(lit, pos), nil
}

func (l *Lexer) Scan() (token.Token, error) {
	if l.peeked {
		l.peeked = false
		return l.peekToken, l.peekErr
	}

	return l.scan()
}

func (l *Lexer) Peek() (token.Token, error) {
	if l.peeked {
		return l.peekToken, l.peekErr
	}

	tkn, err := l.scan()
	l.peeked = true
	l.peekToken = tkn
	l.peekErr = err
	return l.peekToken, l.peekErr
}

func (l *Lexer) ScanStringLiteral() (token.Token, error) {
	firstPos, lit, err := l.scanner.Scan()
	if err != nil {
		return token.Token{}, err
	}
	if lit != `"` {
		return token.Token{}, errors.New("test") // TODO: include Pos in error
	}

	txt := lit
	for {
		_, lit, err := l.scanner.Scan()
		if err != nil {
			return token.Token{}, err
		}
		txt = txt + lit
		if lit == `"` {
			return token.Token{
				TokenType: token.StringLiteral,
				Text:      txt,
				Pos:       firstPos,
			}, nil
		}
	}
}
