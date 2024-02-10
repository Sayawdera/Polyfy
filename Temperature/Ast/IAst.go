package Ast

type INode interface {
	TokenLiteral() string
	String() string
}

type IStatement interface {
	INode
	StatementNode()
}

type IExpression interface {
	INode
	ExpressionNode()
}
