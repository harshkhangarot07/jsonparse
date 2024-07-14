package main

import "strings"

type TokenType int

const (
	TokenLeftBrace TokenType = iota
	TokenRightBrace
	TokenEOF
	TokenInvalid
	TokenString
	TokenComma
	TokenColon
	TokenLeftBracket
	TokenRightBracket
)

type Token struct {
	Type  TokenType
	Value string
}

func lexer(input string) []Token {
	var tokens []Token
	var strBuilder strings.Builder
	inString := false
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{':
			if !inString {
				tokens = append(tokens, Token{Type: TokenLeftBrace, Value: "{"})
			} else {
				strBuilder.WriteByte(input[i])
			}

		case '}':
			if !inString {
				tokens = append(tokens, Token{Type: TokenRightBrace, Value: "}"})
			} else {
				strBuilder.WriteByte(input[i])
			}

		case ':':
			if !inString {
				tokens = append(tokens, Token{Type: TokenColon, Value: ":"})
			} else {
				strBuilder.WriteByte(input[i])
			}

		case ',':
			if !inString {
				tokens = append(tokens, Token{Type: TokenComma, Value: ","})
			} else {
				strBuilder.WriteByte(input[i])
			}

		case '"':
			if inString {
				tokens = append(tokens, Token{Type: TokenString, Value: strBuilder.String()})
				strBuilder.Reset()
			}
			inString = !inString

		case ' ', '\n', '\r', '\t':
			if inString {
				strBuilder.WriteByte(input[i])
			}

		default:
			if inString {
				strBuilder.WriteByte(input[i])
			} else {
				tokens = append(tokens, Token{Type: TokenInvalid, Value: string(input[i])})
			}
		}

	}

	if inString {
		tokens = append(tokens, Token{Type: TokenInvalid, Value: strBuilder.String()})
	}
	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens
}
