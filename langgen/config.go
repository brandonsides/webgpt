package langgen

import (
	"fmt"

	"github.com/brandonsides/gpt/langgen/openai"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Provider       Provider    `yaml:"provider"`
	ProviderConfig interface{} `yaml:"config"`
}

func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	type plain struct {
		Provider       Provider  `yaml:"provider"`
		ProviderConfig yaml.Node `yaml:"config"`
	}
	var p plain
	if err := value.Decode(&p); err != nil {
		return err
	}

	c.Provider = p.Provider
	switch c.Provider {
	case OpenAI:
		var config openai.Config
		if err := p.ProviderConfig.Decode(&config); err != nil {
			return err
		}
		c.ProviderConfig = config
	default:
		return fmt.Errorf("unknown provider: %s", c.Provider)
	}

	return nil
}
