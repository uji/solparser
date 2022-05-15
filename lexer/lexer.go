package lexer

import (
	"bufio"
	"io"
	"unicode/utf8"
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

func (l *Lexer) Scan() (result bool) {
	// Scan until next token.
	for {
		if !l.scanner.Scan() {
			return false
		}
		if err := l.scanner.Err(); err != nil {
			l.err = err
			return false
		}
		result = true

		txt := l.scanner.Text()
		r, _ := utf8.DecodeRune([]byte(txt))
		if r == '\n' {
			l.lineOffset++
			l.offset = 0
			continue
		}
		if isSpace(r) {
			l.offset += len([]rune(txt))
			continue
		}
		break
	}

	txt := l.scanner.Text()
	size := len([]rune(txt))
	pos := Position{
		Column: l.offset + 1,
		Size:   size,
		Line:   l.lineOffset + 1,
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
