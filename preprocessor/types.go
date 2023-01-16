package preprocessor

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/Brianpan/assembler-go/symbol"
)

type Preprocessor struct {
	Insts []Inst
}

type Inst struct {
	Line int
	Txt  string
}

func NewPreprocessor(scanner *bufio.Scanner, symbolTable *symbol.SymbolTable) (p *Preprocessor) {
	p = new(Preprocessor)
	p.Insts = make([]Inst, 0)
	line := 0

	reAddressSym := regexp.MustCompile(`([\w\\._\\$]+).*`)

	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.TrimSpace(txt)
		if strings.HasPrefix(txt, "//") || txt == "" {
			continue
		}
		// is address annotation
		// (xxxx)
		if strings.HasPrefix(txt, "(") {
			r := reAddressSym.FindStringSubmatch(txt)
			sym := r[1]
			symbolTable.Table[sym] = line
			continue
		}

		// instructions to remove comments
		if strings.Contains(txt, "//") {
			idx := strings.Index(txt, "//")
			txt = txt[:idx]
			inst := Inst{
				Line: line,
				Txt:  txt,
			}
			p.Insts = append(p.Insts, inst)
			line += 1
			continue
		}

		// no prune case
		inst := Inst{
			Line: line,
			Txt:  txt,
		}
		p.Insts = append(p.Insts, inst)
		line += 1
	}

	return
}

func (p *Preprocessor) String() (printString string) {
	for _, inst := range p.Insts {
		printString += fmt.Sprintf("%d: %s\n", inst.Line, inst.Txt)
	}

	return
}
