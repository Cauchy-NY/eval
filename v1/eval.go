package v1

import (
	"fmt"
	"math"
)

type Env map[Var]float64

var bool2float64 = map[bool]float64{
	true:  1,
	false: 0,
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case "+":
		return +u.x.Eval(env)
	case "-":
		return -u.x.Eval(env)
	case "!":
		return bool2float64[u.x.Eval(env) == 0]
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case "+":
		return b.x.Eval(env) + b.y.Eval(env)
	case "-":
		return b.x.Eval(env) - b.y.Eval(env)
	case "*":
		return b.x.Eval(env) * b.y.Eval(env)
	case "/":
		return b.x.Eval(env) / b.y.Eval(env)
	case "%":
		return float64(int64(b.x.Eval(env)) % int64(b.y.Eval(env)))
	case ">":
		return bool2float64[b.x.Eval(env) > b.y.Eval(env)]
	case "<":
		return bool2float64[b.x.Eval(env) < b.y.Eval(env)]
	case ">=":
		return bool2float64[b.x.Eval(env) >= b.y.Eval(env)]
	case "<=":
		return bool2float64[b.x.Eval(env) <= b.y.Eval(env)]
	case "==":
		return bool2float64[b.x.Eval(env) == b.y.Eval(env)]
	case "!=":
		return bool2float64[b.x.Eval(env) != b.y.Eval(env)]
	case "&&":
		return bool2float64[b.x.Eval(env) != 0 && b.y.Eval(env) != 0]
	case "||":
		return bool2float64[b.x.Eval(env) != 0 || b.y.Eval(env) != 0]
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
