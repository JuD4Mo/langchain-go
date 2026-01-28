package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func jinjaPromptTemplates() {
	//The template has to be in jinja syntaxt if the template format is TemplateFormatJinja2
	templateWithJinja := prompts.PromptTemplate{
		Template:       "Translate '{{statement}}' from {{lang1}} to {{lang2}}",
		InputVariables: []string{"statement", "lang1", "lang2"},
		TemplateFormat: prompts.TemplateFormatJinja2,
	}

	s, _ := templateWithJinja.Format(map[string]interface{}{
		"statement": "Welcome to China",
		"lang1":     "English",
		"lang2":     "Chinese",
	})

	fmt.Println(s)
}
