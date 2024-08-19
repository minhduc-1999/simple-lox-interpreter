package expr

import (
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
					Left: &Unary{
						Operator: "-",
						Right: &Literal{
							Value: 123,
						},
					},
					Operator: "*",
					Right: &Grouping{
						Expression: &Literal{
							Value: 45.67,
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
					Left: &Binary{
						Left: &Grouping{
							Expression: &Binary{
								Left: &Literal{
									Value: 12,
								},
								Operator: "+",
								Right: &Literal{
									Value: "x",
								},
							},
						},
						Operator: "-",
						Right: &Literal{
							Value: 123,
						},
					},
					Operator: "*",
					Right: &Grouping{
						Expression: &Literal{
							Value: 45.67,
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
