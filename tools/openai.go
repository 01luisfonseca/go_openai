package tools

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

var apiKey string

func init() {
	LoadEnv()
	// Set the API key from the environment variable
	apiKey = GetEnv("OPENAI_API_KEY")
}

func GetCompletion(prompt string, model string) string {
	client := openai.NewClient(apiKey)
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Choices[0].Message.Content
}
