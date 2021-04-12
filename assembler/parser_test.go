package assembler

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	aCommand = "@2"
	cCommand = "A=M;JGT"
)

func TestParserACommand(t *testing.T) {
	r := strings.NewReader(aCommand)
	s := bufio.NewScanner(r)
	p := NewParser(s)

	p.HasNext()
	p.Next()
	require.Equal(t, A_COMMAND, p.CommandType())
	require.Equal(t, "2", p.Symbol())
}

func TestParserCCommand(t *testing.T) {
	r := strings.NewReader(cCommand)
	s := bufio.NewScanner(r)
	p := NewParser(s)

	p.HasNext()
	p.Next()
	require.Equal(t, C_COMMAND, p.CommandType())
	require.Equal(t, "A", p.Dest())
	require.Equal(t, "M", p.Comp())
	require.Equal(t, "JGT", p.Jump())
}
