package symbol

import "fmt"

const (
	SymR0     = "R0"
	SymR1     = "R1"
	SymR2     = "R2"
	SymR3     = "R3"
	SymR4     = "R4"
	SymR5     = "R5"
	SymR6     = "R6"
	SymR7     = "R7"
	SymR8     = "R8"
	SymR9     = "R9"
	SymR10    = "R10"
	SymR11    = "R11"
	SymR12    = "R12"
	SymR13    = "R13"
	SymR14    = "R14"
	SymR15    = "R15"
	SymScreen = "SCREEN"
	SymKbd    = "KBD"
	SymSp     = "SP"
	SymLcl    = "LCL"
	SymArg    = "ARG"
	SymThis   = "THIS"
	SymThat   = "THAT"
)

type SymbolTable struct {
	Table map[string]int
}

func NewSymbolTable() (s *SymbolTable) {
	s = new(SymbolTable)
	s.Table = make(map[string]int, 0)
	// prepare default vals
	for idx := 0; idx < 16; idx++ {
		r := fmt.Sprintf("R%d", idx)
		s.Table[r] = idx
	}

	s.Table[SymScreen] = 16384
	s.Table[SymKbd] = 24576
	s.Table[SymSp] = 0
	s.Table[SymLcl] = 1
	s.Table[SymArg] = 2
	s.Table[SymThis] = 3
	s.Table[SymThat] = 4

	return
}

func (s *SymbolTable) String() (printStr string) {
	for k, v := range s.Table {
		printStr += fmt.Sprintf("Symbol: %s, %d\n", k, v)
	}

	return
}
