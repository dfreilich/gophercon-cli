package cmd

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/stretchr/testify/require"

	"github.com/dfreilich/gophercon-cli/test/mocks"
)

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
