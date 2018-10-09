package lisper

// Env contains the Operators defined by default
var Env = map[Value]Op{
	MINUS:  Sub,
	PLUS:   Add,
	SLASH:  Div,
	STAR:   Mul,
	GT:     Gt,
	GE:     Ge,
	LT:     Lt,
	LE:     Le,
	EQ:     Eq,
	NE:     Ne,
	IF:     If,
	DEFINE: Define,
}

// Sym contains the Symbols defined by default
var Sym = map[Value]Value{
	PI: V(3.141592654),
}
