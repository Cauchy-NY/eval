package parser

import (
	"fmt"
)

type Env map[string]interface{}

func (n IdentNode) Eval(env Env) interface{} {
	return env[n.val]
}

func (n IntNode) Eval(env Env) interface{} {
	return n.val
}

func (n FloatNode) Eval(env Env) interface{} {
	return n.val
}

func (n BoolNode) Eval(env Env) interface{} {
	return n.val
}

func (n StringNode) Eval(env Env) interface{} {
	return n.val
}

func (n UnaryNode) Eval(env Env) interface{} {
	switch n.op {
	case "+":
		return add(0, n.x.Eval(env))
	case "-":
		return sub(0, n.x.Eval(env))
	case "!":
		return not(n.x.Eval(env))
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", n.op))
}

func (n BinaryNode) Eval(env Env) interface{} {
	switch n.op {
	case "+":
		return add(n.x.Eval(env), n.y.Eval(env))
	case "-":
		return sub(n.x.Eval(env), n.y.Eval(env))
	case "*":
		return mul(n.x.Eval(env), n.y.Eval(env))
	case "/":
		return div(n.x.Eval(env), n.y.Eval(env))
	case "%":
		return mod(n.x.Eval(env), n.y.Eval(env))
	case ">":
		return gt(n.x.Eval(env), n.y.Eval(env))
	case "<":
		return lt(n.x.Eval(env), n.y.Eval(env))
	case ">=":
		return ge(n.x.Eval(env), n.y.Eval(env))
	case "<=":
		return le(n.x.Eval(env), n.y.Eval(env))
	case "==":
		return eq(n.x.Eval(env), n.y.Eval(env))
	case "!=":
		return ne(n.x.Eval(env), n.y.Eval(env))
	case "&&":
		return and(n.x.Eval(env), n.y.Eval(env))
	case "||":
		return or(n.x.Eval(env), n.y.Eval(env))
	case "in":
		return n.inArray(env)
	case "not_in":
		return n.notInArray(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", n.op))
}

func (n ArrayNode) Eval(env Env) interface{} {
	var res []interface{}
	for _, v := range n.args {
		res = append(res, v.Eval(env))
	}
	return res
}

func (n FuncNode) Eval(env Env) interface{} {
	switch n.fn {
	case "pow":
		return n.pow(env)
	case "sin":
		return n.sin(env)
	case "sqrt":
		return n.sqrt(env)
	case "len":
		return n.len(env)
	case "lower":
		return n.lower(env)
	case "str_index":
		return n.index(env)
	case "contains":
		return n.contains(env)
	case "has_prefix":
		return n.hasPrefix(env)
	case "has_suffix":
		return n.hasSuffix(env)
	}
	panic(fmt.Sprintf("unsupported function call: %s", n.fn))
}
