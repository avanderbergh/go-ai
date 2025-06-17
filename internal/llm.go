package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/ollama/ollama/api"
)

const model = "avanderbergh/gemma3-translator:4b"

func Translate(text string) string {
	client, err := api.ClientFromEnvironment()

	if err != nil {
		panic(err)
	}

	req := &api.GenerateRequest{
		Model:  model,
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

	translated = strings.TrimSpace(translated)

	return translated
}
