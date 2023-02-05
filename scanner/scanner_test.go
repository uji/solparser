package scanner

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/token"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantPoss []token.Pos
		wantLits []string
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
			wantLits: []string{
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
			wantLits: []string{
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
			wantLits: []string{
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
			lits := make([]string, 0, len(tt.wantLits))

			for {
				pos, lit, err := s.Scan()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("got unexpected error: %s", err)
				}
				poss = append(poss, pos)
				lits = append(lits, lit)
			}

			if diff := cmp.Diff(tt.wantPoss, poss); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(tt.wantLits, lits); diff != "" {
				t.Errorf(diff)
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
