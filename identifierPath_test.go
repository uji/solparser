package solparser_test

import (
	"testing"

	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
	"github.com/uji/solparser/token"
)

func TestParser_ParseIdentifierPath(t *testing.T) {
	tests := TestData[ast.IdentifierPath]{
		{
			input: "global . identifier",
			want: ast.IdentifierPath{
				Elements: []*ast.IdentifierPathElement{
					{
						Identifier: ast.Identifier(tkn(token.Global, "global", pos(1, 1))),
						Period:     posPtr(8, 1),
					},
					{
						Identifier: ast.Identifier(tkn(token.Identifier, "identifier", pos(10, 1))),
					},
				},
			},
		},
	}

	tests.Test(t, func(p *solparser.Parser) (ast.IdentifierPath, error) {
		return p.ParseIdentifierPath()
	})
}
