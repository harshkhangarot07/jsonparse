package main

import (
	"fmt"
	"os"
)

type Parser struct {
	tokens []Token
	pos    int
}

func (p *Parser) nextToken() Token {
	if p.pos < len(p.tokens) {
		token := p.tokens[p.pos]
		p.pos++
		return token
	}
	return Token{Type: TokenEOF}
}

func (p *Parser) parse() error {
	token := p.nextToken()
	if token.Type != TokenLeftBrace {
		return fmt.Errorf("expected '{' but found '%s'", token.Value)
	}

	for {
		token = p.nextToken()
		if token.Type == TokenRightBrace {
			break
		}

		if token.Type != TokenString {
			return fmt.Errorf("expected string key but found '%s'", token.Value)
		}

		token = p.nextToken()
		if token.Type != TokenColon {
			return fmt.Errorf("expected ':' but found '%s'", token.Value)
		}

		if err := p.parseValue(); err != nil {
			return err
		}

		token = p.nextToken()
		if token.Type == TokenRightBrace {
			break
		}

		if token.Type != TokenComma {
			return fmt.Errorf("expected ',' but found '%s'", token.Value)
		}
		if p.pos < len(p.tokens) && p.tokens[p.pos].Type == TokenRightBrace {
			return fmt.Errorf("trailing comma found before '}'")
		}
	}

	token = p.nextToken()
	if token.Type != TokenEOF {
		return fmt.Errorf("expected EOF but found '%s'", token.Value)
	}
	return nil
}

func (p *Parser) parseValue() error {
	token := p.nextToken()
	switch token.Type {
	case TokenLeftBrace:
		return p.parse() // Parse nested objects
	case TokenLeftBracket:
		return p.parseArray() // Parse arrays
	case TokenString:
		return nil // Valid string
	default:
		return fmt.Errorf("unexpected token '%s'", token.Value)
	}
}

func (p *Parser) parseArray() error {
	token := p.nextToken()
	if token.Type == TokenRightBracket {
		return nil // Empty array
	}

	for {
		p.pos-- // Go back for the current token
		if err := p.parseValue(); err != nil {
			return err
		}

		token = p.nextToken()
		if token.Type == TokenRightBracket {
			break
		}

		if token.Type != TokenComma {
			return fmt.Errorf("expected ',' but found '%s'", token.Value)
		}
	}

	return nil
}

func testJSONFile(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("failed to read file %s: %s\n", filepath, err)
		return
	}
	tokens := lexer(string(content))
	parser := &Parser{tokens: tokens}
	err = parser.parse()
	if err != nil {
		fmt.Printf("invalid json: file %s: %s\n", filepath, err)
	} else {
		fmt.Printf("valid %s json\n", filepath)
	}
}

func main() {
	files := []string{
		"tests/step1/valid.json",
		"tests/step1/invalid.json",
		"tests/step2/invalid.json",
		"tests/step2/invalid2.json",
		"tests/step2/valid.json",
		"tests/step2/valid2.json",
	}

	for _, file := range files {
		testJSONFile(file)
	}
}
