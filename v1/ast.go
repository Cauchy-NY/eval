package v1

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

// A Var identifies a variable, e.g. x.
type Var string

// A literal is a numeric constant, e.g. 3.141.
type literal float64

// A unary represents a unary operator expression, e.g. -x.
type unary struct {
	op string
	x  Expr
}

// A binary represents a binary operator expression, e.g. x+y.
type binary struct {
	op   string
	x, y Expr
}

// A call represents a function call expression, e.g. sin(x).
type call struct {
	fn   string
	args []Expr
}
