package expr

// Expr is an arithmetic expression
type Expr interface {
	Eval(env Env) float64
}

// Var identifies a variable, e.g. x.
type Var string

// A literal is a numeric constant, e.g. 3.14
type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
