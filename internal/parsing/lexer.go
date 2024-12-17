package parsing

import (
	"unicode"
)

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		if unicode.IsSpace(rune(ch)) {
			l.pos++
			continue
		}

		if unicode.IsDigit(rune(ch)) {
			start := l.pos
			for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
				l.pos++
			}
			return Token{Type: Number, Value: l.input[start:l.pos]}
		}

		switch ch {
		case '+':
			l.pos++
			return Token{Type: Plus, Value: string(ch)}
		case '-':
			l.pos++
			return Token{Type: Minus, Value: string(ch)}
		case '/':
			l.pos++
			return Token{Type: Divide, Value: string(ch)}
		case '*':
			l.pos++
			return Token{Type: Multiply, Value: string(ch)}
		case '(':
			l.pos++
			return Token{Type: LeftParen, Value: string(ch)}
		case ')':
			l.pos++
			return Token{Type: RightParen, Value: string(ch)}
		}

		l.pos++
	}

	return Token{Type: EOF}
}
