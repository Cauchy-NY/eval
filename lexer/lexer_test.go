package lexer

import (
	"testing"
)

type lexTest struct {
	input  string
	tokens []Token
}

var lexTests = []lexTest{
	{
		`2 < a || b < 9`,
		[]Token{
			{Int, "2"},
			{Operator, "<"},
			{Ident, "a"},
			{Operator, "||"},
			{Ident, "b"},
			{Operator, "<"},
			{Int, "9"},
			{EOF, ""},
		},
	},

	{
		`((a + 2) == 3) > 0`,
		[]Token{
			{Bracket, "("},
			{Bracket, "("},
			{Ident, "a"},
			{Operator, "+"},
			{Int, "2"},
			{Bracket, ")"},
			{Operator, "=="},
			{Int, "3"},
			{Bracket, ")"},
			{Operator, ">"},
			{Int, "0"},
			{EOF, ""},
		},
	},

	{
		`pow(x, 3) + pow(y, 3)`,
		[]Token{
			{Ident, "pow"},
			{Bracket, "("},
			{Ident, "x"},
			{Operator, ","},
			{Int, "3"},
			{Bracket, ")"},
			{Operator, "+"},
			{Ident, "pow"},
			{Bracket, "("},
			{Ident, "y"},
			{Operator, ","},
			{Int, "3"},
			{Bracket, ")"},
			{EOF, ""},
		},
	},

	{
		`!(a > 0) || False`,
		[]Token{
			{Operator, "!"},
			{Bracket, "("},
			{Ident, "a"},
			{Operator, ">"},
			{Int, "0"},
			{Bracket, ")"},
			{Operator, "||"},
			{Bool, "false"},
			{EOF, ""},
		},
	},

	{
		`note == "hello, world"`,
		[]Token{
			{Ident, "note"},
			{Operator, "=="},
			{String, "hello, world"},
			{EOF, ""},
		},
	},

	{
		`grade >= 'a'`,
		[]Token{
			{Ident, "grade"},
			{Operator, ">="},
			{Char, "a"},
			{EOF, ""},
		},
	},

	{
		`a and b`,
		[]Token{
			{Ident, "a"},
			{Operator, "&&"},
			{Ident, "b"},
			{EOF, ""},
		},
	},
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		tokens, err := Parse(test.input)
		if err != nil {
			t.Errorf("%s:\n%v", test.input, err)
			return
		}
		if !compareTokens(tokens, test.tokens) {
			t.Errorf("%s:\ngot\n\t%+v\nexpected\n\t%v", test.input, tokens, test.tokens)
		}
	}
}

func compareTokens(i1, i2 []Token) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].tp != i2[k].tp {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
	}
	return true
}
