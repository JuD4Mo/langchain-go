package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func standardTemplateDefinition() {
	templateWithProps := prompts.PromptTemplate{
		Template:       "Research {{.topic}} on {{.website}}",
		InputVariables: []string{"topic", "website"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	template_with_props, _ := templateWithProps.Format(map[string]interface{}{
		"topic":   "Cookie recipes",
		"website": "The food Network",
	})

	fmt.Println(template_with_props)
}
