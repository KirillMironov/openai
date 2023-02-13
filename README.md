# OpenAI Go Client

- [OpenAI API Reference](https://platform.openai.com/docs/api-reference)
- [OpenAI OpenAPI Specification](https://github.com/openai/openai-openapi)

## Installation
```shell
go get github.com/KirillMironov/openai
```

## Usage
```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/KirillMironov/openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(apiKey, openai.WithTimeout(time.Second*20))

	completion, err := client.Completion(context.Background(), openai.CompletionRequest{
		Model:     "text-davinci-003",
		Prompt:    []string{"Example prompt"},
		MaxTokens: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(completion.Choices[0].Text)
}
```
