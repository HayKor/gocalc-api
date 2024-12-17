package parsing

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input       string
		expectValue float64
		wantErr     bool
	}{
		// Common check
		{"5+5", 10, false},
		{"5+5*2", 15, false},
		{"(5+5)*2", 20, false},
		{"(5+5)*2+2", 22, false},
		{"5-(5-5)", 5, false},

		// Signs before number check
		{"5-(-5)", 10, false},
		{"5--5", 10, false},

		// Parentheses check
		{"(5+5*2+2", 0, true},
		{"5+5*2)", 0, true},
		{"5+5*2))", 0, true},
		{"(5+5*2))", 0, true},
	}

	for _, tc := range tests {
		lexer := NewLexer(tc.input)
		parser := NewParser(lexer)
		gotValue, err := parser.ParseExpression()
		if err != nil && !tc.wantErr {
			t.Errorf("Got error %v", err.Error())
		}
		if err == nil && tc.wantErr {
			t.Errorf("Expected error, got nil in %v", tc.input)
		} else if tc.expectValue != gotValue {
			t.Errorf("Expected %v, got %v", tc.expectValue, gotValue)
		}
	}
}
