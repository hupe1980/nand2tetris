package assembler

import (
	"bufio"
	"strings"
)

type Command int

const (
	N_COMMAND Command = iota // empty command
	A_COMMAND                // address ex: @123
	C_COMMAND                // compute ex: A=M;JGT
	L_COMMAND                // label   ex: (LOOP)
)

type Parser interface {
	HasNext() bool
	Next()
	CommandType() Command
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

type parser struct {
	scanner *bufio.Scanner
	text    string
	dest    string
	comp    string
	jump    string
}

func NewParser(scanner *bufio.Scanner) Parser {
	return &parser{scanner, "", "", "", ""}
}

func (p *parser) HasNext() bool {
	return p.scanner.Scan()
}

func (p *parser) Next() {
	p.text = p.scanner.Text()
	if len(p.text) <= 0 {
		return
	}
	// normalization: remove comments
	tokens := strings.SplitN(p.text, "//", 2)
	if len(tokens) > 0 {
		p.text = tokens[0]
	}
	// normalization: remove spaces and tabs
	p.text = strings.TrimSpace(p.text)
}

func (p *parser) CommandType() Command {
	p.dest = ""
	p.comp = ""
	p.jump = ""
	if len(p.text) == 0 {
		// empty token
		return N_COMMAND
	}
	if p.text[0] == '@' {
		// @Xxx
		return A_COMMAND
	}
	if p.text[0] == '(' {
		// (LABEL)
		return L_COMMAND
	}
	// dest=comp;jump
	p.analyze()
	return C_COMMAND
}

func (p *parser) Symbol() string {
	if strings.HasPrefix(p.text, "@") {
		return strings.TrimPrefix(p.text, "@")
	} else {
		return strings.Trim(p.text, "()")
	}
}

func (p *parser) Dest() string {
	return p.dest
}

func (p *parser) Comp() string {
	return p.comp
}

func (p *parser) Jump() string {
	return p.jump
}

// parse "dest=comp;jump"
func (p *parser) analyze() {
	// "dest=comp", "jump"
	tokens := strings.SplitN(p.text, ";", 2)
	destcomp := tokens[0]
	if len(tokens) == 2 {
		p.jump = tokens[1]
	} else {
		p.jump = ""
	}

	// "dest", "comp"
	tokens = strings.SplitN(destcomp, "=", 2)
	if len(tokens) == 2 {
		p.dest = tokens[0]
		p.comp = tokens[1]
	} else {
		p.dest = ""
		p.comp = destcomp
	}
}
