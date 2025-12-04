package semantics

import "fmt"

type VarKind string

const (
	KindGlobal VarKind = "global"
	KindLocal  VarKind = "local"
	KindParam  VarKind = "param"
	KindTemp   VarKind = "temp"
)

type VariableInfo struct {
	Name       string
	Type       TypeTag
	Kind       VarKind
	Address    int
	Dimensions []int
}

type VariableTable struct {
	Symbols map[string]VariableInfo
}

func NewVariableTable() VariableTable {
	return VariableTable{Symbols: map[string]VariableInfo{}}
}

func (vt *VariableTable) Insert(v VariableInfo) error {
	if _, dup := vt.Symbols[v.Name]; dup {
		return fmt.Errorf("variable doblemente declarada: %s", v.Name)
	}
	vt.Symbols[v.Name] = v
	return nil
}

func (vt *VariableTable) Lookup(name string) (VariableInfo, bool) {
	v, ok := vt.Symbols[name]
	return v, ok
}

type FunctionInfo struct {
	Name       string
	ReturnType TypeTag
	ReturnAddr int // Dirección de la variable de retorno
	Params     []VariableInfo
	Locals     VariableTable
	StartQuad  int
}

type FunctionDirectory struct {
	Fns map[string]*FunctionInfo
}

func NewFunctionDirectory() *FunctionDirectory {
	return &FunctionDirectory{Fns: map[string]*FunctionInfo{}}
}

func (fd *FunctionDirectory) Declare(name string, ret TypeTag) error {
	if _, dup := fd.Fns[name]; dup {
		return fmt.Errorf("función doblemente declarada: %s", name)
	}
	fd.Fns[name] = &FunctionInfo{
		Name:       name,
		ReturnAddr: -1,
		ReturnType: ret,
		Locals:     NewVariableTable(),
	}
	return nil
}

func (fd *FunctionDirectory) Get(name string) (*FunctionInfo, bool) {
	f, ok := fd.Fns[name]
	return f, ok
}

func (fd *FunctionDirectory) AddParam(fn string, v VariableInfo) error {
	f, ok := fd.Get(fn)
	if !ok {
		return fmt.Errorf("función no declarada: %s", fn)
	}
	f.Params = append(f.Params, v)
	return f.Locals.Insert(v)
}

func (fd *FunctionDirectory) AddLocal(fn string, v VariableInfo) error {
	f, ok := fd.Get(fn)
	if !ok {
		return fmt.Errorf("función no declarada: %s", fn)
	}
	return f.Locals.Insert(v)
}
