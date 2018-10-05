package lisper

// A bunch of lexemes used by the parser
var (
	MINUS  = V("-")
	PLUS   = V("+")
	SLASH  = V("/")
	STAR   = V("*")
	GT     = V("gt?")
	GE     = V("ge?")
	LT     = V("lt?")
	LE     = V("le?")
	EQ     = V("eq?")
	NE     = V("ne?")
	IF     = V("if")
	DEFINE = V("define")
	DOT    = V(".")
	PI     = V("PI")
	TRUE   = V(true)
	FALSE  = V(false)
	NIL    = V(nil)
)
