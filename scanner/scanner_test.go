package scanner

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/SteelSeries/bufrr"
	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/token"
)

func TestScanner_Scan(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantPoss []token.Pos
		wantStrs []string
	}{
		{
			name:  "normal",
			input: "pragma solidity ^0.8.13;",
			wantPoss: []token.Pos{
				{Column: 1, Line: 1},
				{Column: 7, Line: 1},
				{Column: 8, Line: 1},
				{Column: 16, Line: 1},
				{Column: 17, Line: 1},
				{Column: 18, Line: 1},
				{Column: 24, Line: 1},
			},
			wantStrs: []string{
				"pragma",
				" ",
				"solidity",
				" ",
				"^",
				"0.8.13",
				";",
			},
		},
		{
			name:  "there is newline character",
			input: "pragma\n solidity",
			wantPoss: []token.Pos{
				{Column: 1, Line: 1},
				{Column: 7, Line: 1},
				{Column: 2, Line: 2},
			},
			wantStrs: []string{
				"pragma",
				"\n ",
				"solidity",
			},
		},
		{
			name:  "there are operators",
			input: `a >> \'test\'`,
			wantPoss: []token.Pos{
				{Column: 1, Line: 1},
				{Column: 2, Line: 1},
				{Column: 3, Line: 1},
				{Column: 5, Line: 1},
				{Column: 6, Line: 1},
				{Column: 8, Line: 1},
				{Column: 12, Line: 1},
			},
			wantStrs: []string{
				"a",
				" ",
				">>",
				" ",
				`\'`,
				"test",
				`\'`,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := New(strings.NewReader(tt.input))

			poss := make([]token.Pos, 0, len(tt.wantPoss))
			strs := make([]string, 0, len(tt.wantStrs))

			for {
				pos, str, err := s.Scan()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("got unexpected error: %s", err)
				}
				poss = append(poss, pos)
				strs = append(strs, str)
			}

			if diff := cmp.Diff(tt.wantPoss, poss); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(tt.wantStrs, strs); diff != "" {
				t.Errorf(diff)
			}
		})
	}

	t.Run("when peeked", func(t *testing.T) {
		wantStr := "str"
		wantPos := token.Pos{
			Column: 3,
			Line:   4,
		}
		wantErr := errors.New(t.Name())
		s := &Scanner{
			r:       bufrr.NewReader(strings.NewReader("input will not scand")),
			peeked:  true,
			peekStr: wantStr,
			peekPos: wantPos,
			peekErr: wantErr,
		}

		pos, str, err := s.Scan()
		if !errors.Is(err, wantErr) {
			t.Fatalf("got unexpected error: %s", err)
		}
		if diff := cmp.Diff(wantPos, pos); diff != "" {
			t.Errorf(diff)
		}
		if str != wantStr {
			t.Errorf("got %s, want %s", str, wantStr)
		}
		if s.peeked != false {
			t.Error("peeked status was not cleared")
		}
	})
}

func TestScanner_Peek(t *testing.T) {
	terr := errors.New(t.Name())
	tests := []struct {
		name      string
		input     string
		peeked    bool
		peekedPos token.Pos
		peekedStr string
		peekedErr error
		wantPos   token.Pos
		wantStr   string
		wantErr   error
	}{
		{
			name:   "not peeked",
			input:  "pragma solidity ^0.8.13;",
			peeked: false,
			wantPos: token.Pos{
				Column: 1,
				Line:   1,
			},
			wantStr: "pragma",
		},
		{
			name:      "peeked",
			peeked:    true,
			peekedStr: "peekedStr",
			peekedPos: token.Pos{
				Column: 2,
				Line:   2,
			},
			wantPos: token.Pos{
				Column: 2,
				Line:   2,
			},
			wantErr: terr,
			wantStr: "peekedStr",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				r:       bufrr.NewReader(strings.NewReader(tt.input)),
				peeked:  tt.peeked,
				peekStr: tt.peekedStr,
				peekPos: tt.peekedPos,
				peekErr: tt.peekedErr,
			}

			pos, str, err := s.Peek()
			if !errors.Is(err, tt.peekedErr) {
				t.Fatalf("got unexpected error: %s", err)
			}
			if diff := cmp.Diff(tt.wantPos, pos); diff != "" {
				t.Errorf(diff)
			}
			if str != tt.wantStr {
				t.Errorf("got %s, want %s", str, tt.wantStr)
			}
		})
	}
}

func TestScanner_scanOperator(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr error
	}{
		{input: "^0.8.13", want: "^"},
		{input: "=>>", want: "=>"},
		{input: "<< ", want: "<<"},
		{input: "<<=a", want: "<<="},
		{input: ">>1", want: ">>"},
		{input: ">>=x", want: ">>="},
		{input: ">>>>", want: ">>>"},
		{input: ">>>==", want: ">>>="},
		{input: "pragma", wantErr: errNotOperator},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			s := New(strings.NewReader(tt.input))

			got, err := s.scanOperator()
			if !errors.Is(err, tt.wantErr) {
				t.Error(err)
			}
			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
			expectedOffset := len(tt.want)
			if s.offset != expectedOffset {
				t.Errorf("offset is %d, expected is %d", s.offset, expectedOffset)
			}
		})
	}
}
