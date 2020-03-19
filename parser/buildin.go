package parser

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func (n BinaryNode) notInArray(env Env) interface{} {
	want := n.x.Eval(env)
	if list, ok := n.y.Eval(env).([]interface{}); ok {
		for _, v := range list {
			if reflect.DeepEqual(want, v) {
				return false
			}
		}
	}
	return true
}

func (n BinaryNode) inArray(env Env) interface{} {
	want := n.x.Eval(env)
	if list, ok := n.y.Eval(env).([]interface{}); ok {
		for _, v := range list {
			if reflect.DeepEqual(want, v) {
				return true
			}
		}
	}
	return false
}

func (n FuncNode) pow(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := num2float64(a)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := num2float64(b)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}

	return math.Pow(x, y)
}

func (n FuncNode) sin(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := num2float64(a)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T)", n.fn, a))
	}
	return math.Sin(x)
}

func (n FuncNode) sqrt(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := num2float64(a)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T)", n.fn, a))
	}
	return math.Sqrt(x)
}

func (n FuncNode) len(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T)", n.fn, a))
	}
	return int64(len(x))
}

func (n FuncNode) lower(env Env) interface{} {
	n.argsCheck(1)
	a := n.args[0].Eval(env)
	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T)", n.fn, a))
	}
	return strings.ToLower(x)
}

func (n FuncNode) index(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := b.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}

	return int64(strings.Index(x, y))
}

func (n FuncNode) contains(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := b.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}

	return strings.Contains(x, y)
}

func (n FuncNode) hasPrefix(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := b.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}

	return strings.HasPrefix(x, y)
}

func (n FuncNode) hasSuffix(env Env) interface{} {
	n.argsCheck(2)

	a := n.args[0].Eval(env)
	b := n.args[1].Eval(env)

	x, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}
	y, ok := b.(string)
	if !ok {
		panic(fmt.Sprintf("invalid arguments: %v(%T, %T)", n.fn, a, b))
	}

	return strings.HasSuffix(x, y)
}

func (n FuncNode) argsCheck(num int) {
	if len(n.args) < num {
		panic(fmt.Sprintf("not enough arguments in call to: %s : ", n.fn))
	}

	if len(n.args) > num {
		panic(fmt.Sprintf("too many arguments in call to: %s : ", n.fn))
	}
}
