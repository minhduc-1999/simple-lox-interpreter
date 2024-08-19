package expr

import "fmt"

type Visitor interface {
	visitBinary(*Binary) string
	visitGrouping(*Grouping) string
	visitLiteral(*Literal) string
	visitUnary(*Unary) string
}

type Expr interface {
	accept(Visitor) string
}

type Token struct {
}

type Binary struct {
	Left     Expr
	Operator string
	Right    Expr
}

func (b *Binary) accept(v Visitor) string {
	return v.visitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (b *Grouping) accept(v Visitor) string {
	return v.visitGrouping(b)
}

type Literal struct {
	Value any
}

func (b *Literal) accept(v Visitor) string {
	return v.visitLiteral(b)
}

type Unary struct {
	Operator string
	Right    Expr
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
	return fmt.Sprintf("%v %v %v", expr.Left.accept(a), expr.Operator, expr.Right.accept(a))
}

func (a *AstPrinter) visitGrouping(expr *Grouping) string {
	return fmt.Sprintf("(%v)", expr.Expression.accept(a))
}

func (a *AstPrinter) visitLiteral(expr *Literal) string {
	return fmt.Sprintf("%v", expr.Value)
}

func (a *AstPrinter) visitUnary(expr *Unary) string {
	return fmt.Sprintf("%v%v", expr.Operator, expr.Right.accept(a))
}
