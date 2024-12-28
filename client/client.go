package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const urlServer string = "http://localhost:8080"

type ResponseDTO struct {
	Valor string `json:"valor"`
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlServer+"/cotacao", nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		panic(string(body))
	}

	var cotacao *ResponseDTO

	json.Unmarshal(body, &cotacao)

	// fmt.Println(cotacao)

	file, err := os.Create("cotacao.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar:%s", cotacao.Valor))

	if err != nil {
		panic(err)
	}
}
