package lisper

var Env = map[Value]Op{
	MINUS: Sub,
	PLUS:  Add,
	SLASH: Div,
	STAR:  Mul,
	GT:    Gt,
	GE:    Ge,
	LT:    Lt,
	LE:    Le,
	EQ:    Eq,
	NE:    Ne,
	IF:    If,
}
