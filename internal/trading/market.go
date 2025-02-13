package trading

import (
	"fmt"
	"time"

	avanza_api "github.com/Melle101/mom-bot-v3/avanza-api"
	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
	"github.com/Melle101/mom-bot-v3/internal/models"
)

func GetPercentageChange(assetID, lookbackPeriod string) (float64, error) {
	priceInfo, err := avanza_api.GetPriceInfo(assetID)
	if err != nil {
		return 0.0, err
	}

	var comparePrice float64

	switch lookbackPeriod {
	case "ONE_WEEK":
		comparePrice = priceInfo.PriceOneWeekAgo
	case "ONE_MONTH":
		comparePrice = priceInfo.PriceOneMonthAgo
	case "THREE_MONTHS":
		comparePrice = priceInfo.PriceThreeMonthsAgo
	case "SIX_MONTHS":
		comparePrice = priceInfo.PriceSixMonthsAgo
	case "ONE_YEAR":
		comparePrice = priceInfo.PriceOneYearAgo
	}

	if comparePrice == 0 {
		return 0.0, fmt.Errorf("compare price is zero")
	}

	return priceInfo.LastPrice / comparePrice, nil
}

func GetRelativeSMA(assetID string, length int) (float64, error) {
	prices, err := avanza_api.GetHistoricalPrices(assetID, "ONE_YEAR")
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for i := range length {
		index := len(prices.Ohlc) - i - 1

		sum += prices.Ohlc[index].Close
	}

	lastPrice, err := GetLastPrice(assetID)
	if err != nil {
		return 0, err
	}

	return lastPrice / (sum / float64(length)), nil
}

func GetLastPrice(assetID string) (float64, error) {
	priceInfo, err := avanza_api.GetPriceInfo(assetID)
	if err != nil {
		return 0, err
	}

	return priceInfo.LastPrice, nil
}

func WarrantMarketOrder(ApiClient *avanza_api.ApiClient, orderInfo models.OrderInfo) (api_models.OrderStatus, error) {

	matchingPrice, err := ApiClient.GetMatchingPrice(orderInfo.AssetID, orderInfo.OrderType, orderInfo.Quantity)
	if err != nil {
		return api_models.OrderStatus{}, fmt.Errorf("failed to get matching price: %w", err)
	}

	reqID, err := ApiClient.GetRequestID()
	if err != nil {
		return api_models.OrderStatus{}, fmt.Errorf("failed to get requestID: %w", err)
	}

	newOrderInfo := api_models.NewOrder{
		AccountId:              orderInfo.AccountID,
		Condition:              "NORMAL",
		IsDividendReinvestment: false,
		OrderbookId:            orderInfo.AssetID,
		Side:                   orderInfo.OrderType,
		Volume:                 int(orderInfo.Quantity),
		Price:                  matchingPrice,
		RequestId:              *reqID,
		Metadata: api_models.Metadata{
			HasTouchedPrice: "true",
			OrderEntryMode:  "STANDARD",
		},
	}

	orderResponse, err := ApiClient.PlaceOrder(newOrderInfo)

	if err != nil {
		return api_models.OrderStatus{}, fmt.Errorf("failed to place order: %w", err)
	} else if orderResponse.OrderRequestStatus != "SUCCESS" {
		return api_models.OrderStatus{}, fmt.Errorf("failed to place order: %v", orderResponse)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(15 * time.Second) // Wait 10 seconds between checks

		orderStatus, err := ApiClient.CheckOrder(orderInfo.AccountID, orderResponse.OrderID)
		if err != nil {
			return api_models.OrderStatus{}, err
		}

		if orderStatus.State == "FILLED" {
			return orderStatus, nil
		}

		// if not filled, modify order
		matchingPrice, err := ApiClient.GetMatchingPrice(orderInfo.AssetID, orderInfo.OrderType, orderInfo.Quantity)
		if err != nil {
			return api_models.OrderStatus{}, fmt.Errorf("failed to get matching price: %w", err)
		}

		resp, err := ApiClient.ModifyOrder(orderInfo.AccountID, orderResponse.OrderID, matchingPrice, orderStatus.Volume)
		if err != nil || resp.OrderRequestStatus != "SUCCESS" {
			return api_models.OrderStatus{}, fmt.Errorf("failed to modify order %s: %w", orderStatus.OrderID, err)
		}
	}
	return api_models.OrderStatus{}, fmt.Errorf("couldn't execute market order for order %s", orderInfo.AssetID)
}
