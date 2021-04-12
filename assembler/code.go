package assembler

import (
	"errors"
	"strconv"
)

var destDict = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var compDict = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

var jumpDict = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type Code interface {
	Addr(string) (string, error)
	Comp(string) (string, error)
	Dest(string) (string, error)
	Jump(string) (string, error)
}

type code struct {
	st      *SymbolTable
	ramAddr int
}

func NewCode(st *SymbolTable) Code {
	return &code{st, 0x0010}
}

func (c *code) Addr(symbol string) (string, error) {
	addr, err := strconv.Atoi(symbol)
	if err == nil {
		// addr
		return int2bin(addr), nil
	}
	// symbol
	if c.st.Contains(symbol) {
		// known symbol
		addr = c.st.GetAddress(symbol)
	} else {
		// new symbol
		addr = c.ramAddr
		c.st.AddEntry(symbol, addr)
		c.ramAddr++
	}
	return int2bin(addr), nil
}

func (c *code) Comp(cmp string) (string, error) {
	comp, ok := compDict[cmp]
	if !ok {
		return "", errors.New("undefined comp: " + cmp)
	}
	return comp, nil
}

func (c *code) Dest(cmp string) (string, error) {
	dest, ok := destDict[cmp]
	if !ok {
		return "", errors.New("undefined dest: " + cmp)
	}
	return dest, nil
}

func (c *code) Jump(cmp string) (string, error) {
	jump, ok := jumpDict[cmp]
	if !ok {
		return "", errors.New("undefined jump: " + cmp)
	}
	return jump, nil
}

func int2bin(num int) string {
	var bin string
	for i := 1 << 14; i > 0; i = i >> 1 {
		if i&num != 0 {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return bin
}
