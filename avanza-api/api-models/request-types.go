package api_models

type WarrantSearch struct {
	Filter Filter `json:"filter"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	SortBy SortBy `json:"sortBy"`
}
type Filter struct {
	SubTypes              []string `json:"subTypes"`
	Issuers               []string `json:"issuers"`
	UnderlyingInstruments []string `json:"underlyingInstruments"`
	EndDates              []string `json:"endDates"`
	Directions            []string `json:"directions"`
	NameQuery             string   `json:"nameQuery,omitempty"`
}
type SortBy struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

// Place Order Info
type NewOrder struct {
	IsDividendReinvestment bool
	Price                  float64
	Volume                 int
	AccountId              string
	Side                   string
	OrderbookId            string
	Condition              string
	RequestId              string
	OpenVolume             *string
	OrderRequestParameters *string
	ValidUntil             *string
	Metadata               Metadata
}

type Metadata struct {
	HasTouchedPrice string `json:",omitempty"`
	OrderEntryMode  string
}

// Modify order
type ModifyOrderInfo struct {
	OrderId    string
	Price      float64
	Volume     int
	OpenVolume *int
	AccountId  string
	ValidUntil *string
	Metadata   Metadata
}

type ValidationRequest struct {
	IsDividendReinvestment bool    `json:"isDividendReinvestment"`
	RequestID              *any    `json:"requestId"`
	OrderRequestParameters *any    `json:"orderRequestParameters"`
	Price                  float64 `json:"price"`
	Volume                 int     `json:"volume"`
	OpenVolume             *any    `json:"openVolume"`
	AccountID              string  `json:"accountId"`
	Side                   string  `json:"side"`
	OrderbookID            string  `json:"orderbookId"`
	ValidUntil             *any    `json:"validUntil"`
	Metadata               *any    `json:"metadata"`
	Condition              string  `json:"condition"`
	Isin                   string  `json:"isin"`
	Currency               string  `json:"currency"`
	MarketPlace            string  `json:"marketPlace"`
}
