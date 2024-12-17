package calculator

import (
	parsing "github.com/HayKor/gocalc-api/internal/parsing"
)

// / Evaluate the expression and return the result and the error.
func Calc(expression string) (float64, error) {
	lexer := parsing.NewLexer(expression)
	parser := parsing.NewParser(lexer)

	result, err := parser.ParseExpression()
	if err != nil {
		return 0, err
	}
	return result, nil
}
