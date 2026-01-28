package main

import (
	"fmt"
	"time"

	"github.com/tmc/langchaingo/prompts"
)

// In InputVariables we just define company_name 'cause date is a partial variable
func partialVariablesInPrompts() {
	templateWithPartials := prompts.PromptTemplate{
		Template:       "I want the {{company_name}} financial breakdown from {{date}}",
		InputVariables: []string{"company_name"},
		PartialVariables: map[string]interface{}{
			"date": func() string {
				return time.Now().Format("2006-01-02")
			},
		},
		TemplateFormat: prompts.TemplateFormatJinja2,
	}

	/*
		Along the line

		templateWithPartials.PartialVariables = map[string]interface{}{
				"date": "2026-01-25"
				}
	*/

	formatted_template, _ := templateWithPartials.Format(map[string]interface{}{
		"company_name": "FZ Pharmaceuticals",
	})

	fmt.Println(formatted_template)
}
