package semantics

import "strconv"

type MemManager struct {
	nextGlobalInt   int
	nextGlobalFloat int
	nextLocalInt    int
	nextLocalFloat  int
	nextTempInt     int
	nextTempFloat   int
	nextConstInt    int
	nextConstFloat  int

	constInts   map[string]int
	constFloats map[string]int
}

func NewMemManager() *MemManager {
	return &MemManager{
		nextGlobalInt:   1000,
		nextGlobalFloat: 2000,
		nextLocalInt:    4000,
		nextLocalFloat:  5000,
		nextTempInt:     7000,
		nextTempFloat:   8000,
		nextConstInt:    9000,
		nextConstFloat:  9500,
		constInts:       map[string]int{},
		constFloats:     map[string]int{},
	}
}

func (m *MemManager) AllocVar(kind VarKind, t TypeTag) int {
	switch kind {
	case KindGlobal:
		switch t {
		case TInt:
			addr := m.nextGlobalInt
			m.nextGlobalInt++
			return addr
		case TFloat:
			addr := m.nextGlobalFloat
			m.nextGlobalFloat++
			return addr
		}
	case KindLocal, KindParam:
		switch t {
		case TInt:
			addr := m.nextLocalInt
			m.nextLocalInt++
			return addr
		case TFloat:
			addr := m.nextLocalFloat
			m.nextLocalFloat++
			return addr
		}
	}
	return -1
}

func (m *MemManager) AllocTemp(t TypeTag) int {
	switch t {
	case TInt:
		addr := m.nextTempInt
		m.nextTempInt++
		return addr
	case TFloat, TBool:
		addr := m.nextTempFloat
		m.nextTempFloat++
		return addr
	default:
		return -1
	}
}

func (m *MemManager) AllocConstInt(lit string) int {
	if addr, ok := m.constInts[lit]; ok {
		return addr
	}
	addr := m.nextConstInt
	m.nextConstInt++
	m.constInts[lit] = addr
	return addr
}

func (m *MemManager) AllocConstFloat(lit string) int {
	if addr, ok := m.constFloats[lit]; ok {
		return addr
	}
	addr := m.nextConstFloat
	m.nextConstFloat++
	m.constFloats[lit] = addr
	return addr
}

func ToAddrString(addr int) string { return strconv.Itoa(addr) }
