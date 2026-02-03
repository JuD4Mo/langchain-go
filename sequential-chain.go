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

/*
	- Synopsis Chain: take a "topic" and generate a "story"
	- Character Chain: take our "story" and extract the main "character"
	- Backstory Chain: take the "character" and build a "backstory"

	"topic" -> SynopsisChain -> story -> CharacterChain -> character -> BackstoryChain -> backstory
*/

func SequentialChains() {
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

	// Synopsis chain
	synopsisTemplate := `
		Create a short story synopsis about {{.input}}
		Synopsis:
	`

	synopsisPrompt := prompts.NewPromptTemplate(synopsisTemplate, []string{"input"})

	synopsisChain := chains.NewLLMChain(llm, synopsisPrompt)

	// Character chain
	characterTemplate := `
		Based on this story synopsis, identify and describe the main character.
		Synopsis: {{.input}}
		Main character:
	`

	characterPrompt := prompts.NewPromptTemplate(characterTemplate, []string{"input"})

	characterChain := chains.NewLLMChain(llm, characterPrompt)

	// Backstory chain
	backstoryTemplate := `
		Create a detailed backstory for this character.
		Character description: {{.input}}
		Backstory:
	`

	backstoryPrompt := prompts.NewPromptTemplate(backstoryTemplate, []string{"input"})

	backstoryChain := chains.NewLLMChain(llm, backstoryPrompt)

	sequentialChain, err := chains.NewSimpleSequentialChain([]chains.Chain{
		synopsisChain, characterChain, backstoryChain,
	})
	if err != nil {
		log.Fatal(err)
	}

	ans, err := chains.Run(ctx, sequentialChain, "A space explorer discovering a new planet")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
