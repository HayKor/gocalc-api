package parsing

import (
	"testing"
)

func TestLexerOnSingleStrings(t *testing.T) {
	tests := []struct {
		input       string
		expectToken Token
	}{
		{"+", Token{Type: Plus, Value: "+"}},
		{"-", Token{Type: Minus, Value: "-"}},
		{"/", Token{Type: Divide, Value: "/"}},
		{"*", Token{Type: Multiply, Value: "*"}},

		{"5", Token{Type: Number, Value: "5"}},
		{"555", Token{Type: Number, Value: "555"}},
	}

	for _, tc := range tests {
		lexer := NewLexer(tc.input)
		gotToken := lexer.NextToken()
		if tc.expectToken != gotToken {
			t.Errorf("Expected %v, got %v", tc.expectToken, gotToken)
		}
	}
}
