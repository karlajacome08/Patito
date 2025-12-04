package semantics

import "fmt"

type Cube map[string]map[TypeTag]map[TypeTag]TypeTag

func NewCube() Cube {
	c := make(Cube)
	add := func(op string, l, r, res TypeTag) {
		if _, ok := c[op]; !ok {
			c[op] = map[TypeTag]map[TypeTag]TypeTag{}
		}
		if _, ok := c[op][l]; !ok {
			c[op][l] = map[TypeTag]TypeTag{}
		}
		c[op][l][r] = res
	}

	// Aritméticos
	for _, op := range []string{"+", "-", "*", "/"} {
		add(op, TInt, TInt, TInt)
		add(op, TInt, TFloat, TFloat)
		add(op, TFloat, TInt, TFloat)
		add(op, TFloat, TFloat, TFloat)
	}
	add("%", TInt, TInt, TInt)

	for _, op := range []string{"==", "!=", "<", ">"} {
		add(op, TInt, TInt, TBool)
		add(op, TInt, TFloat, TBool)
		add(op, TFloat, TInt, TBool)
		add(op, TFloat, TFloat, TBool)
	}
	for _, op := range []string{"==", "!="} {
		add(op, TChar, TChar, TBool)
		add(op, TBool, TBool, TBool)
		add(op, TString, TString, TBool)
	}

	return c
}

func (c Cube) Result(op string, l, r TypeTag) (TypeTag, error) {
	if m1, ok := c[op]; ok {
		if m2, ok := m1[l]; ok {
			if res, ok := m2[r]; ok {
				return res, nil
			}
		}
	}
	return -1, fmt.Errorf("operación inválida: %s (%s, %s)", op, l, r)
}

func AssignCompatible(lhs, rhs TypeTag) bool {
	if lhs == rhs {
		return true
	}
	if lhs == TFloat && rhs == TInt {
		return true
	}
	return false
}

func ResultType(op Op, left, right TypeTag) (TypeTag, bool) {
	switch op {
	case OAdd, OSub, OMul, ODiv:
		if left == TFloat || right == TFloat {
			if (left == TInt || left == TFloat) && (right == TInt || right == TFloat) {
				return TFloat, true
			}
		}
		if left == TInt && right == TInt {
			return TInt, true
		}
		return TInvalid, false
	case OEq, ONEq, OLt, OGt:
		if (left == TInt || left == TFloat) && (right == TInt || right == TFloat) {
			return TBool, true
		}
		return TInvalid, false

	default:
		return TInvalid, false
	}
}

func IsCompatibleAssign(dst, src TypeTag) bool {
	if dst == src {
		return true
	}
	if dst == TFloat && src == TInt {
		return true
	}
	return false
}
