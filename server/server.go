package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CotacaoDTO struct {
	Usdbrl struct {
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

type ResponseDTO struct {
	Valor string `json:"valor"`
}

func main() {

	http.HandleFunc("/cotacao", handleCotacao)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	data, err := BuscaCotacao(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar cotação"))
		fmt.Println(err)
		return
	}

	response := &ResponseDTO{
		Valor: data.Usdbrl.Bid,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func BuscaCotacao(ctx context.Context) (*CotacaoDTO, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data CotacaoDTO
	err = json.Unmarshal(res, &data)

	if err != nil {
		return nil, err
	}

	err = SalvarCotacao(ctx, &data)

	return &data, err
}

func SalvarCotacao(ctx context.Context, data *CotacaoDTO) error {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	db, err := sql.Open("sqlite3", "cotacao.db")

	if err != nil {
		return err
	}

	defer db.Close()

	err = MigrateTables(db)

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("insert into cotacoes (cotacao, dt_consulta) values ($1, $2)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, data.Usdbrl.Bid, time.Now())

	return err
}

func MigrateTables(db *sql.DB) error {
	const create string = `
		create table if not exists cotacoes (
			id integer primary key autoincrement,
			cotacao number not null,
			dt_consulta timestamp not null
		)
	`

	_, err := db.Exec(create)

	return err
}
