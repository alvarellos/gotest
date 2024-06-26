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
	//print(string(data))

	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel("llama3"))
	//llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	prompt := "make a unit test Go function for function " + string(data)
	//prompt := "make a Go unit test function for function " + string(data)
	//prompt := "return in plain text format the next text" + string(data)
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
