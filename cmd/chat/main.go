package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KleyvissonMatias/go-game-chatgpt/internal/constant"
	chat "github.com/KleyvissonMatias/go-game-chatgpt/pkg/chat/model/response"
)

func main() {

	apiKey := os.Getenv(constant.ENV_API_KEY)

	var response chat.ChatGPTResponse

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "Brasil",
		"max_tokens":  4000,
		"temperature": 1.0,
	})

	if err != nil {
		fmt.Println("Erro ao criar o corpo da solicitação JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Erro ao criar a solicitação HTTP:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler response", err.Error())
	}

	if err := json.Unmarshal([]byte(string(body)), &response); err != nil {
		fmt.Println("Erro ao fazer unmarshal do JSON:", err)
		return
	}

	fmt.Println("Warning:", response.Warning)
	fmt.Println("ID:", response.ID)
	fmt.Println("Object:", response.Object)
	fmt.Println("Created:", response.Created)
	fmt.Println("Model:", response.Model)

	fmt.Println("Choices:")
	for _, choice := range response.Choices {
		fmt.Println("  Text:", choice.Text)
		fmt.Println("  Index:", choice.Index)
		fmt.Println("  Logprobs:", choice.Logprobs)
		fmt.Println("  FinishReason:", choice.FinishReason)
	}

	fmt.Println("Usage:")
	fmt.Println("  PromptTokens:", response.Usage.PromptTokens)
	fmt.Println("  CompletionTokens:", response.Usage.CompletionTokens)
	fmt.Println("  TotalTokens:", response.Usage.TotalTokens)
}
