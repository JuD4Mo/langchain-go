package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms/fake"
)

func UsingFakesLlms() {
	res := []string{
		"Hello!",
		"Nahgh",
		"Omori goat",
	}

	llm := fake.NewFakeLLM(res)

	ctx := context.Background()

	ans, _ := llm.Call(ctx, "Hi there")

	fmt.Println(ans)
}
