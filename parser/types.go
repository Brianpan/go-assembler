package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Brianpan/assembler-go/preprocessor"
	"github.com/Brianpan/assembler-go/symbol"
)

const startAddress = 16

type Parser struct {
	Preprocessor *preprocessor.Preprocessor
	Asms         []string
}

func NewParser(preprocessor *preprocessor.Preprocessor) (parser *Parser) {
	parser = new(Parser)
	parser.Preprocessor = preprocessor
	parser.Asms = make([]string, len(preprocessor.Insts))

	return
}

func (parser *Parser) String() (printStr string) {
	for _, asm := range parser.Asms {
		printStr += fmt.Sprintf("%s\n", asm)
	}

	return
}
func (parser *Parser) FirstScan(symbolTable *symbol.SymbolTable) {
	addr := startAddress

	for _, inst := range parser.Preprocessor.Insts {
		txt := inst.Txt
		if ok, sym := parser.ParseAInstruction(txt); ok {
			// it is symbol
			if _, err := strconv.Atoi(sym); err != nil {
				if _, ok := symbolTable.Table[sym]; !ok {
					symbolTable.Table[sym] = addr
					addr += 1
				}
			}
		}
	}
}

func (parser *Parser) ParseAInstruction(txt string) (ok bool, c string) {
	rAcommand := regexp.MustCompile(`@([\w\\._\\$]+)`)
	if rMatches := rAcommand.FindStringSubmatch(txt); len(rMatches) > 1 {
		ok = true
		c = rMatches[1]
		return
	}
	return
}

func (parser *Parser) Parse(symbolTable *symbol.SymbolTable) {
	for idx, inst := range parser.Preprocessor.Insts {
		txt := inst.Txt
		txt = strings.TrimSpace(txt)

		// A instruction
		if ok, sym := parser.ParseAInstruction(txt); ok {
			addr, ok := symbolTable.Table[sym]
			if !ok {
				addr, _ = strconv.Atoi(sym)
			}
			asm := fmt.Sprintf("0%s", addressConv(addr))
			parser.Asms[idx] = asm
		} else {
			// C instruction
			dest := ""
			comp := ""
			jmp := ""
			// has dest and jmp
			if strings.Contains(txt, "=") && strings.Contains(txt, ";") {
				eqIdx := strings.Index(txt, "=")
				colonIdx := strings.Index(txt, ";")

				dest = txt[:eqIdx]
				comp = txt[eqIdx+1 : colonIdx]
				jmp = txt[colonIdx+1:]

				// has dest no jump
			} else if strings.Contains(txt, "=") && !strings.Contains(txt, ";") {
				eqIdx := strings.Index(txt, "=")
				dest = txt[:eqIdx]
				comp = txt[eqIdx+1:]
				jmp = ""
				// has jump no dest
			} else {
				colonIdx := strings.Index(txt, ";")
				dest = ""
				comp = txt[:colonIdx]
				jmp = txt[colonIdx+1:]
			}
			// fmt.Printf("%s, comp: %s, dest: %s, jmp: %s\n", txt, comp, dest, jmp)
			asm := fmt.Sprintf("111%s%s%s", convertComp(comp), convertDest(dest), convertJmp(jmp))
			parser.Asms[idx] = asm
		}
	}
}

func addressConv(addr int) (r string) {
	s := fmt.Sprintf("%b", addr)
	padding := 15 - len(s)
	if padding > 0 {
		for i := 0; i < padding; i++ {
			r += "0"
		}
	}
	r += s
	return
}

func convertDest(dest string) (r string) {
	switch dest {
	case "":
		r = "000"
		return
	case "M":
		r = "001"
		return
	case "D":
		r = "010"
		return
	case "MD":
		r = "011"
		return
	case "A":
		r = "100"
		return
	case "AM":
		r = "101"
		return
	case "AD":
		r = "110"
		return
	case "AMD":
		r = "111"
		return
	default:
	}
	return
}

func convertComp(comp string) (r string) {
	switch comp {
	case "0":
		r = "0101010"
		return
	case "1":
		r = "0111111"
		return
	case "-1":
		r = "0111010"
		return
	case "D":
		r = "0001100"
		return
	case "A":
		r = "0110000"
		return
	case "M":
		r = "1110000"
		return
	case "!D":
		r = "0001101"
		return
	case "!A":
		r = "0110001"
		return
	case "!M":
		r = "1110001"
		return
	case "-D":
		r = "0001111"
		return
	case "-A":
		r = "0110011"
		return
	case "-M":
		r = "1110011"
		return
	case "D+1":
		r = "0011111"
		return
	case "A+1":
		r = "0110111"
		return
	case "M+1":
		r = "1110111"
		return
	case "D-1":
		r = "0001110"
		return
	case "A-1":
		r = "0110010"
		return
	case "M-1":
		r = "1110010"
		return
	case "D+A":
		r = "0000010"
		return
	case "D+M":
		r = "1000010"
		return
	case "D-A":
		r = "0010011"
		return
	case "D-M":
		r = "1010011"
		return
	case "A-D":
		r = "0000111"
		return
	case "M-D":
		r = "1000111"
		return
	case "D&A":
		r = "0000000"
		return
	case "D&M":
		r = "1000000"
		return
	case "D|A":
		r = "0010101"
		return
	case "D|M":
		r = "1010101"
		return
	default:
	}
	return
}

func convertJmp(jmp string) (r string) {
	switch jmp {
	case "":
		r = "000"
		return
	case "JGT":
		r = "001"
		return
	case "JEQ":
		r = "010"
		return
	case "JGE":
		r = "011"
		return
	case "JLT":
		r = "100"
		return
	case "JNE":
		r = "101"
		return
	case "JLE":
		r = "110"
		return
	case "JMP":
		r = "111"
		return
	default:
	}
	return
}
