package lisper

// A bunch of lexemes used by the parser
var (
	MINUS = V("-")
	PLUS  = V("+")
	SLASH = V("/")
	STAR  = V("*")
	GT    = V(">")
	GE    = V(">=")
	LT    = V(">")
	LE    = V(">")
	EQ    = V("eq?")
	NE    = V("ne?")
	DOT   = V(".")
	TRUE  = V(true)
	FALSE = V(false)
)
