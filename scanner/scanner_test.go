package scanner_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/uji/solparser/scanner"
)

func TestScan(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		splitFunc    bufio.SplitFunc
		exptText     string
		exptErr      error
		exptPeekText string
		exptPeekErr  error
	}{
		{
			name:         "normal",
			input:        "abc def",
			splitFunc:    bufio.ScanWords,
			exptText:     "abc",
			exptErr:      nil,
			exptPeekText: "def",
			exptPeekErr:  nil,
		},
		{
			name:         "scan one token",
			input:        "abc",
			splitFunc:    bufio.ScanLines,
			exptText:     "abc",
			exptErr:      nil,
			exptPeekText: "",
			exptPeekErr:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := scanner.New(strings.NewReader(c.input))
			s.Split(c.splitFunc)
			if s.Scan() != true {
				t.Errorf("scan returns false")
			}

			gotText := s.Text()
			if gotText != c.exptText {
				t.Errorf("Text() wont %s, get %s", c.exptText, gotText)
			}

			gotErr := s.Err()
			if gotErr != c.exptErr {
				t.Errorf("Err() wont %s, get %s", c.exptErr, gotErr)
			}

			gotPeekText := s.PeekText()
			if gotPeekText != c.exptPeekText {
				t.Errorf("PeekText() wont %s, get %s", c.exptPeekText, gotPeekText)
			}

			gotPeekErr := s.PeekErr()
			if gotPeekErr != c.exptPeekErr {
				t.Errorf("PeekErr() wont %s, get %s", c.exptPeekErr, gotPeekErr)
			}
		})
	}
}
