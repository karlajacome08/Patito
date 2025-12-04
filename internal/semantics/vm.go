package semantics

import (
	"fmt"
	"math"
	"strconv"
)

// VM ejecuta la lista de cuádruplos usando direcciones virtuales.
type VM struct {
	quads *QuadQueue
	mem   []float64
	fd    *FunctionDirectory

	callStack   []int
	currentCall string
	ip          int
}

func NewVM(quads *QuadQueue, memMgr *MemManager, fd *FunctionDirectory) *VM {
	vm := &VM{
		quads: quads,
		fd:    fd,
	}
	vm.initMemory(memMgr)
	vm.loadConstants(memMgr)
	vm.ip = vm.findMainStart()
	return vm
}

func (vm *VM) initMemory(m *MemManager) {
	max := -1
	candidates := []int{
		m.nextGlobalInt - 1,
		m.nextGlobalFloat - 1,
		m.nextLocalInt - 1,
		m.nextLocalFloat - 1,
		m.nextTempInt - 1,
		m.nextTempFloat - 1,
		m.nextConstInt - 1,
		m.nextConstFloat - 1,
	}
	for _, v := range candidates {
		if v > max {
			max = v
		}
	}
	if max < 0 {
		max = 0
	}
	vm.mem = make([]float64, max+1)
}

func (vm *VM) loadConstants(m *MemManager) {
	for lit, addr := range m.constInts {
		v, err := strconv.Atoi(lit)
		if err == nil && addr >= 0 && addr < len(vm.mem) {
			vm.mem[addr] = float64(v)
		}
	}
	for lit, addr := range m.constFloats {
		v, err := strconv.ParseFloat(lit, 64)
		if err == nil && addr >= 0 && addr < len(vm.mem) {
			vm.mem[addr] = v
		}
	}
}

func (vm *VM) findMainStart() int {
	lastEnd := -1
	for i := 0; i < vm.quads.Len(); i++ {
		if vm.quads.At(i).Op == OEndFunc {
			lastEnd = i
		}
	}
	return lastEnd + 1
}

func parseIntOrZero(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (vm *VM) get(addrStr string) float64 {
	if addrStr == "" {
		return 0
	}
	addr, err := strconv.Atoi(addrStr)
	if err != nil || addr < 0 || addr >= len(vm.mem) {
		return 0
	}
	return vm.mem[addr]
}

func (vm *VM) set(addrStr string, val float64) {
	if addrStr == "" {
		return
	}
	addr, err := strconv.Atoi(addrStr)
	if err != nil || addr < 0 || addr >= len(vm.mem) {
		return
	}
	vm.mem[addr] = val
}

func (vm *VM) boolFrom(addrStr string) bool {
	return vm.get(addrStr) != 0
}

func (vm *VM) execWrite(arg1 string) {
	if arg1 == "" {
		fmt.Println()
		return
	}
	// ¿Es dirección numérica?
	if addr, err := strconv.Atoi(arg1); err == nil {
		if addr >= 0 && addr < len(vm.mem) {
			v := vm.mem[addr]
			if math.Trunc(v) == v {
				fmt.Println(int(v))
			} else {
				fmt.Println(v)
			}
		} else {
			fmt.Println(0)
		}
		return
	}

	s := arg1
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	fmt.Println(s)
}

func (vm *VM) execParam(q Quad) {
	if vm.currentCall == "" {
		return
	}
	fn, ok := vm.fd.Get(vm.currentCall)
	if !ok {
		return
	}
	idx := parseIntOrZero(q.Result)
	if idx < 0 || idx >= len(fn.Params) {
		return
	}
	srcAddr, err := strconv.Atoi(q.Arg1)
	if err != nil || srcAddr < 0 || srcAddr >= len(vm.mem) {
		return
	}
	dstAddr := fn.Params[idx].Address
	if dstAddr < 0 || dstAddr >= len(vm.mem) {
		return
	}
	vm.mem[dstAddr] = vm.mem[srcAddr]
}

func (vm *VM) Run() error {
	for vm.ip >= 0 && vm.ip < vm.quads.Len() {
		q := vm.quads.At(vm.ip)

		switch q.Op {
		// Aritméticos
		case OAdd:
			vm.set(q.Result, vm.get(q.Arg1)+vm.get(q.Arg2))
		case OSub:
			vm.set(q.Result, vm.get(q.Arg1)-vm.get(q.Arg2))
		case OMul:
			vm.set(q.Result, vm.get(q.Arg1)*vm.get(q.Arg2))
		case ODiv:
			den := vm.get(q.Arg2)
			if den == 0 {
				return fmt.Errorf("división entre cero en quad %d", vm.ip)
			}
			vm.set(q.Result, vm.get(q.Arg1)/den)

		// Asignación
		case OAssign:
			vm.set(q.Result, vm.get(q.Arg1))

		case OEq:
			if vm.get(q.Arg1) == vm.get(q.Arg2) {
				vm.set(q.Result, 1)
			} else {
				vm.set(q.Result, 0)
			}
		case ONEq:
			if vm.get(q.Arg1) != vm.get(q.Arg2) {
				vm.set(q.Result, 1)
			} else {
				vm.set(q.Result, 0)
			}
		case OLt:
			if vm.get(q.Arg1) < vm.get(q.Arg2) {
				vm.set(q.Result, 1)
			} else {
				vm.set(q.Result, 0)
			}
		case OGt:
			if vm.get(q.Arg1) > vm.get(q.Arg2) {
				vm.set(q.Result, 1)
			} else {
				vm.set(q.Result, 0)
			}

		// I/O
		case OWrite:
			vm.execWrite(q.Arg1)

		// Saltos
		case OGoto:
			vm.ip = parseIntOrZero(q.Result)
			continue
		case OGotoF:
			if !vm.boolFrom(q.Arg1) {
				vm.ip = parseIntOrZero(q.Result)
				continue
			}

		// Funciones
		case OEra:
			vm.currentCall = q.Arg1

		case OParam:
			vm.execParam(q)

		case OGoSub:
			ret := vm.ip + 1
			vm.callStack = append(vm.callStack, ret)
			vm.ip = parseIntOrZero(q.Result)
			continue

		case OEndFunc:
			n := len(vm.callStack)
			if n == 0 {
				return nil
			}
			ret := vm.callStack[n-1]
			vm.callStack = vm.callStack[:n-1]
			vm.ip = ret
			continue

		default:
		}

		vm.ip++
	}
	return nil
}
