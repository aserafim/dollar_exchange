package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	db, err := sql.Open("sqlite3", "../db/db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria a tabela "logs" se ela não existir
	createTableSQL := `CREATE TABLE IF NOT EXISTS logs (
		"idLog" TEXT NOT NULL PRIMARY KEY,
		"cot" TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erro ao criar a tabela logs:", err)
	}

	// Insere um registro na tabela
	stmt, err := db.Prepare("INSERT INTO logs(idLog, cot) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("123", "teste")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Registro inserido com sucesso")
}
