package scanner

import (
	"io"

	"github.com/SteelSeries/bufrr"
	"github.com/uji/solparser/token"
)

type Scanner struct {
	// source
	r *bufrr.Reader

	// position state
	offset     int
	lineOffset int
}

func New(reader io.Reader) *Scanner {
	r := bufrr.NewReader(reader)
	return &Scanner{
		r: r,
	}
}

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

func isMultiLengthOperatorSymbol(r rune) bool {
	switch r {
	case '=', '|', '^', '&', '+', '-', '*', '/', '%', '<', '>', '!':
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

const invalidRune = -1

func (s *Scanner) readRune() (rune, error) {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return invalidRune, err
	}
	if r == '\n' {
		s.offset = 0
		s.lineOffset++
		return r, nil
	}
	s.offset++
	return r, nil
}

func (s *Scanner) Scan() (token.Pos, string, error) {
	startPos := token.Pos{
		Column: s.offset + 1,
		Line:   s.lineOffset + 1,
	}

	ch, err := s.readRune()
	if err != nil {
		return token.Pos{}, "", err
	}
	if ch == bufrr.EOF {
		return token.Pos{}, "", io.EOF
	}
	if isSplitSymbol(ch) || ch == '\n' {
		return startPos, string(ch), nil
	}

	readingSpace := isSpace(ch)
	runes := []rune{ch}

	for {
		ch, _, err := s.r.PeekRune()
		if err != nil {
			return token.Pos{}, "", err
		}
		if isSpace(ch) != readingSpace || isSplitSymbol(ch) || ch == '\n' || ch == bufrr.EOF {
			return startPos, string(runes), nil
		}
		ch, err = s.readRune()
		if err != nil {
			return token.Pos{}, "", err
		}
		runes = append(runes, ch)
	}
}

func (s *Scanner) Peek() (pos token.Pos, lit string, err error) {
	return token.Pos{}, "", nil
}
