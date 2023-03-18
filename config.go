package main

import (
	"github.com/brandonsides/gpt/langgen"
	"github.com/brandonsides/gpt/server"
)

type Config struct {
	Server            server.Config  `yaml:"server"`
	LanguageGenerator langgen.Config `yaml:"language_generator"`
}
