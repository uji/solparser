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

	readingSpace := token.IsSpace(ch)
	runes := []rune{ch}

	for {
		ch, _, err := s.r.PeekRune()
		if err != nil {
			return token.Pos{}, "", err
		}
		if token.IsSpace(ch) != readingSpace || isSplitSymbol(ch) || ch == bufrr.EOF {
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
