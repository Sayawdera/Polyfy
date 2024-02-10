package Ast

import (
	"Polyfy/Temperature/Token"
)

type ExpressionStatement struct {
	Token      Token.Token
	Expression IExpression
}

type Identifier struct {
	Token Token.Token
	Value string
}

type Boolean struct {
	Token Token.Token
	Value bool
}

type IntegerLiteral struct {
	Token Token.Token
	Value int64
}

type FloatLiteral struct {
	Token Token.Token
	Value float64
}

type NullLiteral struct {
	Token Token.Token
	Value interface{}
}

type StringLiteral struct {
	Token Token.Token
	Value string
}

type ArrayLiteral struct {
	Token Token.Token
	Elems []IExpression
}

type ObjectLiteral struct {
	Token Token.Token
	Elems map[string]IExpression
}

type PrefixExpression struct {
	Token    Token.Token
	Operator string
	Right    IExpression
}

type InfixExpression struct {
	Token    Token.Token
	Left     IExpression
	Operator string
	Right    IExpression
}

type CallExpression struct {
	Token     Token.Token
	Function  IExpression
	Arguments []IExpression
}
