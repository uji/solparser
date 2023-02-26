package ast

import (
	"strings"

	"github.com/uji/solparser/token"
)

type Node interface {
	Pos() token.Pos
	End() token.Pos
}

type SymbolAlias struct {
	Symbol Identifier
	As     *token.Pos
	Alias  *Identifier
}

type SymbolAliases struct {
	LBrace  token.Pos
	Aliases []*SymbolAlias
	Commas  []*token.Pos
	RBrace  token.Pos
}

type PragmaDirective struct {
	Pragma       token.Pos
	PragmaTokens []*token.Token
	Semicolon    token.Pos
}

type ModifierList struct {
	Visibility      *Visibility
	StateMutability *StateMutability
}

type Visibility = token.Token

// Parameter is type of ParameterList elements
type Parameter struct {
	TypeName TypeName
}

type ParameterList []*Parameter

type StateMutability = token.Token

// identifier | fallback | recevie
type FunctionDescriptor = token.Token

type FunctionDefinitionReturns struct {
	From          token.Pos
	LParen        token.Pos
	ParameterList ParameterList
	RParen        token.Pos
}

type FunctionDefinition struct {
	From               token.Pos
	FunctionDescriptor FunctionDescriptor
	LParen             token.Pos
	RParen             token.Pos
	ModifierList       *ModifierList
	Returns            *FunctionDefinitionReturns
	Block              *Block
}

func (f FunctionDefinition) Pos() token.Pos {
	return f.From
}

func (f FunctionDefinition) End() token.Pos {
	return f.Block.End()
}

type InheritanceSpecifier struct {
	IdentifierPath   string
	CallArgumentList CallArgumentList
}

type ContractBodyElement interface {
	Node
	contractBodyElementNode()
}

func (f *FunctionDefinition) contractBodyElementNode() {}

type CallArgumentList struct{}

type ContractDefinition struct {
	Abstract             *token.Pos
	Contract             token.Pos
	Identifier           Identifier
	LBrace               token.Pos
	ContractBodyElements []ContractBodyElement
	RBrace               token.Pos
}

// A File node represents a Solidity source file.
type SourceUnit struct {
	PragmaDirective    *PragmaDirective
	ContractDefinition *ContractDefinition
	FunctionDefinition *FunctionDefinition
}

// ----------------------------------------------------------------------------
// TypeName Nodes

type TypeName interface {
	Node
	typeNameNode()
}

type ElementaryTypeName []*token.Token

func (e ElementaryTypeName) Pos() token.Pos {
	return e[0].Position
}

func (e ElementaryTypeName) End() token.Pos {
	eTkn := e[len(e)-1]
	return token.Pos{
		Column: eTkn.Position.Column + len(eTkn.Value),
		Line:   eTkn.Position.Line,
	}
}

func (e ElementaryTypeName) typeNameNode() {}

// ----------------------------------------------------------------------------
// Expression Nodes

type Expression interface {
	Node
	expressionNode()
}

// ----------------------------------------------------------------------------

type Identifier = token.Token

// ----------------------------------------------------------------------------
// Literal Nodes

type Literal interface {
	Node
	literalNode()
	expressionNode()
}

type BooleanLiteral struct {
	Token token.Token
}

func (b *BooleanLiteral) Pos() token.Pos {
	return b.Token.Position
}

func (b *BooleanLiteral) End() token.Pos {
	return token.Pos{
		Column: b.Token.Position.Column + len(b.Token.Value),
		Line:   b.Token.Position.Line,
	}
}

// EmptyStringLiteral | NonEmptyStringLiteral
type StringLiteral token.Token

func (s StringLiteral) Pos() token.Pos {
	return s.Position
}

func (s StringLiteral) End() (pos token.Pos) {
	// Calculate Line and Offset by referring to the number of new line codes
	nc := strings.Count(s.Value, "\n")

	pos.Line = s.Position.Line + nc
	if nc == 0 {
		pos.Column = s.Position.Column + len(s.Value) - 1
		return
	}

	pos.Column = len(s.Value) - strings.LastIndexByte(s.Value, '\n') - 1
	return
}

type HexStringLiteral []*HexString

type UnicordStringLiteral []*UnicordStrings

type NumberLiteral struct {
	Number     Node // DecimalNumber | HexNumber
	NumberUnit *NumberUnit
}

func (*BooleanLiteral) literalNode()       {}
func (*StringLiteral) literalNode()        {}
func (*HexStringLiteral) literalNode()     {}
func (*UnicordStringLiteral) literalNode() {}
func (*NumberLiteral) literalNode()        {}

func (*BooleanLiteral) expressionNode()       {}
func (*StringLiteral) expressionNode()        {}
func (*HexStringLiteral) expressionNode()     {}
func (*UnicordStringLiteral) expressionNode() {}
func (*NumberLiteral) expressionNode()        {}

// ----------------------------------------------------------------------------

type NumberUnit struct {
	Pos   token.Pos
	Value string
}

type EmptyStringLiteral struct {
	Pos          token.Pos
	SingleQuoted bool
}

type NonEmptyStringLiteral struct {
	Pos  token.Pos
	List []Printable
}

// ----------------------------------------------------------------------------
// Printable Nodes

type Printable interface {
	Node
	printableNode()
}

type SingleQuotedPrintable struct {
	Begin  token.Pos
	String string
}

func (s *SingleQuotedPrintable) Pos() token.Pos {
	return s.Begin
}

func (s *SingleQuotedPrintable) End() token.Pos {
	return token.Pos{
		Column: s.Begin.Column + len(s.String) + 1,
		Line:   s.Begin.Line,
	}
}

type DoubleQuotedPrintable struct {
	Begin  token.Pos
	String string
}

func (d *DoubleQuotedPrintable) Pos() token.Pos {
	return d.Begin
}

func (d *DoubleQuotedPrintable) End() token.Pos {
	return token.Pos{
		Column: d.Begin.Column + len(d.String) + 1,
		Line:   d.Begin.Line,
	}
}

type EscapeSequence struct {
	Begin  token.Pos
	String string
}

func (e *EscapeSequence) Pos() token.Pos {
	return e.Begin
}

func (e *EscapeSequence) End() token.Pos {
	return token.Pos{
		Column: e.Begin.Column + len(e.String),
		Line:   e.Begin.Line,
	}
}

func (*SingleQuotedPrintable) printableNode() {}
func (*DoubleQuotedPrintable) printableNode() {}
func (*EscapeSequence) printableNode()        {}

// ----------------------------------------------------------------------------

// unicode-string-literal (https://github.com/ethereum/solidity/blob/develop/docs/grammar/SolidityParser.g4#L407)
type UnicordStrings struct {
	Pos  token.Pos
	List []Node // string | EscapeSequence
}

type HexString struct {
	Pos          token.Pos
	SingleQuoted bool
	String       string
}

// ----------------------------------------------------------------------------

type Block struct {
	LBracePos token.Pos
	RBracePos token.Pos
	Nodes     []Node // statement | unblocked-block
}

func (b Block) Pos() token.Pos {
	return b.LBracePos
}

func (b Block) End() token.Pos {
	return b.RBracePos
}

// ----------------------------------------------------------------------------
// Statement Nodes

type Statement interface {
	Node
	statementNode()
}

type ReturnStatement struct {
	From       token.Pos
	SemiPos    token.Pos
	Expression Expression
}

func (r *ReturnStatement) Pos() token.Pos {
	return r.From
}
func (r *ReturnStatement) End() token.Pos {
	return r.SemiPos
}

func (s *ReturnStatement) statementNode() {}
