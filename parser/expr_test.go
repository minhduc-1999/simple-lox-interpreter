package parser

import (
	"lox/lexer"
	"testing"
)

func TestAstPrinter_PrintExpr(t *testing.T) {
	type args struct {
		expr Expr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				expr: &Binary{
					left: &Unary{
						operator: lexer.NewToken(lexer.MINUS, "", "", 0),
						right: &Literal{
							value: 123,
						},
					},
					operator: lexer.NewToken(lexer.STAR, "", "", 0),
					right: &Grouping{
						expression: &Literal{
							value: 45.67,
						},
					},
				},
			},
			want: "-123 * (45.67)",
		},
		{
			name: "case 2",
			args: args{
				expr: &Binary{
					left: &Binary{
						left: &Grouping{
							expression: &Binary{
								left: &Literal{
									value: 12,
								},
								operator: lexer.NewToken(lexer.PLUS, "", "", 0),
								right: &Literal{
									value: "x",
								},
							},
						},
						operator: lexer.NewToken(lexer.MINUS, "", "", 0),
						right: &Literal{
							value: 123,
						},
					},
					operator: lexer.NewToken(lexer.STAR, "", "", 0),
					right: &Grouping{
						expression: &Literal{
							value: 45.67,
						},
					},
				},
			},
			want: "(12 + x) - 123 * (45.67)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AstPrinter{}
			if got := a.PrintExpr(tt.args.expr); got != tt.want {
				t.Errorf("AstPrinter.PrintExpr()\ngot:\t%v\nwant:\t%v", got, tt.want)
			}
		})
	}
}
