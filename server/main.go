package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
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

type Cotacao struct {
	IDCotacao string
	Cotacao   string
}

func NewCotacao(_Cotacao string) *Cotacao {
	return &Cotacao{
		IDCotacao: uuid.New().String(),
		Cotacao:   _Cotacao,
	}
}

func writeLog(db *sql.DB, cot *Cotacao) error {
	// Criar um contexto com timeout de 3 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Preparar a instrução usando o contexto
	stmt, err := db.PrepareContext(ctx, "INSERT INTO logs(idLog, cot) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executar a instrução com os valores de `cot` e o contexto
	_, err = stmt.ExecContext(ctx, cot.IDCotacao, cot.Cotacao)
	if err != nil {
		return err
	}

	return nil
}

func createTable(db *sql.DB) error {

	// Cria a tabela "logs" se ela não existir
	createTableSQL := `CREATE TABLE IF NOT EXISTS logs (
		"idLog" TEXT NOT NULL PRIMARY KEY,
		"cot" TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func GetDollPrice(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "/home/aserafim/dev-repos/go-env/dollar_exchange/db/db.db")  
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	createTable(db)

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

	//verificar se vamos usar
	cot := NewCotacao(string(body))
	writeLog(db, cot)
	//writeLog(db, cot, ctxDB)

	var d DollPrice
	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}

	bid := struct {
		Bid string `json:"bid"`
	}{
		Bid: d.USDBRL.Bid,
	}

	// Envia o JSON no ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bid)

}

func main() {
	http.HandleFunc("/cotacao", GetDollPrice)
	http.ListenAndServe(":8080", nil)

}
