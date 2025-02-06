package avanza_api

type ApiClient struct {
	CredInfo    CredInfo
	CredHeaders map[string][]string
}

type InitAuthInfo struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	TOTP     string `toml:"TOTP"`
}

type CredInfo struct {
	Secret_token string
	Auth_session string
	Push_sub_id  string
	Cs_token     string
}

type MarketOrderResult struct {
	Suceess   bool
	AssetName string
	AssetID   string
	Quantity  float64
	Price     float64
}

// WarrantMarketORder
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
