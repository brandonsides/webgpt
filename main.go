// This file sets up the server and routes
package main

import (
	"os"

	"github.com/brandonsides/gpt/langgen"
	"github.com/brandonsides/gpt/server"
	"gopkg.in/yaml.v3"
)

func main() {
	// load config
	config, err := loadConfig(os.Getenv("GPT_CONFIG"))
	if err != nil {
		panic(err)
	}

	// create language generator
	languageGenerator, err := langgen.NewLanguageGenerator(config.LanguageGenerator)
	if err != nil {
		panic(err)
	}

	// create server
	server := server.NewServerFromConfig(config.Server, languageGenerator)

	// start server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func loadConfig(path string) (Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	var config Config
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
