package semantics

type TypeTag int

const (
	TInt TypeTag = iota
	TFloat
	TChar
	TBool
	TString
	TVoid
)

func (t TypeTag) String() string {
	switch t {
	case TInt:
		return "int"
	case TFloat:
		return "float"
	case TChar:
		return "char"
	case TBool:
		return "bool"
	case TString:
		return "string"
	case TVoid:
		return "void"
	default:
		return "?"
	}
}
