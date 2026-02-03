package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/tools"
)

type wordCounterTool struct {
}

// The Tool interface has 3 functions: Name(), Description(), Call()

func (wc wordCounterTool) Name() string {
	return "WordCounter"
}

// The more we describe the tool, the better the agent will implement it

func (wc wordCounterTool) Description() string {
	return "Count words in text. Input: the text to count"
}

func (wc wordCounterTool) Call(ctx context.Context, input string) (string, error) {
	count := len(strings.Fields(input))

	return fmt.Sprintf("%d", count), nil
}

func UsingCustomTools() {
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

	//Gather tools

	calculator := tools.Calculator{}
	wordCounter := wordCounterTool{}

	agentTools := []tools.Tool{calculator, wordCounter}

	agent := agents.NewOneShotAgent(llm, agentTools)

	executor := agents.NewExecutor(agent)

	mathPrompt := "Count the words in 'Omori is the best game' and multiply it by the square root of 4"

	resesut, err := executor.Call(ctx, map[string]any{
		"input": mathPrompt,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resesut["output"])
}
