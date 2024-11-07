package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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

	file, err := os.Create("/home/aserafim/dev-repos/go-env/dollar_exchange/out/cotacao.txt")
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
	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
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
