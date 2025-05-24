package main

type Expression interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	VisitBinaryExpression(binaryExpression *BinaryExpression)
}

type BinaryExpression struct {
	left     Expression
	operator Token
	right    Expression
}

func (expr *BinaryExpression) Accept(visitor Visitor) { visitor.VisitBinaryExpression(expr) }
