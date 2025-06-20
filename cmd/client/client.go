package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var apiUrl string = "http://localhost:8080/cotacao"
var client = &http.Client{}
var timeoutRequestApi = time.Millisecond * 300

type ApiResponseWrapper struct {
	ApiData    ApiDataResponse `json:"data"`
	ApiMessage string          `json:"message"`
	ApiStatus  string          `json:"status"`
}

type ApiDataResponse struct {
	Dolar string `json:"Bid"`
}

func main() {
	var response ApiResponseWrapper
	ctx, cancel := context.WithTimeout(context.Background(), timeoutRequestApi)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}
	println("Dólar:", response.ApiData.Dolar)

	file := getFile()
	defer file.Close()

	// Write/append the usd value to the file
	_, err = file.WriteString("Dólar: " + response.ApiData.Dolar + "\n")
	if err != nil {
		panic(err)
	}
}

func getFile() *os.File {
	var file *os.File
	filePath := "cotacao.txt"

	// Check if the file exists, if not create it
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}
	return file

}
