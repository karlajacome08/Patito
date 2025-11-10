package semantics

import (
	"errors"
	"fmt"
	"strings"

	p "Entrega1_GO/gen/grammar"

	"github.com/antlr4-go/antlr/v4"
)

type SemanticListener struct {
	*p.BasePatitoListener

	FD   *FunctionDirectory
	Cube Cube

	scope   []string
	types   Stack[TypeTag]
	callRet Stack[TypeTag]

	errors []error
}

func NewSemanticListener() *SemanticListener {
	fd := NewFunctionDirectory()
	_ = fd.Declare("global", TVoid)

	return &SemanticListener{
		FD:    fd,
		Cube:  NewCube(),
		scope: []string{"global"},
	}
}

func (s *SemanticListener) pushScope(name string) { s.scope = append(s.scope, name) }
func (s *SemanticListener) popScope()             { s.scope = s.scope[:len(s.scope)-1] }
func (s *SemanticListener) curScope() string      { return s.scope[len(s.scope)-1] }

func (s *SemanticListener) err(tok antlr.Token, msg string) {
	if tok != nil {
		s.errors = append(s.errors, fmt.Errorf("[L%d:C%d] %s", tok.GetLine(), tok.GetColumn(), msg))
	} else {
		s.errors = append(s.errors, fmt.Errorf("%s", msg))
	}
}

func (s *SemanticListener) Error() error {
	if len(s.errors) == 0 {
		return nil
	}
	var b strings.Builder
	for _, e := range s.errors {
		b.WriteString(e.Error() + "\n")
	}
	return errors.New(b.String())
}
