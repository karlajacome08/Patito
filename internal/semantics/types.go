package semantics

type TypeTag int

const (
	TInvalid TypeTag = iota
	TInt
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
	case TInvalid:
		return "invalid"
	default:
		return "?"
	}
}

type Op int

const (
	OAdd    Op = iota // +
	OSub              // -
	OMul              // *
	ODiv              // /
	OAssign           // =
	OEq               // ==
	ONEq              // !=
	OLt               // <
	OLe               // <=
	OGt               // >
	OGe               // >=
	OAnd              // &&
	OOr               // ||
	OUMinus           // unary -
	ONot              // !
	ORead             // READ id
	OWrite            // WRITE expr/string
	OParen            // '(' marcador interno para la pila
)

func (o Op) String() string {
	switch o {
	case OAdd:
		return "+"
	case OSub:
		return "-"
	case OMul:
		return "*"
	case ODiv:
		return "/"
	case OAssign:
		return "="
	case OEq:
		return "=="
	case ONEq:
		return "!="
	case OLt:
		return "<"
	case OLe:
		return "<="
	case OGt:
		return ">"
	case OGe:
		return ">="
	case OAnd:
		return "&&"
	case OOr:
		return "||"
	case OUMinus:
		return "unary -"
	case ONot:
		return "!"
	case ORead:
		return "READ"
	case OWrite:
		return "WRITE"
	case OParen:
		return "("
	default:
		return "?"
	}
}

type Addr struct {
	Name string
	Type TypeTag
}
