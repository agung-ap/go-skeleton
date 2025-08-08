package config

import "github.com/spf13/viper"

type AppConfig struct {
	DocsPath string
}

var App AppConfig

func initAppConfig() {
	// Make DOCS_PATH optional with a sensible default to avoid panics during tests
	docsPath := viper.GetString("DOCS_PATH")
	if docsPath == "" {
		docsPath = "./docs"
	}
	App.DocsPath = docsPath
}
