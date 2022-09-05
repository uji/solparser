package scanner

import (
	"io"

	"github.com/uji/solparser/token"
)

type Scanner struct {
	// source
	r io.Reader

	// position state
	offset     int
	lineOffset int
}

func New(reader io.Reader) *Scanner {
	return &Scanner{
		r: reader,
	}
}

func (s *Scanner) readRune() (rune, error) {
	char := make([]byte, 2)

	if _, err := s.r.Read(char); err != nil {
		return 0, err
	}
	return ' ', nil
}

func (s *Scanner) Scan() (pos token.Pos, lit string, err error) {
	return token.Pos{}, "", nil
}

func (s *Scanner) Peek() (pos token.Pos, lit string, err error) {
	return token.Pos{}, "", nil
}
