package lexer

import (
	"strings"
	"text/scanner"
)

func Parse(input string) ([]Token, error) {
	lex := NewLexer()

	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.GoTokens

	for lex.next(); lex.cur != scanner.EOF; lex.next() {
		err := state(lex)
		if err != nil {
			return nil, err
		}
	}

	lex.emit(EOF)

	return lex.tokens, nil
}

func NewLexer() *Lexer {
	return &Lexer{}
}

type Lexer struct {
	tokens []Token
	cur    rune // Scanner look ahead
	scan   scanner.Scanner
}

func (lex *Lexer) next() { lex.cur = lex.scan.Scan() }

func (lex *Lexer) peek() rune { return lex.scan.Peek() }

func (lex *Lexer) text() string { return lex.scan.TokenText() }

func (lex *Lexer) accept(valid string) bool { return strings.ContainsRune(valid, lex.peek()) }

func (lex *Lexer) emitWithVal(t Type, v string) { lex.tokens = append(lex.tokens, Token{t, v}) }

func (lex *Lexer) emit(t Type) { lex.emitWithVal(t, lex.text()) }
