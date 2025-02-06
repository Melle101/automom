package models

import (
	avanza_api "github.com/Melle101/mom-bot-v3/avanza-api"
)

// Config
type Config struct {
	Settings Settings `toml:"settings"`
	Assets   []Asset  `toml:"assets"`
}

type Settings struct {
	AGG             int    `toml:"AGG"`
	LookbackPeriod  string `toml:"lookbackPeriod"`
	SMAFilterLength int    `toml:"SMAFilterLength"`
	BackupAssetID   string `toml:"backupAsset"`
	BackupType      string `toml:"backupType"`
	HoldPeriod      int    `toml:"holdPeriod"`
	HoldPeriodType  string `toml:"holdPeriodType"`
	AccountURL      string `toml:"accountURL"`
	AccountID       string `toml:"accountID"`
}

type Asset struct {
	UnderlyingName string `toml:"asset"`
	UnderlyingID   string `toml:"assetID"`
	AssetType      string `toml:"assetType"`
	TargetLev      int    `toml:"targetLev"`
}

// Client
type Client struct {
	ApiClient avanza_api.ApiClient
	Cfg       Config
}

// AccountPositions
type AccountPositions struct {
	AccountURL   string
	AccountID    string
	CashPosition CashPosition
	Positions    []Position
}

type CashPosition struct {
	Value    float64
	Currency string
}

type Position struct {
	ID                      string
	AssetType               string
	AssetName               string
	OrderbookID             string
	Currency                string
	Quantity                float64
	CurrentValue            float64
	AcquisitionValue        float64
	AverageAcquisitionPrice float64
	UnderlyingID            string
}

// TreadeInfo
type OrderInfo struct {
	UnderlyingName string
	UnderlyingID   string
	AssetID        string
	OrderType      string
	Quantity       float64
	AssetType      string
	AccountID      string
}

type TradesInfo struct {
	Sells           []OrderInfo
	Buys            []OrderInfo
	BackupPositions int
}

type NewOrder struct {
	IsDividendReinvestment bool
	Price                  float64
	Volume                 int
	AccountID              string
	Side                   string
	OrderbookID            string
	Condition              string
}
