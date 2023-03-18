package server

import (
	"fmt"
	"net/http"

	"github.com/brandonsides/gpt/langgen"
	"github.com/brandonsides/gpt/server/router"
)

func NewServerFromConfig(config Config, lg langgen.LanguageGenerator) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router.NewRouter(config.Router, lg),
	}
}
