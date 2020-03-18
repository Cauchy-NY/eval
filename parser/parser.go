package parser

import (
	"errors"
	"fmt"
	"github.com/Cauchy-NY/eval/lexer"
	"strconv"
)

func Parse(input string) (Node, error) {
	tokens, err := lexer.Parse(input)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, errors.New(fmt.Sprintf("input [%s] has none valid ast", input))
	}

	p := NewParser(tokens)

	node := p.parseExpr()

	if !p.cur.Is(lexer.EOF) {
		return nil, errors.New(fmt.Sprintf("unexpected ast [%v]", p.cur))
	}

	if p.err != nil {
		return nil, p.err
	}

	return node, nil
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		cur:    tokens[0],
	}
}

type parserPanic string

type Parser struct {
	tokens []lexer.Token
	cur    lexer.Token
	pos    int
	err    error
}

func (p *Parser) describe() string {
	switch p.cur.Type() {
	case lexer.Ident:
		return fmt.Sprintf("identifier %s", p.cur.Value())
	case lexer.Int, lexer.Float:
		return fmt.Sprintf("number %s", p.cur.Value())
	case lexer.Bool:
		return fmt.Sprintf("bool %s", p.cur.Value())
	case lexer.Char, lexer.String:
		return fmt.Sprintf("string %s", p.cur.Value())
	case lexer.Operator:
		return fmt.Sprintf("operator %s", p.cur.Value())
	case lexer.Bracket:
		return fmt.Sprintf("bracket %s", p.cur.Value())
	case lexer.EOF:
		return "end of file"
	}
	return fmt.Sprintf("%v", p.cur.Value()) // any other rune
}

func (p *Parser) error(format string, args ...interface{}) {
	if p.err == nil { // show first error
		p.err = errors.New(fmt.Sprintf(format, args))
	}
}

func (p *Parser) next() {
	p.pos++
	if p.pos >= len(p.tokens) {
		p.error("unexpected end of expression")
		return
	}
	p.cur = p.tokens[p.pos]
}

func (p *Parser) expect(tp lexer.Type, values ...string) {
	if p.cur.Is(tp, values...) {
		p.next()
		return
	}
	p.error("unexpected ast %v", p.cur)
}

func (p *Parser) parseExpr() Node {
	return p.parseBinary(1)
}

func (p *Parser) parseBinary(basePrec int) Node {
	left := p.parseUnary()
	for prec := precedence(p.cur.Value()); prec >= basePrec; prec-- {
		for precedence(p.cur.Value()) == prec {
			op := p.cur.Value()
			p.next() // consume operator
			right := p.parseBinary(prec + 1)
			left = BinaryNode{op, left, right}
		}
	}
	return left
}

func (p *Parser) parseUnary() Node {
	if p.cur.Is(lexer.Operator, "+", "-", "!") {
		op := string(p.cur.Value())
		p.next() // consume "+", "-" or "!"
		return UnaryNode{op, p.parseUnary()}
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Node {
	switch p.cur.Type() {
	case lexer.Ident:
		ident := p.cur.Value()
		p.next()                  // consume Ident
		if p.cur.Value() == "(" { //deal with buildin func
			p.next() // consume '('
			var args []Node
			if p.cur.Value() != ")" {
				for {
					args = append(args, p.parseExpr())
					if p.cur.Value() != "," {
						break
					}
					p.next() // consume ','
				}
				if p.cur.Value() != ")" {
					msg := fmt.Sprintf("got %v, want ')'", p.cur.Value())
					panic(parserPanic(msg))
				}
			}
			p.next() // consume ')'
			return FuncNode{ident, args}
		} else {
			return IdentNode{ident}
		}
	case lexer.Int:
		i, err := strconv.ParseInt(p.cur.Value(), 10, 64)
		if err != nil {
			panic(parserPanic(err.Error()))
		}
		p.next() // consume int
		return IntNode{i}
	case lexer.Float:
		f, err := strconv.ParseFloat(p.cur.Value(), 64)
		if err != nil {
			panic(parserPanic(err.Error()))
		}
		p.next() // consume float
		return FloatNode{f}
	case lexer.Bool:
		b, err := strconv.ParseBool(p.cur.Value())
		if err != nil {
			panic(parserPanic(err.Error()))
		}
		p.next() // consume bool
		return BoolNode{b}
	case lexer.Char, lexer.String:
		str := p.cur.Value()
		p.next() // consume string or char
		return StringNode{str}
	case lexer.Bracket:
		if p.cur.Value() == "(" {
			p.next() // consume '('
			node := p.parseExpr()
			if p.cur.Value() != ")" {
				msg := fmt.Sprintf("got %v, want ')'", p.describe())
				panic(parserPanic(msg))
			}
			p.next() // consume ')'
			return node
		} else if p.cur.Value() == "[" { // deal with array node
			p.next() // consume '['
			var args []Node
			if p.cur.Value() != "]" {
				for {
					args = append(args, p.parseExpr())
					if p.cur.Value() != "," {
						break
					}
					p.next() // consume ','
				}
				if p.cur.Value() != "]" {
					msg := fmt.Sprintf("got %v, want ']'", p.cur.Value())
					panic(parserPanic(msg))
				}
			}
			p.next() // consume ']'
			return ArrayNode{args}
		} else {
			// map is not support for now
		}
	}
	msg := fmt.Sprintf("unexpected %s", p.describe())
	panic(parserPanic(msg))
}
