package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/prompts"
)

func LlmChains() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	opt := googleai.WithAPIKey(os.Getenv("GEMINI_API_KEY"))
	opt2 := googleai.WithDefaultModel(os.Getenv("AI_MODEL"))

	llm, err := googleai.New(ctx, opt, opt2)
	if err != nil {
		log.Fatal(err)
	}

	templateString := `
		You're a helpful assistant that translates text.
		Transale the following text from {{.lang1}} to {{.lang2}}

		Text: {{.text}}
		Translation:
	`

	prompt := prompts.NewPromptTemplate(
		templateString,
		[]string{"lang1", "lang2", "text"},
	)

	llmChain := chains.NewLLMChain(llm, prompt)

	chain_input := map[string]interface{}{
		"lang1": "English",
		"lang2": "French",
		"text":  "Well done!",
	}

	//Run(): single input, single output
	//Predict(): multiple input, single output
	//Call(): multiple input, multiple output

	result, err := chains.Predict(ctx, llmChain, chain_input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
