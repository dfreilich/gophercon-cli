package complex_cli

import (
	"github.com/spf13/cobra"

	"github.com/dfreilich/gophercon-cli/pkg/complex"
)

func NewCommandWithInterface(actor complex.Actor) *cobra.Command {
	return &cobra.Command{
		Use:     "complex",
		Version: "0.0.1",
		Short:   "Complex uses a library to abstract out",
		RunE: func(cmd *cobra.Command, args []string) error {
			return actor.DoSomething(&complex.DoSomethingArgs{
				Message: "Hello Gophercon!",
			})
		},
	}
}
