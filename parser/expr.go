package parser

import (
	"fmt"
	"lox/lexer"
)

type Visitor interface {
	visitBinary(*Binary) (any, error)
	visitGrouping(*Grouping) (any, error)
	visitLiteral(*Literal) (any, error)
	visitUnary(*Unary) (any, error)
}

type Expr interface {
	accept(Visitor) (any, error)
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

func (b *Binary) accept(v Visitor) (any, error) {
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

func (b *Grouping) accept(v Visitor) (any, error) {
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

func (b *Literal) accept(v Visitor) (any, error) {
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

func (b *Unary) accept(v Visitor) (any, error) {
	return v.visitUnary(b)
}

type AstPrinter struct {
}

func (a *AstPrinter) PrintExpr(expr Expr) string {
	val, err := expr.accept(a)
	if err != nil {
		// TODO
		return fmt.Sprintf("error %v", err)
	}
	return fmt.Sprintf("%v", val)
}

func (a *AstPrinter) visitBinary(expr *Binary) (any, error) {
	leftVal, err := expr.left.accept(a)
	if err != nil {
		return nil, err
	}
	rightVal, err := expr.right.accept(a)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%v %v %v", leftVal, expr.operator, rightVal), nil
}

func (a *AstPrinter) visitGrouping(expr *Grouping) (any, error) {
	val, err := expr.expression.accept(a)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("(%v)", val), nil
}

func (a *AstPrinter) visitLiteral(expr *Literal) (any, error) {
	return fmt.Sprintf("%v", expr.value), nil
}

func (a *AstPrinter) visitUnary(expr *Unary) (any, error) {
	val, err := expr.right.accept(a)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%v%v", expr.operator, val), nil
}
