package simple

import "github.com/spf13/cobra"

func NewTestCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "test",
		Version: "0.0.1",
		Short:   "Test is a fabulous CLI",
		Long: `A Fast and Flexible Test CLI built with
                love by David Freilich.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("Hello Gophercon!\n")
		},
	}
}
