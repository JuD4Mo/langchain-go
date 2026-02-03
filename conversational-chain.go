package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
)

func ConversationalChain() {
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

	bufferMemory := memory.NewConversationBuffer()

	conversationChain := chains.NewConversation(llm, bufferMemory)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nUser:")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
		}

		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("Warning! please enter a message.")
			continue
		}

		fmt.Print("Robot Agent: ")
		fmt.Print("thinking...")

		response, err := chains.Run(ctx, conversationChain, userInput)

		fmt.Print("\rRobor Agent: ")
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println(response)
	}
}
