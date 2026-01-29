package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/prompts"
)

func chatPromptTemplate() {

	systemPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		"You're to give your responses in {{.language}}",
		[]string{"language"},
	)

	humanMessageTemplate := prompts.NewHumanMessagePromptTemplate(
		"Translate this: {{.phrase}}",
		[]string{"phrase"},
	)

	chatPromptTemplate := prompts.NewChatPromptTemplate(
		[]prompts.MessageFormatter{
			systemPromptTemplate,
			humanMessageTemplate,
		},
	)

	formattedChatPrompt, _ := chatPromptTemplate.Format(map[string]interface{}{
		"language": "Spanish",
		"phrase":   "Thank you",
	})

	// fmt.Println(formattedChatPrompt)

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

	res, err := llms.GenerateFromSinglePrompt(ctx, llm, formattedChatPrompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
