package scanner

import (
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
				{
					Column: 1,
					Line:   1,
				},
				{
					Column: 7,
					Line:   1,
				},
				{
					Column: 8,
					Line:   1,
				},
				{
					Column: 16,
					Line:   1,
				},
				{
					Column: 17,
					Line:   1,
				},
				{
					Column: 18,
					Line:   1,
				},
				{
					Column: 24,
					Line:   1,
				},
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
			name:  "normal",
			input: "pragma\n solidity",
			wantPoss: []token.Pos{
				{
					Column: 1,
					Line:   1,
				},
				{
					Column: 7,
					Line:   1,
				},
				{
					Column: 1,
					Line:   2,
				},
				{
					Column: 2,
					Line:   2,
				},
			},
			wantLits: []string{
				"pragma",
				"\n",
				" ",
				"solidity",
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
