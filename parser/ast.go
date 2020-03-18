package parser

type Node interface {
	Eval(env Env) interface{}
}

type IdentNode struct {
	val string
}

type IntNode struct {
	val int64
}

type FloatNode struct {
	val float64
}

type BoolNode struct {
	val bool
}

type StringNode struct {
	val string
}

type UnaryNode struct {
	op string
	x  Node
}

type BinaryNode struct {
	op   string
	x, y Node
}

type ArrayNode struct {
	args []Node
}

type FuncNode struct {
	fn   string
	args []Node
}
