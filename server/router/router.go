package router

import (
	"net/http"

	"github.com/brandonsides/gpt/langgen"
)

func NewRouter(conf Config, lg langgen.LanguageGenerator) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text, err := lg.Generate(r.Context(), conf.Prompt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(text))
	})

	return router
}
