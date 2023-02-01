# gophercon-cli

Tutorial: 
1. Initialize `go.mod`: `go mod init github.com/dfreilich/gophercon-cli`
2. Create a cmd directory: `mkdir -p cmd/`
3. Get cobra: `go get github.com/spf13/cobra`
4. Create a file in the `cmd/` directory with the contents: 
```go
var (
	Version = "0.0.1"
)

func NewJokerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "joker",
		Aliases: []string{"joke"},
		Short:   "Joker!",
		Long:    "This returns ChatGPT jokes!",
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello Gophercon!")
			return nil
		},
	}
}
```
5. Initialize the `main.go` file to run the command: 
```go
func main() {
	root := cmd.NewJokerCmd()
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```
6. Create a `Makefile` with the contents:
```
GOCMD?=go
NAME?=mycli

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

.PHONY: build test
```
7. Get library for `ChatGPT`: `go get github.com/sashabaranov/go-gpt3`
8. Use ChatGPT in `root.go`: 
```go
c := gogpt.NewClient(apiKey)
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
resp, err := c.CreateCompletion(ctx, req)
if err != nil {
    return err
}
```
10. Get a nice rendering library: `go get github.com/charmbracelet/lipgloss`
11. Make the output :sparkles: fabulous :sparkles: :
```go
style := lipgloss.NewStyle().
    Bold(true).
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("228")).
    Foreground(lipgloss.Color("12"))

fmt.Println(style.Render(strings.TrimSpace(resp.Choices[0].Text)))
```