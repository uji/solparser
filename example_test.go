package solparser_test

import (
	"fmt"
	"strings"

	"github.com/uji/solparser"
	"github.com/uji/solparser/ast"
)

func Example() {
	input :=
		`pragma solidity ^0.8.13;

contract HelloWorld {
    function hello() public pure returns (string) {
        return "Hello World!!";
    }
}`

	parser := solparser.New(strings.NewReader(input))

	got, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(got.ContractDefinition.Identifier.Value)
	fmt.Println(got.ContractDefinition.ContractBodyElements[0].(*ast.FunctionDefinition).FunctionDescriptor.Value)

	// Output:
	// HelloWorld
	// hello
}
