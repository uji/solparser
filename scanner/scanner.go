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

func isOperatorRune(r rune) bool {
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

var (
	errNotOperator = errors.New("Not operator.")
)

func (s *Scanner) scanOperator() (string, error) {
	ch1, _, err := s.r.PeekRune()
	if err != nil {
		return "", err
	}
	if !isOperatorRune(ch1) {
		return "", errNotOperator
	}
	if _, err := s.readRune(); err != nil {
		return "", err
	}
	ch2, _, err := s.r.PeekRune()
	if err != nil {
		return string(ch1), nil
	}
	oprt := string([]rune{ch1, ch2})
	switch oprt {
	case "=>", "->", "|=", "^=", "&=", "+=", "-=", "*=", "/=", "%=", "==", "||", "&&", "**", "!=", "<-", ">-", "++", "--", `\'`:
		if _, err := s.readRune(); err != nil {
			return "", err
		}
		return oprt, nil
	case "<<":
		if _, err := s.readRune(); err != nil {
			return "", err
		}
		ch3, _, err := s.r.PeekRune()
		if err != nil {
			return oprt, nil
		}
		if ch3 == '=' {
			if _, err := s.readRune(); err != nil {
				return "", err
			}
			return "<<=", nil
		}
		return oprt, nil
	case ">>":
		if _, err := s.readRune(); err != nil {
			return "", err
		}
		ch3, _, err := s.r.PeekRune()
		if err != nil {
			return oprt, nil
		}
		if ch3 == '=' {
			if _, err := s.readRune(); err != nil {
				return "", err
			}
			return ">>=", nil
		}
		if ch3 == '>' {
			if _, err := s.readRune(); err != nil {
				return "", err
			}
			oprt = ">>>"
			ch4, _, err := s.r.PeekRune()
			if err != nil {
				return oprt, nil
			}
			if ch4 == '=' {
				if _, err := s.readRune(); err != nil {
					return "", err
				}
				return ">>>=", nil
			}
			return oprt, nil
		}
		return oprt, nil
	}
	return string(ch1), nil
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
		return token.Pos{}, "", err
	}
	if ch == bufrr.EOF {
		return token.Pos{}, "", io.EOF
	}
	if isOperatorRune(ch) {
		// scan operator.
		oprt, err := s.scanOperator()
		if err != nil {
			return token.Pos{}, "", err
		}
		return startPos, oprt, nil
	}

	if _, err := s.readRune(); err != nil {
		return token.Pos{}, "", err
	}

	readingSpace := token.IsSpace(ch)
	rslt := []rune{ch}

	for {
		ch, _, err := s.r.PeekRune()
		if err != nil {
			return token.Pos{}, "", err
		}
		if token.IsSpace(ch) != readingSpace || isOperatorRune(ch) || ch == bufrr.EOF {
			return startPos, string(rslt), nil
		}
		if _, err = s.readRune(); err != nil {
			return token.Pos{}, "", err
		}
		rslt = append(rslt, ch)
	}
}

func (s *Scanner) Peek() (pos token.Pos, lit string, err error) {
	return token.Pos{}, "", nil
}
