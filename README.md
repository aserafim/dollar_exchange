# Dollar Exchange API

Este projeto consiste em uma aplicação desenvolvida em Go (Golang) composta por dois serviços: um **servidor HTTP** que consulta a cotação do dólar na API pública [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas) e armazena no banco SQLite, e um **cliente** que consome essa API local e salva a cotação em um arquivo `.txt`.

## 🗂️ Estrutura do Projeto

├── bkp # Backups e versões anteriores
├── client # Cliente que consome a API local e gera um arquivo com a cotação
├── db # Banco de dados SQLite
├── out # Saída dos arquivos gerados pelo cliente
├── server # Servidor HTTP que fornece a cotação
├── testes # Arquivos de teste e scripts auxiliares


## 🚀 Funcionalidades

- ✅ Servidor local exposto na porta `8080` com o endpoint `/cotacao`.
- ✅ Consulta a cotação do dólar (`USD-BRL`) na AwesomeAPI.
- ✅ Persiste o JSON completo retornado da API em um banco SQLite (`logs`).
- ✅ Cliente que consome o endpoint local `/cotacao` e gera um arquivo `cotacao.txt` contendo o valor do dólar.

## 🛠️ Tecnologias

- [Golang](https://golang.org/)
- [SQLite](https://www.sqlite.org/)
- [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas)

## 🔧 Pré-requisitos

- Go instalado (versão 1.18 ou superior)
- Acesso à internet (para consumo da API AwesomeAPI)

## 🏗️ Instalação e Execução

### 1️⃣ Clone o repositório:

```bash
git clone https://github.com/seu-usuario/dollar_exchange.git
cd dollar_exchange
```

### 2️⃣ Instale as dependências:

```bash
go mod tidy
```

### 3️⃣ Execute o servidor:

```bash
cd server
go run main.go
```
O servidor estará rodando em http://localhost:8080/cotacao.

### 4️⃣ Execute o cliente em outro terminal:

```bash
cd client
go run main.go
```

### ✔️ Resultado

```bash
out/cotacao.txt
```
Com o seguinte conteúdo (exemplo): Dólar: 5.7559

Além disso, o banco SQLite (db/db.db) armazenará o JSON completo da cotação.


### 📦 Banco de Dados
O banco de dados SQLite (db.db) possui uma tabela chamada logs:

```sql
CREATE TABLE logs (
    idLog TEXT PRIMARY KEY,
    cot TEXT
);
```
Cada requisição ao endpoint /cotacao salva uma entrada no banco, armazenando o JSON retornado da AwesomeAPI.

### 🔗 Endpoints

**Método**   | **Endpoint**   | **Descrição**
--------- | ------  | ------
GET | /cotacao | Retorna o valor atual do dólar no formato JSON: {"bid": "5.7559"}


### ⚠️ Limitações e Observações
- O servidor implementa timeout para evitar requisições travadas tanto no acesso à API externa quanto na gravação no banco.
- O cliente também possui timeout configurado.
- A persistência no banco de dados é simples e baseada no JSON bruto retornado da API.
- Projeto didático, ideal para aprendizado de Go, HTTP, SQLite e integração com APIs.
