# gophercon-cli

## 1. Create initial functionality
1. Initialize `go.mod`: `go mod init github.com/dfreilich/gophercon-cli`
2. Create a cmd directory: `mkdir -p cmd/`
3. Get cobra: `go get github.com/spf13/cobra`
4. Create a file in the `cmd/` directory with the contents: 
```go
func NewJokerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "joker",
		Aliases: []string{"joke"},
		Short:   "This returns GPT3 Dad jokes!",
		Version: "0.0.1",
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
NAME?=joker

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

run:
	$(GOCMD) run ./...

.PHONY: build test run
```
7. Write initial test in `root_test.go`:
```go
func TestNewJokerCmd(t *testing.T) {
	cmd := NewJokerCmd()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	err := cmd.Execute()
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
}
```

## 2. Sprinkle Some AI
8. Get library for `ChatGPT`: `go get github.com/sashabaranov/go-gpt3`
9. Make `API Key` at https://beta.openai.com/account/api-keys and set it as an environment variable `export OPEN_AI_KEY=MY_KEY`
9. Use ChatGPT in `root.go`: 
```go
// Note: For this, you need to make an API KEY at https://beta.openai.com/account/api-keys
c := gogpt.NewClient(os.Getenv("OPEN_AI_KEY"))
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

## 3. Add Some Style
10. Get a nice rendering library: `go get github.com/charmbracelet/lipgloss`
11. Make the output :sparkles: fabulous :sparkles: :
```go
style := lipgloss.NewStyle().
    Bold(true).
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("12")). // Light Blue
    Foreground(lipgloss.Color("5")) // Magenta

fmt.Println(style.Render(strings.TrimSpace(resp.Choices[0].Text)))
```

## 4. Add Mocking/Interface
12. Go get assertion library: `go get github.com/stretchr/testify/require`
14. Install `mockgen` : `go install github.com/golang/mock/mockgen@v1.6.0` and `go get github.com/golang/mock/mockgen/model`
15. Create the interface and annotations:
```go
//go:generate mockgen -package mocks -destination ../test/mocks/mock_asker.go github.com/dfreilich/gophercon-cli/cmd Asker
type Asker interface {
	CreateCompletion(ctx context.Context, request gogpt.CompletionRequest) (response gogpt.CompletionResponse, err error)
}
```
16. Change the Command to use it: 
```go
func NewJokerCmd(asker Asker) *cobra.Command {
...
    resp, err := asker.CreateCompletion(ctx, req)
```
17. Change the main to se it: 
```go
	c := gogpt.NewClient(os.Getenv("OPEN_AI_KEY"))
	root := cmd.NewJokerCmd(c)
```
18. Change the test to use it:
```go
func TestNewJokerCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	testActor := mocks.NewMockAsker(ctrl)
	cmd := NewJokerCmd(testActor)
	testActor.EXPECT().CreateCompletion(gomock.Any(), gomock.Any()).Return(gogpt.CompletionResponse{
		Choices: []gogpt.CompletionChoice{{Text: "Some funny joke!"}},
	}, nil)
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	err := cmd.Execute()
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "Some funny joke!")
}
```

## 5. Run
![asciicast](test_output.gif)