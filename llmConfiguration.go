package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func UsingModelConfigurations() {

	backgroundContext := context.Background()

	apiKey := "xxx-xx-xxxx"

	llm, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithModel("gpt-5"),
	)

	if err != nil {
		log.Fatal(err)
	}

	response, err := llms.GenerateFromSinglePrompt(
		backgroundContext,
		llm,
		"Who invented the microphone",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

}
