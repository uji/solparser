package lexer

import (
	"bufio"
	"io"
	"unicode/utf8"
)

type Lexer struct {
	scanner *bufio.Scanner
	token   Token
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
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

func isMiscToken(r rune) bool {
	return asMiscToken(r) != Invalid
}

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	// Return misk token.
	r, width := utf8.DecodeRune(data[start:])
	if isMiscToken(r) {
		return start + width, data[start : start+width], nil
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isMiscToken(r) {
			return i, data[start:i], nil
		}
		if isSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
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

	txt := l.scanner.Text()
	l.token = Token{
		TokenType: asKeyword(txt),
		Text:      txt,
	}

	return true
}

func (l Lexer) Token() Token {
	return l.token
}
