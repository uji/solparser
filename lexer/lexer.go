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
	pos, str, err := l.scanner.Peek()
	if err != nil {
		return token.Token{}, err
	}
	if str == "" {
		return token.Token{}, errors.New("Empty character scanned.")
	}

	if str == `"` || str == `\"` {
		return l.ScanStringLiteral()
	}

	l.scanner.Scan()

	// If space, scan for the next token
	if token.IsSpace([]rune(str)[0]) {
		return l.scan()
	}

	return token.NewToken(str, pos), nil
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

// ScanStringLiteral parse NonEmptyStringLiteral or EmptyStringLiteral then return StringLiteral token.
func (l *Lexer) ScanStringLiteral() (token.Token, error) {
	firstPos, str, err := l.scanner.Scan()
	if err != nil {
		return token.Token{}, err
	}
	if str != `"` && str != `\'` {
		return token.Token{}, token.NewPosError(firstPos, `not found " or \'`)
	}

	quote, txt := str, str
	tokenType := token.EmptyStringLiteral
	for {
		_, str, err := l.scanner.Scan()
		if err != nil {
			return token.Token{}, err
		}
		txt = txt + str
		if str == quote {
			return token.Token{
				Type:     tokenType,
				Value:    txt,
				Position: firstPos,
			}, nil
		}
		tokenType = token.NonEmptyStringLiteral
	}
}
