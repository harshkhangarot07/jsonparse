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
	token = p.nextToken()
	if token.Type != TokenRightBrace {
		return fmt.Errorf("expected '}' but found '%s' ", token.Value)
	}

	token = p.nextToken()
	if token.Type != TokenEOF {
		return fmt.Errorf("expected EOF but found '%s' ", token.Value)
	}
	return nil
}

func testJSONFile(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("failed to read file %s : %s\n", filepath, err)
		return
	}
	tokens := lexer(string(content))
	parser := &Parser{tokens: tokens}
	err = parser.parse()
	if err != nil {
		fmt.Printf("invalid json : file %s : %s\n", filepath, err)
	} else {
		fmt.Printf("valid %s json  \n", filepath)
	}
}

func main() {
	files := []string{
		"step1/valid.json",
		"step1/invalid.json",
	}

	for _, file := range files {
		testJSONFile(file)
	}
}
