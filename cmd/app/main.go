package main

import (
	"go.uber.org/fx"
	"log"
	"url-shortener/internal/di"
	"url-shortener/pkg/env"
)

func main() {
	if err := env.NewEnv(".env"); err != nil {
		log.Fatal(err)
	}
	fx.New(di.NewModule()).Run()
}
