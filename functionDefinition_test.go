package solparser_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_ParseVisibility(t *testing.T) {
	tests := []struct {
		input string
		want  ast.Visibility
		err   *token.PosError
	}{
		{
			input: "internal",
			want: ast.Visibility{
				TokenType: token.Internal,
				Text:      "internal",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found visibility keyword.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseVisibility()

			var sErr *token.PosError
			if errors.As(err, &sErr) {
				if diff := cmp.Diff(tt.err, sErr); diff != "" {
					t.Errorf("%s", diff)
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestParser_ParseStateMutability(t *testing.T) {
	tests := []struct {
		input string
		want  ast.StateMutability
		err   *token.PosError
	}{
		{
			input: "pure",
			want: ast.StateMutability{
				TokenType: token.Pure,
				Text:      "pure",
				Pos:       token.Pos{Column: 1, Line: 1},
			},
		},
		{
			input: "pragma",
			err: &token.PosError{
				Pos: token.Pos{Column: 1, Line: 1},
				Msg: "not found state-mutability keyword.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			r := strings.NewReader(tt.input)
			p := solparser.New(r)

			got, err := p.ParseStateMutability()

			var sErr *token.PosError
			if errors.As(err, &sErr) {
				if diff := cmp.Diff(tt.err, sErr); diff != "" {
					t.Errorf("%s", diff)
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
