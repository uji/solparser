package ast_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

var (
	_ ast.ImportDirectiveElement = &ast.ImportDirectivePathElement{}
	_ ast.ImportDirectiveElement = &ast.ImportDirectiveSymbolAliasesElement{}
	_ ast.ImportDirectiveElement = &ast.ImportDirectiveMulElement{}
	_ ast.Node                   = &ast.ImportDirective{}
	_ ast.Node                   = &ast.Path{}
	_ ast.Node                   = &ast.SymbolAlias{}
	_ ast.Node                   = &ast.SymbolAliases{}
	_ ast.Node                   = &ast.PragmaDirective{}
	_ ast.ContractBodyElement    = &ast.FunctionDefinition{}
	_ ast.TypeName               = ast.ElementaryTypeName{}
	_ ast.Expression             = &ast.Identifier{}
	_ ast.Literal                = &ast.BooleanLiteral{}
	_ ast.Literal                = &ast.StringLiteral{}
	_ ast.Node                   = &ast.Block{}
	_ ast.Statement              = &ast.ReturnStatement{}
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
		{
			name: "StringLiteral",
			node: &ast.StringLiteral{
				Type:  token.NonEmptyStringLiteral,
				Value: "\"Hello world!!\"",
				Position: token.Pos{
					Column: 4,
					Line:   3,
				},
			},
			exptEnd: token.Pos{
				Column: 18,
				Line:   3,
			},
		},
		{
			name: "StringLiteral with new line codes",
			node: &ast.StringLiteral{
				Type:  token.NonEmptyStringLiteral,
				Value: "\"Hello \n new \nworld!!\"",
				Position: token.Pos{
					Column: 4,
					Line:   3,
				},
			},
			exptEnd: token.Pos{
				Column: 8,
				Line:   5,
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
