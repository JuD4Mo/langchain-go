package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func usingLocalModels() {
	ctx := context.Background()

	model, err := ollama.New(
		ollama.WithModel("gemma3"),
	)

	if err != nil {
		log.Fatal(err)
	}

	prompt := "What is Omori?"

	res, err := llms.GenerateFromSinglePrompt(ctx, model, prompt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
