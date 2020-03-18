package lexer

import "fmt"

type Type string

const (
	Ident    Type = "Ident"
	Int           = "int"
	Float         = "float"
	Char          = "char"
	Bool          = "bool"
	String        = "String"
	Operator      = "Operator"
	Bracket       = "Bracket"
	EOF           = "EOF"
)

type Token struct {
	tp  Type
	val string
}

func (t Token) Value() string {
	return t.val
}

func (t Token) Type() Type {
	return t.tp
}

func (t Token) String() string {
	if t.val == "" {
		return string(t.tp)
	}
	return fmt.Sprintf("%s(%#v)", t.tp, t.val)
}

func (t Token) Is(tp Type, vals ...string) bool {
	if len(vals) == 0 {
		return tp == t.tp
	}

	for _, v := range vals {
		if v == t.val {
			goto found
		}
	}
	return false

found:
	return tp == t.tp
}
