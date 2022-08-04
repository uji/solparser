package scanner

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
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

func isSplitSymbol(r rune) bool {
	switch r {
	case '(', ')', '[', ']', '{', '}', ':', ';', '?', '=', '|', '^', '&', '<', '>', '+', '-', '*', '/', '%', ',', '!', '~', '"', '\'', '\\':
		return true
	}
	return false
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	// Return newline code or misk token.
	r, width := utf8.DecodeRune(data[start:])
	if r == '\n' || isSplitSymbol(r) {
		return start + width, data[start : start+width], nil
	}

	tokenIsSpace := isSpace(r)
	// Scan until isSpace result changed, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '\n' || isSplitSymbol(r) || isSpace(r) != tokenIsSpace {
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

type Scanner struct {
	scanner *bufio.Scanner

	// scan result
	text string
	err  error

	// peek state
	peeked   bool
	peekText string
	peekErr  error
}

func New(input io.Reader) *Scanner {
	s := bufio.NewScanner(input)
	s.Split(Split)

	return &Scanner{
		scanner: s,
	}
}

func (s *Scanner) Scan() bool {
	if s.peeked {
		s.text = s.peekText
		s.err = s.peekErr
		s.peeked = false
		return true
	}

	if !s.scanner.Scan() {
		return false
	}

	s.text = s.scanner.Text()
	s.err = s.scanner.Err()
	return true
}

func (s Scanner) Text() string {
	return s.text
}

func (s Scanner) Err() error {
	return s.err
}

func (s *Scanner) Peek() bool {
	if s.peeked {
		return true
	}

	if !s.scanner.Scan() {
		return false
	}

	s.peeked = true
	s.peekText = s.scanner.Text()
	s.peekErr = s.scanner.Err()
	return true
}

func (s Scanner) PeekText() string {
	return s.peekText
}

func (s Scanner) PeekErr() error {
	return s.peekErr
}
