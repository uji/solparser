package lexer

import (
	"fmt"
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

	if token.IsSpace([]rune(lit)[0]) {
		return l.scan()
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

// ScanStringLiteral parse NonEmptyStringLiteral or EmptyStringLiteral then return StringLiteral token.
func (l *Lexer) ScanStringLiteral() (token.Token, error) {
	firstPos, lit, err := l.scanner.Scan()
	if err != nil {
		return token.Token{}, err
	}
	fmt.Println(firstPos, lit)
	if lit != `"` && lit != `\'` {
		return token.Token{}, token.NewPosError(firstPos, `not found " or \'`)
	}

	quote, txt := lit, lit
	for {
		_, lit, err := l.scanner.Scan()
		if err != nil {
			return token.Token{}, err
		}
		txt = txt + lit
		if lit == quote {
			return token.Token{
				TokenType: token.StringLiteral,
				Text:      txt,
				Pos:       firstPos,
			}, nil
		}
	}
}
