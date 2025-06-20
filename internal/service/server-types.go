package service

var ApiUsdExchangeRateUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type ApiUsdExchangeRateWrapper struct {
	USDBRL ApiUsdExchangeRate `json:"USDBRL"`
}

type ApiUsdExchangeRate struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ApiUsdExchangeRateResponse struct {
	Bid string
}

/*
{
  "USDBRL": {
    "code": "USD",
    "codein": "BRL",
    "name": "Dólar Americano/Real Brasileiro",
    "high": "5.4952",
    "low": "5.48259",
    "varBid": "0.00405",
    "pctChange": "0.073806",
    "bid": "5.4914",
    "ask": "5.4944",
    "timestamp": "1750402897",
    "create_date": "2025-06-20 04:01:37"
  }
}*/
