package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

/*
const (
	limitRequestAPIDollar = 200
)
*/

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

func GetDollPrice(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var d DollPrice
	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}

	// Envia o JSON no ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)

	//decoder := json.NewDecoder(r.Body)
	//fmt.Print(d)

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// io.Copy(os.Stdout, res.Body)
	//res, err := http.DefaultClient.Do(req)

	// 	if error != nil {
	// 		return nil, error
	// 	}
	// 	var c ViaCEP
	// 	error = json.Unmarshal(body, &c)
	// 	if error != nil {
	// 		return nil, error
	// 	}

	// 	cache[cep] = c

	// 	return &c, nil
	// }

}

func main() {
	http.HandleFunc("/", GetDollPrice)
	http.ListenAndServe(":8082", nil)

	// req, err := http.NewRequest("GET", "https://economia.awesomeapi.com.br/json/last/USD-BR", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// io.Copy(os.Stdout, res.Body)
}
