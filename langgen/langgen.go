package langgen

import (
	"context"
	"fmt"

	"github.com/brandonsides/gpt/langgen/openai"
)

type LanguageGenerator interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

func NewLanguageGenerator(config Config) (LanguageGenerator, error) {
	switch config.Provider {
	case OpenAI:
		return openai.NewOpenAILanguageGenerator(config.ProviderConfig.(openai.Config))
	default:
		return nil, fmt.Errorf("unknown provider: %s", config.Provider)
	}
}
