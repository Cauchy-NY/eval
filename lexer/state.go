package lexer

import (
	"errors"
	"fmt"
	"strings"
	"text/scanner"
)

var str2op = map[string]string{
	"and": "&&",
	"or":  "||",
	"eq":  "==",
	"ne":  "!=",
	"le":  "<=",
	"ge":  ">=",
	"lt":  "<",
	"gt":  ">",
}

func state(lex *Lexer) error {
	switch lex.cur {

	case scanner.Int:
		lex.emit(Int)
	case scanner.Float:
		lex.emit(Float)
	case scanner.Char:
		lex.emitWithVal(Char, lex.text()[1:len(lex.text())-1])
	case scanner.String:
		lex.emitWithVal(String, lex.text()[1:len(lex.text())-1])
	case scanner.RawString:
		// ignore this for now
	case scanner.Comment:
		// ignore this for now

	case scanner.Ident:
		switch lex.text() {
		case "t", "T", "true", "True", "f", "F", "false", "False", "TRUE", "FALSE":
			lex.emitWithVal(Bool, strings.ToLower(lex.text()))
		case "and", "AND", "or", "OR",
			"le", "LE", "ge", "GE", "lt", "LT", "gt", "GT",
			"eq", "EQ", "ne", "NE": // logic operator
			lex.emitWithVal(Operator, str2op[strings.ToLower(lex.text())])
		default:
			lex.emit(Ident)
		}

	default:
		switch {
		case strings.ContainsRune("{[()]}", lex.cur):
			lex.emit(Bracket)
		case strings.ContainsRune("#,?:%+-/", lex.cur): // single rune operator
			lex.emit(Operator)
		case strings.ContainsRune("&|!=*<>", lex.cur): // possible double rune operator
			op := lex.text()
			if lex.accept("&|=*") {
				lex.next()
				op += lex.text()
			}
			lex.emitWithVal(Operator, op)
		default:
			return errors.New(fmt.Sprintf("unrecognized character: %#U", lex.cur))
		}
	}
	return nil
}
