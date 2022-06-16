package scanner

import (
	"bufio"
	"io"
)

type Scanner struct {
	scanner        *bufio.Scanner
	result         bool
	text           string
	err            error
	peekResult     bool
	peekText       string
	peekErr        error
	firstPeekEnded bool
}

func New(r io.Reader) *Scanner {
	return &Scanner{scanner: bufio.NewScanner(r)}
}

func (s *Scanner) Scan() bool {
	if s.firstPeekEnded {
		s.result = s.peekResult
		s.text = s.peekText
		s.err = s.peekErr
	} else {
		s.result = s.scanner.Scan()
		s.text = s.scanner.Text()
		s.err = s.scanner.Err()
		s.firstPeekEnded = true
	}

	if !s.result {
		return false
	}

	if s.err != nil {
		return false
	}

	s.peekResult = s.scanner.Scan()
	s.peekText = s.scanner.Text()
	s.peekErr = s.scanner.Err()
	return s.result
}

func (s *Scanner) Buffer(buf []byte, max int) {
	s.scanner.Buffer(buf, max)
}

func (s *Scanner) Bytes() []byte {
	return s.scanner.Bytes()
}

func (s *Scanner) Split(split bufio.SplitFunc) {
	s.scanner.Split(split)
}

func (s *Scanner) Text() string {
	return s.text
}

func (s *Scanner) Err() error {
	return s.scanner.Err()
}

func (s *Scanner) PeekText() string {
	return s.peekText
}

func (s *Scanner) PeekErr() error {
	return s.peekErr
}
