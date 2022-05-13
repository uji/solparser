package lexer

import (
	"bufio"
	"io"
)

type Lexer struct {
	scanner *bufio.Scanner
	token   Token
	err     error

	// position state
	offset     int
	lineOffset int
}

func New(input io.Reader) *Lexer {
	s := bufio.NewScanner(input)
	s.Split(ScanTokens)

	return &Lexer{
		scanner: s,
	}
}

func (l *Lexer) Scan() bool {
	if !l.scanner.Scan() {
		return false
	}
	if err := l.scanner.Err(); err != nil {
		l.err = err
	}

	txt := l.scanner.Text()
	size := len([]rune(txt))
	pos := Position{
		Start: l.offset,
		Size:  size,
		Line:  l.lineOffset,
	}
	l.token = NewToken(txt, pos)

	l.offset += size

	return true
}

func (l Lexer) Token() Token {
	return l.token
}

func (l Lexer) Error() error {
	return l.err
}
