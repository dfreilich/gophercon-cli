/*
Copyright © 2023 David Freilich freilich.david@gmail.com
*/
package main

import (
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"

	"github.com/dfreilich/gophercon-cli/cmd"
)

var (
	apiKey = os.Getenv("OPEN_AI_KEY")
)

func main() {
	c := gogpt.NewClient(apiKey)
	root := cmd.NewJokerCmd(c)
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
