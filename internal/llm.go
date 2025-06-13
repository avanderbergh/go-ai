package internal

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func Translate(text string) string {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		panic(err)
	}

	req := &api.GenerateRequest{
		Model:  "gemma3",
		Prompt: Prompt + text,
		Stream: new(bool),
	}

	ctx := context.Background()

	var translated string

	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)
		translated = resp.Response
		return nil
	}

	client.Generate(ctx, req, respFunc)

	return translated

}
