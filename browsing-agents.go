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
	"github.com/tmc/langchaingo/tools/serpapi"
)

func BrowsingAgents() {
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

	Seropt := serpapi.WithAPIKey(os.Getenv("SERPAPI_API_KEY"))
	search, err := serpapi.New(Seropt)

	if err != nil {
		log.Fatal(err)
	}

	agentTools := []tools.Tool{
		search,
	}

	agent := agents.NewOneShotAgent(llm, agentTools, agents.WithMaxIterations(3))

	executor := agents.NewExecutor(agent)

	query := "Who won the arm wrestling world tournament in 2024"

	result, err := executor.Call(ctx, map[string]interface{}{
		"input": query,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result["output"])
}
