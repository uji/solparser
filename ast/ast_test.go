package ast_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestNode_End(t *testing.T) {
	tests := []struct {
		name    string
		node    ast.Node
		exptEnd token.Pos
	}{
		{
			name: "SingleQuotedPrintable",
			node: &ast.SingleQuotedPrintable{
				Begin: token.Pos{
					Column: 4,
					Line:   3,
				},
				String: "test",
			},
			exptEnd: token.Pos{
				Column: 9,
				Line:   3,
			},
		},
		{
			name: "DoubleQuotedPrintable",
			node: &ast.DoubleQuotedPrintable{
				Begin: token.Pos{
					Column: 4,
					Line:   3,
				},
				String: "test",
			},
			exptEnd: token.Pos{
				Column: 9,
				Line:   3,
			},
		},
		{
			name: "EscapeSequence",
			node: &ast.EscapeSequence{
				Begin: token.Pos{
					Column: 4,
					Line:   3,
				},
				String: `\\u2000`,
			},
			exptEnd: token.Pos{
				Column: 11,
				Line:   3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.node.End(), tt.exptEnd); diff != "" {
				t.Error(diff)
			}
		})
	}
}
