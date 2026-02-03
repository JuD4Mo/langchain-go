package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/tools"
)

func ReActAgent() {
	//type OneShotAgent

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

	calculator := tools.Calculator{}

	agentTools := []tools.Tool{
		calculator,
	}

	agent := agents.NewOneShotAgent(llm, agentTools, agents.WithMaxIterations(5))

	executor := agents.NewExecutor(agent)

	mathProblem := "What is 25 * 47 plus the square root of 144"

	result, err := executor.Call(ctx, map[string]interface{}{
		"input": mathProblem,
	})

	if err != nil {
		log.Printf("error: %v\n", err)
	} else {
		fmt.Println(result["output"])
	}
}
