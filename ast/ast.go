package ast

type version string
type expression string

type pragmaValue struct {
	version    string
	expression expression
}

type pragmaDirective struct {
	pragmaName  string
	pragmaValue pragmaValue
}

type functionDescriptor struct {
	name string
}

type stateMutability struct {
	pure bool
}

type modifierList struct {
	stateMutability stateMutability
}

type typeName struct {
	elementalyTypeName string
}

type eventParameter struct {
	typeName typeName
}

type parameterList struct {
	eventParameter eventParameter
}

type returnParameters struct {
	parameterList parameterList
}

type functionDefinition struct {
	functionDescriptor functionDescriptor
	modifierList       modifierList
	returnParameters   returnParameters
}

type contractPart struct {
	functionDefinition functionDefinition
}

type contractDefinition struct {
	contractPart contractPart
}

// A File node represents a Solidity source file.
type SourceUnit struct {
	pragmaDirective    pragmaDirective
	contractDefinition contractDefinition
}
