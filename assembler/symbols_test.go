package assembler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSymbolTable(t *testing.T) {
	s := NewSymbolTable()

	s.AddEntry("LOOP", 16)
	require.True(t, s.Contains("LOOP"))
	actual := s.GetAddress("LOOP")
	require.Equal(t, 16, actual)
}
