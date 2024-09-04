package parser

import (
	"fmt"
	"lox/lexer"
)

type Visitor interface {
	visitBinary(*Binary) string
	visitGrouping(*Grouping) string
	visitLiteral(*Literal) string
	visitUnary(*Unary) string
}

type Expr interface {
	accept(Visitor) string
}

type Binary struct {
	left     Expr
	operator lexer.Token
	right    Expr
}

func NewBinary(left Expr, operator lexer.Token, right Expr) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (b *Binary) accept(v Visitor) string {
	return v.visitBinary(b)
}

type Grouping struct {
	expression Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		expression: expr,
	}
}

func (b *Grouping) accept(v Visitor) string {
	return v.visitGrouping(b)
}

type Literal struct {
	value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		value: value,
	}
}

func (b *Literal) accept(v Visitor) string {
	return v.visitLiteral(b)
}

type Unary struct {
	operator lexer.Token
	right    Expr
}

func NewUnary(operator lexer.Token, right Expr) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

func (b *Unary) accept(v Visitor) string {
	return v.visitUnary(b)
}

type AstPrinter struct {
}

func (a *AstPrinter) PrintExpr(expr Expr) string {
	return expr.accept(a)
}

func (a *AstPrinter) visitBinary(expr *Binary) string {
	return fmt.Sprintf("%v %v %v", expr.left.accept(a), expr.operator, expr.right.accept(a))
}

func (a *AstPrinter) visitGrouping(expr *Grouping) string {
	return fmt.Sprintf("(%v)", expr.expression.accept(a))
}

func (a *AstPrinter) visitLiteral(expr *Literal) string {
	return fmt.Sprintf("%v", expr.value)
}

func (a *AstPrinter) visitUnary(expr *Unary) string {
	return fmt.Sprintf("%v%v", expr.operator, expr.right.accept(a))
}
