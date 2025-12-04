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
	OAdd Op = iota
	OSub
	OMul
	ODiv
	OAssign
	OEq
	ONEq
	OLt
	OGt
	ORead
	OWrite
	OParen
	OGoto
	OGotoF
	OEra
	OParam
	OGoSub
	OReturn
	OEndFunc
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
	case OGt:
		return ">"
	case ORead:
		return "READ"
	case OWrite:
		return "WRITE"
	case OParen:
		return "("
	case OGoto:
		return "GOTO"
	case OGotoF:
		return "GOTOF"
	case OEra:
		return "ERA"
	case OParam:
		return "PARAM"
	case OGoSub:
		return "GOSUB"
	case OReturn:
		return "RETURN"
	case OEndFunc:
		return "ENDFUNC"
	default:
		return "?"
	}
}

type Addr struct {
	Name    string
	Type    TypeTag
	Address int
}
