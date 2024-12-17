package parsing

type TokenType int

const (
	Number TokenType = iota
	Plus
	Minus
	Multiply
	Divide
	EOF
	LeftParen
	RightParen
)

type Token struct {
	Type  TokenType
	Value string
}
