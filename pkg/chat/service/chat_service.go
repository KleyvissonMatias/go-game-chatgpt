package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KleyvissonMatias/go-game-chatgpt/internal/constant"
	"github.com/KleyvissonMatias/go-game-chatgpt/pkg/chat/model/request"
	chat "github.com/KleyvissonMatias/go-game-chatgpt/pkg/chat/model/response"
)

func Question(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	requestBody, _ := io.ReadAll(r.Body)
	var chatGPTRequest request.ChatGPTRequest
	err := json.Unmarshal(requestBody, &chatGPTRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
		return
	}
	question(chatGPTRequest.Texto)
}

func question(texto string) {
	var response chat.ChatGPTResponse

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      texto,
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
	req.Header.Set("Authorization", "Bearer "+os.Getenv(constant.ENV_API_KEY))

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
}
