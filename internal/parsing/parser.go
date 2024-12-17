package parsing

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lexer   *Lexer
	current Token
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.current = parser.lexer.NextToken()
	return parser
}

// Parse minus and plus as they are least prioritizied
func (p *Parser) ParseExpression() (float64, error) {
	result, err := p.ParseTerm()
	if err != nil {
		return 0, err
	}

	for p.current.Type == Plus || p.current.Type == Minus {
		op := p.current
		p.current = p.lexer.NextToken()
		term, err := p.ParseTerm()
		if err != nil {
			return 0, err
		}
		if op.Type == Plus {
			result += term
		} else {
			result -= term
		}
	}

	return result, nil
}

// Parse Multiply and Division as they are second prioritizied
func (p *Parser) ParseTerm() (float64, error) {
	result, err := p.ParseFactor()
	if err != nil {
		return 0, err
	}

	for p.current.Type == Multiply || p.current.Type == Divide {
		op := p.current
		p.current = p.lexer.NextToken()
		factor, err := p.ParseFactor()
		if err != nil {
			return 0, err
		}
		if op.Type == Multiply {
			result *= factor
		} else {
			if factor == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			result /= factor
		}
	}

	return result, nil
}

func (p *Parser) ParseFactor() (float64, error) {
	if p.current.Type == LeftParen {
		p.current = p.lexer.NextToken()    // Consume '('
		result, err := p.ParseExpression() // Parse the expression inside parentheses
		if err != nil {
			return 0, err
		}
		if p.current.Type != RightParen {
			return 0, fmt.Errorf("expected ')', got %s", p.current.Value)
			// Consume ')'
		}
		p.current = p.lexer.NextToken()
		return result, nil
	}

	if p.current.Type == Number {
		value, err := strconv.ParseFloat(p.current.Value, 64)
		if err != nil {
			return 0, err
		}
		p.current = p.lexer.NextToken()
		return value, nil
	}

	return 0, fmt.Errorf("unexpected token %s", p.current.Value)
}
