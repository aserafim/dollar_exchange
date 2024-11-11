package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DollPrice struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Bid struct {
	Bid string `json:"bid"`
}

func writeFile(content string) error {

	file, err := os.Create("out/cotacao.txt")
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 100000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	// Executa a requisição HTTP
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Verifica se o erro foi devido a timeout
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout: o tempo limite foi excedido")
		} else {
			log.Printf("Erro na requisição: %v", err)
		}
		return
	}
	defer res.Body.Close()

	//io.Copy(os.Stdout, res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var b Bid	
	err = json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}

	ret := "Dólar: " + b.Bid

	err = writeFile(ret)
	if err != nil {
		panic(err)
	}

}
