package server

import (
	"github.com/brandonsides/gpt/server/router"
)

type Config struct {
	Port   uint          `yaml:"port"`
	Router router.Config `yaml:"router"`
}
