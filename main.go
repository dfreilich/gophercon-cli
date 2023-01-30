/*
Copyright © 2023 David Freilich freilich.david@gmail.com
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

func main() {
	root := NewTestCommand()
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//go:generate mockgen -package mocks -destination test/mocks/mock_test_interface.go github.com/dfreilich/gophercon-cli TestInterface
type TestInterface interface {
	DoSomething(args *DoSomethingArgs) error
}

type DoSomethingArgs struct {
	Message string
}

func NewCommandWithInterface(t TestInterface) *cobra.Command {
	return &cobra.Command{
		Use:     "test",
		Version: "0.0.1",
		Short:   "Test uses a library to abstract out",
		RunE: func(cmd *cobra.Command, args []string) error {
			return t.DoSomething(&DoSomethingArgs{
				Message: "Hello Gophercon!",
			})
		},
	}
}
