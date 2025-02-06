package api_models

// Auth()
type Auth_response struct {
	Tfa_response          Tfa_response `json:"twoFactorLogin"`
	AuthenticationSession string       `json:"authenticationSession"`
	PushSubscriptionId    string       `json:"pushSubscriptionId"`
	CustomerId            string       `json:"customerId"`
	CustomerSessionToken  string       `json:"customerSessionToken"`
	RegistrationComplete  bool         `json:"registrationComplete"`
}

type Tfa_response struct {
	Transaction_id string `json:"transactionId"`
	Method         string `json:"method"`
}

// Disconnect()
type SessionClose struct {
	Session string `json:"Session"`
}

// GetSessionInfo()
type SessionInfo struct {
	InvalidSessionID string `json:"invalidSessionId"`
	User             struct {
		LoggedIn           bool   `json:"loggedIn"`
		GreetingName       string `json:"greetingName"`
		PushSubscriptionID string `json:"pushSubscriptionId"`
		PushBaseURL        string `json:"pushBaseUrl"`
		SecurityToken      string `json:"securityToken"`
		Company            bool   `json:"company"`
		Minor              bool   `json:"minor"`
		Start              bool   `json:"start"`
		CustomerGroup      string `json:"customerGroup"`
		ID                 string `json:"id"`
	} `json:"user"`
}

// GetAccounts()
type Accounts []struct {
	Name            string `json:"name"`
	AccountID       string `json:"accountId"`
	AccountTypeName string `json:"accountTypeName"`
	AccountType     string `json:"accountType"`
	URLParameterID  string `json:"urlParameterId"`
	IsTradable      bool   `json:"isTradable"`
}

// GetPositions()
type PositionsRaw struct {
	WithOrderbook []struct {
		Account struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			Name           string `json:"name"`
			URLParameterID string `json:"urlParameterId"`
			HasCredit      bool   `json:"hasCredit"`
		} `json:"account"`
		Instrument struct {
			ID        string `json:"id"`
			Type      string `json:"type"`
			Name      string `json:"name"`
			Orderbook struct {
				ID          string `json:"id"`
				FlagCode    any    `json:"flagCode"`
				Name        string `json:"name"`
				Type        string `json:"type"`
				TradeStatus string `json:"tradeStatus"`
				Quote       struct {
					Highest struct {
						Value            float64 `json:"value"`
						Unit             string  `json:"unit"`
						UnitType         string  `json:"unitType"`
						DecimalPrecision int     `json:"decimalPrecision"`
					} `json:"highest"`
					Lowest struct {
						Value            float64 `json:"value"`
						Unit             string  `json:"unit"`
						UnitType         string  `json:"unitType"`
						DecimalPrecision int     `json:"decimalPrecision"`
					} `json:"lowest"`
					Buy    any `json:"buy"`
					Sell   any `json:"sell"`
					Latest struct {
						Value            float64 `json:"value"`
						Unit             string  `json:"unit"`
						UnitType         string  `json:"unitType"`
						DecimalPrecision int     `json:"decimalPrecision"`
					} `json:"latest"`
					Change struct {
						Value            float64 `json:"value"`
						Unit             string  `json:"unit"`
						UnitType         string  `json:"unitType"`
						DecimalPrecision int     `json:"decimalPrecision"`
					} `json:"change"`
					ChangePercent struct {
						Value            float64 `json:"value"`
						Unit             string  `json:"unit"`
						UnitType         string  `json:"unitType"`
						DecimalPrecision int     `json:"decimalPrecision"`
					} `json:"changePercent"`
					Updated string `json:"updated"`
				} `json:"quote"`
				Turnover struct {
					Volume struct {
						Value            int    `json:"value"`
						Unit             string `json:"unit"`
						UnitType         string `json:"unitType"`
						DecimalPrecision int    `json:"decimalPrecision"`
					} `json:"volume"`
					Value any `json:"value"`
				} `json:"turnover"`
				LastDeal struct {
					Date string `json:"date"`
					Time any    `json:"time"`
				} `json:"lastDeal"`
			} `json:"orderbook"`
			Currency     string  `json:"currency"`
			Isin         string  `json:"isin"`
			VolumeFactor float64 `json:"volumeFactor"`
		} `json:"instrument"`
		Volume struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"volume"`
		Value struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"value"`
		AverageAcquiredPrice struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"averageAcquiredPrice"`
		AverageAcquiredPriceInstrumentCurrency struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"averageAcquiredPriceInstrumentCurrency"`
		AcquiredValue struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"acquiredValue"`
		LastTradingDayPerformance struct {
			Absolute struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"absolute"`
			Relative struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"relative"`
		} `json:"lastTradingDayPerformance"`
		CollateralFactor struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"collateralFactor"`
		SuperInterestApproved bool   `json:"superInterestApproved"`
		ID                    string `json:"id"`
	} `json:"withOrderbook"`
	WithoutOrderbook []any `json:"withoutOrderbook"`
	CashPositions    []struct {
		Account struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			Name           string `json:"name"`
			URLParameterID string `json:"urlParameterId"`
			HasCredit      bool   `json:"hasCredit"`
		} `json:"account"`
		TotalBalance struct {
			Value            float64 `json:"value"`
			Unit             string  `json:"unit"`
			UnitType         string  `json:"unitType"`
			DecimalPrecision int     `json:"decimalPrecision"`
		} `json:"totalBalance"`
		ID string `json:"id"`
	} `json:"cashPositions"`
	WithCreditAccount bool `json:"withCreditAccount"`
}

// GetPriceInfo()
type PriceInfo struct {
	PriceOneWeekAgo     float64 `json:"priceOneWeekAgo,omitempty"`
	PriceOneMonthAgo    float64 `json:"priceOneMonthAgo,omitempty"`
	PriceSixMonthsAgo   float64 `json:"priceSixMonthsAgo,omitempty"`
	PriceThreeYearsAgo  float64 `json:"priceThreeYearsAgo,omitempty"`
	PriceFiveYearsAgo   float64 `json:"priceFiveYearsAgo,omitempty"`
	PriceThreeMonthsAgo float64 `json:"priceThreeMonthsAgo,omitempty"`
	PriceAtStartOfYear  float64 `json:"priceAtStartOfYear,omitempty"`
	PriceOneYearAgo     float64 `json:"priceOneYearAgo,omitempty"`
	NumberOfPriceAlerts int     `json:"numberOfPriceAlerts,omitempty"`
	PushPermitted       bool    `json:"pushPermitted,omitempty"`
	LowestPrice         float64 `json:"lowestPrice,omitempty"`
	Change              float64 `json:"change,omitempty"`
	HighestPrice        float64 `json:"highestPrice,omitempty"`
	Name                string  `json:"name,omitempty"`
	ID                  string  `json:"id,omitempty"`
	Currency            string  `json:"currency,omitempty"`
	ChangePercent       float64 `json:"changePercent,omitempty"`
	FlagCode            string  `json:"flagCode,omitempty"`
	LastPrice           float64 `json:"lastPrice,omitempty"`
	LastPriceUpdated    string  `json:"lastPriceUpdated,omitempty"`
	QuoteUpdated        string  `json:"quoteUpdated,omitempty"`
}

// GetHistoricalPrices()
type HistoricalPrices struct {
	Ohlc []struct {
		Timestamp         int64   `json:"timestamp,omitempty"`
		Open              float64 `json:"open,omitempty"`
		Close             float64 `json:"close,omitempty"`
		Low               float64 `json:"low,omitempty"`
		High              float64 `json:"high,omitempty"`
		TotalVolumeTraded int     `json:"totalVolumeTraded,omitempty"`
	} `json:"ohlc,omitempty"`
	Metadata struct {
		Resolution struct {
			ChartResolution      string   `json:"chartResolution,omitempty"`
			AvailableResolutions []string `json:"availableResolutions,omitempty"`
		} `json:"resolution,omitempty"`
	} `json:"metadata,omitempty"`
	From                 string  `json:"from,omitempty"`
	To                   string  `json:"to,omitempty"`
	PreviousClosingPrice float64 `json:"previousClosingPrice,omitempty"`
}

// GetWarrentInfo()
type WarrantInfo struct {
	OrderbookID string `json:"orderbookId"`
	Name        string `json:"name"`
	Isin        string `json:"isin"`
	Tradable    string `json:"tradable"`
	Listing     struct {
		ShortName             string `json:"shortName"`
		TickerSymbol          string `json:"tickerSymbol"`
		CountryCode           string `json:"countryCode"`
		Currency              string `json:"currency"`
		MarketPlaceCode       string `json:"marketPlaceCode"`
		MarketPlaceName       string `json:"marketPlaceName"`
		TickSizeListID        string `json:"tickSizeListId"`
		MarketTradesAvailable bool   `json:"marketTradesAvailable"`
	} `json:"listing"`
	HistoricalClosingPrices struct {
		OneDay      float64 `json:"oneDay"`
		OneWeek     float64 `json:"oneWeek"`
		OneMonth    float64 `json:"oneMonth"`
		StartOfYear float64 `json:"startOfYear"`
		Start       float64 `json:"start"`
		StartDate   string  `json:"startDate"`
	} `json:"historicalClosingPrices"`
	KeyIndicators struct {
		Parity         int     `json:"parity"`
		BarrierLevel   float64 `json:"barrierLevel"`
		FinancingLevel float64 `json:"financingLevel"`
		Direction      string  `json:"direction"`
		StrikePrice    float64 `json:"strikePrice"`
		Leverage       float64 `json:"leverage"`
		NumberOfOwners int     `json:"numberOfOwners"`
		SubType        string  `json:"subType"`
		IsAza          bool    `json:"isAza"`
	} `json:"keyIndicators"`
	Quote struct {
		Buy               float64 `json:"buy"`
		Sell              float64 `json:"sell"`
		Last              float64 `json:"last"`
		Change            float64 `json:"change"`
		ChangePercent     float64 `json:"changePercent"`
		Spread            float64 `json:"spread"`
		TimeOfLast        int64   `json:"timeOfLast"`
		TotalValueTraded  float64 `json:"totalValueTraded"`
		TotalVolumeTraded int     `json:"totalVolumeTraded"`
		Updated           int64   `json:"updated"`
		IsRealTime        bool    `json:"isRealTime"`
	} `json:"quote"`
	Type       string `json:"type"`
	Underlying struct {
		OrderbookID       string `json:"orderbookId"`
		Name              string `json:"name"`
		InstrumentType    string `json:"instrumentType"`
		InstrumentSubType string `json:"instrumentSubType"`
		Quote             struct {
			Buy               float64 `json:"buy"`
			Sell              float64 `json:"sell"`
			Last              float64 `json:"last"`
			Highest           float64 `json:"highest"`
			Lowest            float64 `json:"lowest"`
			Change            float64 `json:"change"`
			ChangePercent     float64 `json:"changePercent"`
			Spread            float64 `json:"spread"`
			TimeOfLast        int64   `json:"timeOfLast"`
			TotalValueTraded  int     `json:"totalValueTraded"`
			TotalVolumeTraded int     `json:"totalVolumeTraded"`
			Updated           int64   `json:"updated"`
			IsRealTime        bool    `json:"isRealTime"`
		} `json:"quote"`
		Listing struct {
			ShortName             string `json:"shortName"`
			TickerSymbol          string `json:"tickerSymbol"`
			CountryCode           string `json:"countryCode"`
			Currency              string `json:"currency"`
			MarketPlaceCode       string `json:"marketPlaceCode"`
			MarketPlaceName       string `json:"marketPlaceName"`
			TickSizeListID        string `json:"tickSizeListId"`
			MarketTradesAvailable bool   `json:"marketTradesAvailable"`
		} `json:"listing"`
		PreviousClosingPrice float64 `json:"previousClosingPrice"`
	} `json:"underlying"`
}

// GetWarrantList()
type WarrantList struct {
	Warrants []struct {
		OrderbookID          string `json:"orderbookId,omitempty"`
		CountryCode          string `json:"countryCode,omitempty"`
		Name                 string `json:"name,omitempty"`
		Direction            string `json:"direction,omitempty"`
		Issuer               string `json:"issuer,omitempty"`
		SubType              string `json:"subType,omitempty"`
		HasPosition          bool   `json:"hasPosition,omitempty"`
		UnderlyingInstrument struct {
			Name           string `json:"name,omitempty"`
			OrderbookID    string `json:"orderbookId,omitempty"`
			InstrumentType string `json:"instrumentType,omitempty"`
			CountryCode    string `json:"countryCode,omitempty"`
		} `json:"underlyingInstrument,omitempty"`
		TotalValueTraded    int     `json:"totalValueTraded,omitempty"`
		StopLoss            float64 `json:"stopLoss,omitempty"`
		FinancingLevel      float64 `json:"financingLevel,omitempty"`
		Spread              float64 `json:"spread,omitempty"`
		Leverage            float64 `json:"leverage,omitempty"`
		OneDayChangePercent float64 `json:"oneDayChangePercent,omitempty"`
		BuyPrice            float64 `json:"buyPrice,omitempty"`
		SellPrice           float64 `json:"sellPrice,omitempty"`
	} `json:"warrants,omitempty"`
	Filter struct {
		SubTypes              []any `json:"subTypes,omitempty"`
		Issuers               []any `json:"issuers,omitempty"`
		UnderlyingInstruments []any `json:"underlyingInstruments,omitempty"`
		EndDates              []any `json:"endDates,omitempty"`
		Directions            []any `json:"directions,omitempty"`
	} `json:"filter,omitempty"`
	Pagination struct {
		Offset int `json:"offset,omitempty"`
		Limit  int `json:"limit,omitempty"`
	} `json:"pagination,omitempty"`
	SortBy struct {
		Field string `json:"field,omitempty"`
		Order string `json:"order,omitempty"`
	} `json:"sortBy,omitempty"`
	TotalNumberOfOrderbooks int `json:"totalNumberOfOrderbooks,omitempty"`
	FilterOptions           struct {
		Issuers []struct {
			Value              string `json:"value,omitempty"`
			DisplayName        string `json:"displayName,omitempty"`
			NumberOfOrderbooks int    `json:"numberOfOrderbooks,omitempty"`
		} `json:"issuers,omitempty"`
		UnderlyingInstruments []struct {
			Value              string `json:"value,omitempty"`
			DisplayName        string `json:"displayName,omitempty"`
			NumberOfOrderbooks int    `json:"numberOfOrderbooks,omitempty"`
		} `json:"underlyingInstruments,omitempty"`
		EndDates []struct {
			Value              string `json:"value,omitempty"`
			DisplayName        string `json:"displayName,omitempty"`
			NumberOfOrderbooks int    `json:"numberOfOrderbooks,omitempty"`
		} `json:"endDates,omitempty"`
		SubTypes []struct {
			Value              string `json:"value,omitempty"`
			DisplayName        string `json:"displayName,omitempty"`
			NumberOfOrderbooks int    `json:"numberOfOrderbooks,omitempty"`
		} `json:"subTypes,omitempty"`
		Directions []struct {
			Value              string `json:"value,omitempty"`
			DisplayName        string `json:"displayName,omitempty"`
			NumberOfOrderbooks int    `json:"numberOfOrderbooks,omitempty"`
		} `json:"directions,omitempty"`
	} `json:"filterOptions,omitempty"`
}

// OrderResponse
type OrderResponse struct {
	OrderRequestStatus string   `json:"orderRequestStatus,omitempty"`
	Message            string   `json:"message,omitempty"`
	Parameters         []string `json:"parameters,omitempty"`
	OrderID            string   `json:"orderId,omitempty"`
}

// GetMatchingPrice
type MatchingPrice struct {
	Price float64 `json:"price,omitempty"`
}

// Handshake response
type HandshakeResponse []struct {
	MinimumVersion           string   `json:"minimumVersion,omitempty"`
	ClientID                 string   `json:"clientId,omitempty"`
	SupportedConnectionTypes []string `json:"supportedConnectionTypes,omitempty"`
	Advice                   struct {
		Interval  int    `json:"interval,omitempty"`
		Timeout   int    `json:"timeout,omitempty"`
		Reconnect string `json:"reconnect,omitempty"`
	} `json:"advice,omitempty"`
	Channel    string `json:"channel,omitempty"`
	Version    string `json:"version,omitempty"`
	Successful bool   `json:"successful,omitempty"`
}

// CheckOrder()
type OrderStatus struct {
	OrderID         string  `json:"orderId,omitempty"`
	OrderbookID     string  `json:"orderbookId,omitempty"`
	Side            string  `json:"side,omitempty"`
	State           string  `json:"state,omitempty"`
	MarketReference string  `json:"marketReference,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Message         string  `json:"message,omitempty"`
	Volume          int     `json:"volume,omitempty"`
	OriginalVolume  int     `json:"originalVolume,omitempty"`
	AccountID       string  `json:"accountId,omitempty"`
	Condition       string  `json:"condition,omitempty"`
	ValidUntil      string  `json:"validUntil,omitempty"`
	Modifiable      bool    `json:"modifiable,omitempty"`
	Deletable       bool    `json:"deletable,omitempty"`
}

type ValidationResponse struct {
	CommissionWarning struct {
		Valid bool `json:"valid"`
	} `json:"commissionWarning"`
	EmployeeValidation struct {
		Valid bool `json:"valid"`
	} `json:"employeeValidation"`
	LargeInScaleWarning struct {
		Valid bool `json:"valid"`
	} `json:"largeInScaleWarning"`
	OrderValueLimitWarning struct {
		Valid bool `json:"valid"`
	} `json:"orderValueLimitWarning"`
	PriceRampingWarning struct {
		Valid bool `json:"valid"`
	} `json:"priceRampingWarning"`
}
