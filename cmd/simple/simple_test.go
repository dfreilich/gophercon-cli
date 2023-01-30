package simple

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {
	cmd := NewTestCommand()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	require.NoError(t, cmd.Execute())
	require.Equal(t, buf.String(), "Hello Gophercon!\n")
}
