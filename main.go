package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

//go:embed text.txt
var f embed.FS

func main() {

	data, _ := f.ReadFile("text.txt")

	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel("llama3"))

	if err != nil {
		log.Fatal(err)
	}
	prompt := "make a unit test Go function for function " + string(data)

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
