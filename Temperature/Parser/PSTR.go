package Parser

import (
	"Polyfy/Temperature/Ast"
	"Polyfy/Temperature/Lexer"
	"Polyfy/Temperature/Token"
	"go/token"
)

const (
	_ int = iota
	LOWEST
	ANDOR
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	ARRAYDEFINE
	OBJECTDEFINE
	CALL
)

var Precedences = map[Token.TokenType]int{
	Token.EQ:       EQUALS,
	Token.NOT_EQ:   EQUALS,
	Token.LT:       LESSGREATER,
	Token.GT:       LESSGREATER,
	Token.PLUS:     SUM,
	Token.MINUS:    SUM,
	Token.SLASH:    PRODUCT,
	Token.ASTERISK: PRODUCT,
	Token.LBRACKET: ARRAYDEFINE,
	Token.LBRACE:   OBJECTDEFINE,
	Token.LPAREN:   CALL,
	Token.AND:      ANDOR,
	Token.OR:       ANDOR,
}

type (
	PrefixParseFn func() Ast.IExpression
	InfixParseFn  func(Ast.IExpression) Ast.IExpression
)

type Parser struct {
	L              *Lexer.Lexer
	Errors         []string
	CurToken       token.Token
	PeekToken      token.Token
	PrefixParseFns map[Token.TokenType]PrefixParseFn
	InfixParseFns  map[Token.TokenType]InfixParseFn
}
