package openai

import (
	"context"

	openaiPkg "github.com/sashabaranov/go-openai"
)

type OpenAILanguageGenerator struct {
	model     string
	client    *openaiPkg.Client
	maxTokens int
}

func NewOpenAILanguageGenerator(config Config) (OpenAILanguageGenerator, error) {
	return OpenAILanguageGenerator{
		model:     config.Model,
		client:    openaiPkg.NewClient(config.APIKey),
		maxTokens: config.MaxTokens,
	}, nil
}

func (o OpenAILanguageGenerator) Generate(ctx context.Context, prompt string) (string, error) {
	resp, err := o.client.CreateCompletion(ctx, openaiPkg.CompletionRequest{
		Model:     o.model,
		Prompt:    prompt,
		MaxTokens: o.maxTokens,
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}
