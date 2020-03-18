package parser

import (
	"fmt"
	"math"
)

func (n FuncNode) pow(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := a.(float64)
	if !ok {
		panic(fmt.Sprintf("invalid input type: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := b.(float64)
	if !ok {
		panic(fmt.Sprintf("invalid input type: %v(%T, %T)", n.fn, a, b))
	}

	return math.Pow(x, y)
}

func (n FuncNode) sin(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := a.(float64)
	if !ok {
		panic(fmt.Sprintf("invalid input type: %v(%T)", n.fn, a))
	}
	return math.Sin(x)
}

func (n FuncNode) sqrt(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := a.(float64)
	if !ok {
		panic(fmt.Sprintf("invalid input type: %v(%T)", n.fn, a))
	}
	return math.Sqrt(x)
}

func (n FuncNode) argsCheck(num int) {
	if len(n.args) < num {
		panic(fmt.Sprintf("not enough arguments in call to: %s : ", n.fn))
	}

	if len(n.args) > num {
		panic(fmt.Sprintf("too many arguments in call to: %s : ", n.fn))
	}
}
