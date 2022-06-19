package lexer

import (
	"bufio"
	"io"
	"unicode/utf8"
)

type scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type Lexer struct {
	scanner scanner

	// scan result
	token Token
	err   error

	// position state
	offset     int
	lineOffset int

	// peek state
	peeked    bool
	peekToken Token
	peekErr   error
}

func New(input io.Reader) *Lexer {
	s := bufio.NewScanner(input)
	s.Split(ScanTokens)

	return &Lexer{
		scanner: s,
	}
}

func (l *Lexer) scan() (result bool, token Token, err error) {
	offset := l.offset
	lineOffset := l.lineOffset

	// Scan until next token.
	for {
		if !l.scanner.Scan() {
			return false, Token{}, nil
		}
		if err := l.scanner.Err(); err != nil {
			return false, Token{}, err
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
	pos := Position{
		Column: offset + 1,
		Line:   lineOffset + 1,
	}

	return true, NewToken(txt, pos), nil
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

func (l Lexer) Token() Token {
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

func (l Lexer) PeekToken() Token {
	return l.peekToken
}

func (l Lexer) PeekError() error {
	return l.peekErr
}
