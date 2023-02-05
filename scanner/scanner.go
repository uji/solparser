package scanner

import (
	"errors"
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
	case '=', '|', '^', '&', '+', '-', '*', '/', '\\', '%', '<', '>', '!':
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

func (s *Scanner) scanOperator() (string, error) {
	ch, _, err := s.r.PeekRune()
	if err != nil {
		return "", err
	}
	if !isSplitSymbol(ch) {
		return "", errors.New("Not operator")
	}
	s.readRune()
	if !isMultiLengthOperatorSymbol(ch) {
		return string(ch), nil
	}
	ch2, _, err := s.r.PeekRune()
	if err != nil {
		return string(ch), nil
	}
	oprt := string([]rune{ch, ch2})
	switch oprt {
	case "=>", "->", "|=", "^=", "&=", "+=", "-=", "*=", "/=", "%=", "==", "||", "&&", "**", "!=", "<-", ">-", "++", "--", `\'`:
		s.readRune()
		return oprt, nil
	case "<<":
		s.readRune()
		ch, _, err := s.r.PeekRune()
		if err != nil {
			return oprt, nil
		}
		if ch == '=' {
			return "<<=", nil
		}
		return oprt, nil
	case ">>":
		s.readRune()
		ch3, _, err := s.r.PeekRune()
		if err != nil {
			return oprt, nil
		}
		if ch3 == '=' {
			return ">>=", nil
		}
		if ch3 == '>' {
			s.readRune()
			oprt = ">>>"
			ch4, _, err := s.r.PeekRune()
			if err != nil {
				return oprt, nil
			}
			if ch4 == '=' {
				s.readRune()
				return ">>>=", nil
			}
			return oprt, nil
		}
		return oprt, nil
	}
	return string(ch), nil
}

// Scan divides a string into the smallest units from the bufrr.Reader and returns them one by one.
//
// - If first rune is an error or EOF, return an error and exit.
// - If first rune is operator character, scan operator.
// - If first rune is a space character, scan until the end of the blank character.
// - Else scan to next space or operator string.
func (s *Scanner) Scan() (token.Pos, string, error) {
	startPos := token.Pos{
		Column: s.offset + 1,
		Line:   s.lineOffset + 1,
	}

	ch, _, err := s.r.PeekRune()
	if err != nil {
		s.readRune()
		return token.Pos{}, "", err
	}
	if ch == bufrr.EOF {
		s.readRune()
		return token.Pos{}, "", io.EOF
	}
	if isSplitSymbol(ch) {
		// scan operator.
		oprt, err := s.scanOperator()
		if err != nil {
			return token.Pos{}, "", err
		}
		return startPos, oprt, nil
	}

	s.readRune()

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
