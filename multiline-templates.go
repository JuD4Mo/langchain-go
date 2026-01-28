package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func multilineTemplates() {
	templateString := `
		Display the following list of {{.profession}}

		{{.names}}

		{{if .display_examples}}
		Examples:
		{{range .display_examples}}
		- {{.}}

		{{end}}

		{{end}}
	`

	template_multiline := prompts.PromptTemplate{
		Template:       templateString,
		InputVariables: []string{"profession", "names", "display_examples"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	s, _ := template_multiline.Format(map[string]interface{}{
		"profession":       "Footballers",
		"names":            []string{"Ronaldo", "Messi", "Haaland"},
		"display_examples": []string{"kaka - AC Milan", "Rooney - Manchester United"},
	})

	fmt.Println(s)
}
