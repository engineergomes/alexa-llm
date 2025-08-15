package openai

import (
	"context"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func Chat(message string) string {
	key, _ := os.LookupEnv("OPENAI_API_KEY")

	client := openai.NewClient(
		option.WithAPIKey(key), 
	)

	
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	})
	if err != nil {
		panic(err.Error())
	}
	return chatCompletion.Choices[0].Message.Content
}
