package complex_cli_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/dfreilich/gophercon-cli/cmd/complex_cli"
	"github.com/dfreilich/gophercon-cli/pkg/complex"
	"github.com/dfreilich/gophercon-cli/test/mocks"
)

func TestNewCommandWithInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	testInt := mocks.NewMockActor(ctrl)
	cmd := complex_cli.NewCommandWithInterface(testInt)
	testInt.EXPECT().
		DoSomething(
			&complex.DoSomethingArgs{Message: "Hello Gophercon!"})
	require.NoError(t, cmd.Execute())
}
