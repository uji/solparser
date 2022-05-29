package solparser_test

import (
	"testing"

	"github.com/uji/solparser"
	"github.com/uji/solparser/lexer"
)

func TestErrorError(t *testing.T) {
	tests := []struct {
		name string
		err  solparser.Error
		want string
	}{
		{
			name: "normal case",
			err: solparser.Error{
				Pos: lexer.Position{
					Column: 1,
					Line:   1,
				},
				Msg: "not found pragma",
			},
			want: "1:1: not found pragma",
		},
		{
			name: "invalid position case",
			err: solparser.Error{
				Pos: lexer.Position{
					Column: 0,
					Line:   1,
				},
				Msg: "not found pragma",
			},
			want: "not found pragma",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Fatalf("want %v, but %v:", tt.want, got)
			}
		})
	}
}
