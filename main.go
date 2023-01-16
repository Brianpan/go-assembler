package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Brianpan/assembler-go/parser"
	"github.com/Brianpan/assembler-go/preprocessor"
	"github.com/Brianpan/assembler-go/symbol"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	argFileName := os.Args[1]
	outputFileName := os.Args[2]

	readFile, err := os.Open(argFileName)
	if err != nil {
		return
	}
	defer readFile.Close()
	symbolTable := symbol.NewSymbolTable()

	fileScanner := bufio.NewScanner(readFile)

	preprocessor := preprocessor.NewPreprocessor(fileScanner, symbolTable)

	parser := parser.NewParser(preprocessor)
	parser.FirstScan(symbolTable)

	fmt.Println("symbol table: \n", symbolTable)
	// fmt.Println("preprocesssor: \n", preprocessor)

	parser.Parse(symbolTable)

	// write to file
	f, err := os.Create(outputFileName)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(parser.String())
	w.Flush()
}
