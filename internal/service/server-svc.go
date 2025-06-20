package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var client = &http.Client{}
var timeoutPersistenceDatabase = time.Millisecond * 10
var timeoutRequestApi = time.Millisecond * 2000

type ServerSvc struct{}

func NewServerSvc() *ServerSvc {
	return &ServerSvc{}
}

func (s *ServerSvc) GetUsdExchangeRate() (ApiUsdExchangeRateResponse, error) {
	exchangeRateUsdToBrl, err := getCurrentUsdExchange()
	if err != nil {
		return ApiUsdExchangeRateResponse{}, fmt.Errorf("error getting current USD exchange rate: %w", err)
	}

	db, err := initDB()
	if err != nil {
		return ApiUsdExchangeRateResponse{}, fmt.Errorf("error initializing database: %w", err)
	}
	defer db.Close()

	err = s.PersistUsdExchangeRate(db, exchangeRateUsdToBrl, timeoutPersistenceDatabase)
	if err != nil {
		return ApiUsdExchangeRateResponse{}, fmt.Errorf("error persisting USD exchange rate: %w", err)
	}

	return ApiUsdExchangeRateResponse{
		Bid: exchangeRateUsdToBrl.Bid,
	}, nil
}

func (s *ServerSvc) PersistUsdExchangeRate(db *sql.DB, exchangeRate ApiUsdExchangeRate, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	insertSQL := `INSERT INTO exchange_rates (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.ExecContext(ctx, insertSQL,
		exchangeRate.Code,
		exchangeRate.CodeIn,
		exchangeRate.Name,
		exchangeRate.High,
		exchangeRate.Low,
		exchangeRate.VarBid,
		exchangeRate.PctChange,
		exchangeRate.Bid,
		exchangeRate.Ask,
		exchangeRate.Timestamp,
		exchangeRate.CreateDate)

	if err != nil {
		fmt.Printf("Insert failed: %v\n", err)
		return fmt.Errorf("error inserting exchange rate into database: %w", err)
	}

	return nil
}

func initDB() (*sql.DB, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return nil, err
	}

	dbDir := filepath.Join(rootPath, "internal", "db")
	err = os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create db directory: %v\n", err)
		return nil, err
	}

	dbPath := filepath.Join(dbDir, "exchange_rates.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return nil, err
	}

	// Create table EXCHANGE_RATES if not exists
	createTableSQL := `CREATE TABLE IF NOT EXISTS exchange_rates (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return nil, err
	}

	return db, nil
}

func getCurrentUsdExchange() (ApiUsdExchangeRate, error) {
	var apiResponse ApiUsdExchangeRateWrapper
	// time out of 200milliseconds to make the request
	ctx, cancel := context.WithTimeout(context.Background(), timeoutRequestApi)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ApiUsdExchangeRateUrl, nil)
	if err != nil {
		fmt.Printf("Error creating new request: %v\n", err)
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making GET request to API: %v\n", err)
		return ApiUsdExchangeRate{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return ApiUsdExchangeRate{}, err
	}

	// fmt.Printf("\n 🔵 Response: %s", body)

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("Error decoding API response: %v\n", err)
		return ApiUsdExchangeRate{}, err
	}

	return apiResponse.USDBRL, nil
}
