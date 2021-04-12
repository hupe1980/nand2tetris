package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hupe1980/nand2tetris/assembler"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No asm file")
		os.Exit(1)
	}

	filename := args[0]
	hackfile := strings.TrimSuffix(filename, ".asm") + ".hack"

	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)

	fo, err := os.Create(hackfile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	writer := bufio.NewWriter(fo)

	// 1st pass
	st := firstPass(scanner)
	code := assembler.NewCode(st)

	// 2nd pass
	fi.Seek(0, 0)
	scanner = bufio.NewScanner(fi)
	p := assembler.NewParser(scanner)
	for p.HasNext() {
		p.Next()
		switch p.CommandType() {
		case assembler.A_COMMAND:
			addr, err := code.Addr(p.Symbol())
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(writer, "0"+addr)
		case assembler.C_COMMAND:
			comp, err := code.Comp(p.Comp())
			if err != nil {
				panic(err)
			}
			dest, err := code.Dest(p.Dest())
			if err != nil {
				panic(err)
			}
			jump, err := code.Jump(p.Jump())
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(writer, "111"+comp+dest+jump)

		case assembler.L_COMMAND:
			// nope
		}
		writer.Flush()
	}

}

func firstPass(scanner *bufio.Scanner) *assembler.SymbolTable {
	st := assembler.NewSymbolTable()
	romAddr := 0
	p := assembler.NewParser(scanner)
	for p.HasNext() {
		p.Next()
		switch p.CommandType() {
		case assembler.A_COMMAND, assembler.C_COMMAND:
			romAddr++
		case assembler.L_COMMAND:
			st.AddEntry(p.Symbol(), romAddr)
		}
	}
	return &st
}
