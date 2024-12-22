package parser

import (
	"lox/lexer"
	"reflect"
	"testing"
)

func TestInterpreter_evaluate(t *testing.T) {
	type args struct {
		expr Expr
	}
	tests := []struct {
		name   string
		args   args
		want   any
		errMsg string
	}{
		{
			name: "Case 1",
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
			want: -5617.41,
		},
		{
			name: "Addition",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 12.34,
					},
					operator: lexer.NewToken(lexer.PLUS, "", "", 0),
					right: &Literal{
						value: 56.78,
					},
				},
			},
			want: 69.12,
		},

		// Case with subtraction
		{
			name: "Subtraction",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 100.0,
					},
					operator: lexer.NewToken(lexer.MINUS, "", "", 0),
					right: &Literal{
						value: 50.0,
					},
				},
			},
			want: 50.0,
		},

		// Case with division
		{
			name: "Division",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 10.0,
					},
					operator: lexer.NewToken(lexer.SLASH, "", "", 0),
					right: &Literal{
						value: 2.0,
					},
				},
			},
			want: 5.0,
		},

		// Case with negative number in division
		{
			name: "Negative division",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: -10.0,
					},
					operator: lexer.NewToken(lexer.SLASH, "", "", 0),
					right: &Literal{
						value: 2.0,
					},
				},
			},
			want: -5.0,
		},

		// Case with unary minus
		{
			name: "Unary minus",
			args: args{
				expr: &Unary{
					operator: lexer.NewToken(lexer.MINUS, "", "", 0),
					right: &Literal{
						value: 42.0,
					},
				},
			},
			want: -42.0,
		},

		// Case with parentheses (grouping)
		{
			name: "Parentheses",
			args: args{
				expr: &Grouping{
					expression: &Binary{
						left: &Literal{
							value: 2.0,
						},
						operator: lexer.NewToken(lexer.PLUS, "", "", 0),
						right: &Literal{
							value: 3.0,
						},
					},
				},
			},
			want: 5.0,
		},

		// Case with nested parentheses
		{
			name: "Nested parentheses",
			args: args{
				expr: &Grouping{
					expression: &Binary{
						left: &Grouping{
							expression: &Binary{
								left: &Literal{
									value: 5.0,
								},
								operator: lexer.NewToken(lexer.PLUS, "", "", 0),
								right: &Literal{
									value: 3.0,
								},
							},
						},
						operator: lexer.NewToken(lexer.STAR, "", "", 0),
						right: &Literal{
							value: 2.0,
						},
					},
				},
			},
			want: 16.0, // (5 + 3) * 2 = 16
		},

		// Case with large numbers
		{
			name: "Large numbers",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 1000000000.0,
					},
					operator: lexer.NewToken(lexer.PLUS, "", "", 0),
					right: &Literal{
						value: 500000000.0,
					},
				},
			},
			want: 1500000000.0,
		},

		// Case with zero division
		{
			name: "Zero division",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 10.0,
					},
					operator: lexer.NewToken(lexer.SLASH, "", "", 0),
					right: &Literal{
						value: 0.0,
					},
				},
			},
			errMsg: "Runtime error: division by zero error", // Adjust according to error handling in your code
		},

		// Case with invalid operator (assuming code should fail with invalid operators)
		{
			name: "Invalid operator",
			args: args{
				expr: &Binary{
					left: &Literal{
						value: 10.0,
					},
					operator: lexer.NewToken(lexer.EQUAL, "", "", 0), // Assuming EQUAL is not supported in this context
					right: &Literal{
						value: 5.0,
					},
				},
			},
			errMsg: "Runtime error: invalid operator", // Adjust according to your error handling logic
		},

		// Case with logical operation (if you support boolean logic)
		// {
		// 	name: "Logical AND",
		// 	args: args{
		// 		expr: &Binary{
		// 			left: &Literal{
		// 				value: true,
		// 			},
		// 			operator: lexer.NewToken(lexer.AND, "", "", 0), // Assuming AND is supported in your lexer
		// 			right: &Literal{
		// 				value: false,
		// 			},
		// 		},
		// 	},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Interpreter{}
			got, err := a.evaluate(tt.args.expr)
			if err != nil && !reflect.DeepEqual(err.Error(), tt.errMsg) {
				t.Errorf("[%v] Interpreter.evaluate() = %v, want %v", tt.name, err.Error(), tt.errMsg)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[%v] Interpreter.evaluate() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
