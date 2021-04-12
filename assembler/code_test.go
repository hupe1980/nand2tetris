package assembler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComp(t *testing.T) {
	s := NewSymbolTable()
	c := NewCode(&s)

	comp, _ := c.Comp("A")
	require.Equal(t, "0110000", comp)
}

func TestDest(t *testing.T) {
	s := NewSymbolTable()
	c := NewCode(&s)

	dest, _ := c.Dest("M")
	require.Equal(t, "001", dest)
}

func TestJump(t *testing.T) {
	s := NewSymbolTable()
	c := NewCode(&s)

	jump, _ := c.Jump("JEQ")
	require.Equal(t, "010", jump)
}
