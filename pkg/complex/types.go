package complex

import "fmt"

//go:generate mockgen -package mocks -destination ../../test/mocks/mock_complex_actor.go github.com/dfreilich/gophercon-cli/pkg/complex Actor
type Actor interface {
	DoSomething(args *DoSomethingArgs) error
}

type DoSomethingArgs struct {
	Message string
}

type MyActor struct{}

func (m *MyActor) DoSomething(args *DoSomethingArgs) error {
	fmt.Println(args.Message)
	return nil
}
