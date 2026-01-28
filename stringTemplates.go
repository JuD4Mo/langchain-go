package main

import (
	"fmt"
	"log"

	"github.com/tmc/langchaingo/prompts"
)

func stringPrompTemplates() {
	simpleTemplate := prompts.NewPromptTemplate(
		"Write a {{.content_type}} about {{.subject}}",
		[]string{"content_type", "subject"},
	)

	templateInput := map[string]interface{}{
		"content_type": "poem",
		"subject":      "cats",
	}

	s, err := simpleTemplate.Format(templateInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
