package cmd

import (
	"context"
	"strings"

	"github.com/charmbracelet/lipgloss"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
)

var (
	Version   = "0.0.1"
	maxTokens = 150
)

//go:generate mockgen -package mocks -destination ../test/mocks/mock_asker.go github.com/dfreilich/gophercon-cli/cmd Asker
type Asker interface {
	CreateCompletion(
		ctx context.Context,
		request gogpt.CompletionRequest,
	) (response gogpt.CompletionResponse, err error)
}

func NewJokerCmd(asker Asker) *cobra.Command {
	return &cobra.Command{
		Use:     "joker",
		Aliases: []string{"joke"},
		Short:   "This returns GPT3 Dad jokes!",
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			req := gogpt.CompletionRequest{
				Model:            gogpt.GPT3TextDavinci003,
				MaxTokens:        maxTokens,
				Temperature:      1,
				Prompt:           "Tell me a corny dad joke",
				TopP:             1,
				FrequencyPenalty: 1,
				PresencePenalty:  1,
			}
			resp, err := asker.CreateCompletion(ctx, req)
			if err != nil {
				return err
			}

			style := lipgloss.NewStyle().
				Bold(true).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("12")). // Light Blue
				Foreground(lipgloss.Color("5"))         // Magenta

			cmd.Println(style.Render(strings.TrimSpace(resp.Choices[0].Text)))
			return nil
		},
	}
}
