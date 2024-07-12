package main

type TokenType int

const (
	TokenLeftBrace TokenType = iota
	TokenRightBrace
	TokenEOF
	TokenInvalid
)

type Token struct {
	Type  TokenType
	Value string
}

func lexer(input string) []Token {
	var tokens []Token
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{':
			tokens = append(tokens, Token{Type: TokenLeftBrace, Value: "{"})

		case '}':
			tokens = append(tokens, Token{Type: TokenRightBrace, Value: "}"})

		case ' ', '\n', '\r', '\t':

		default:
			tokens = append(tokens, Token{Type: TokenInvalid, Value: string(input[i])})
		}

	}
	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens
}
