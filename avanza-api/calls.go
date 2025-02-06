package avanza_api

import (
	"fmt"
	"strconv"
	"strings"

	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
)

func (c ApiClient) GetSessionInfo() (*api_models.SessionInfo, error) {
	url := api_models.BASE_URL + api_models.SESSION_INFO_URL

	resp, err := HTTPGet[api_models.SessionInfo](url, c.CredHeaders)
	if err != nil {
		return resp, fmt.Errorf("failed to get session info: %w", err)
	}

	return resp, err
}

func (c *ApiClient) GetAccounts() (api_models.Accounts, error) {
	url := api_models.BASE_URL + api_models.ACCOUNTS_URL

	resp, err := HTTPGet[api_models.Accounts](url, c.CredHeaders)
	if err != nil {
		return *resp, fmt.Errorf("failed to get accounts: %w", err)
	}

	return *resp, err
}

func (c *ApiClient) GetPositions() (api_models.PositionsRaw, error) {
	url := api_models.BASE_URL + api_models.POSITIONS_URL

	resp, err := HTTPGet[api_models.PositionsRaw](url, c.CredHeaders)
	if err != nil {
		return *resp, fmt.Errorf("failed to get positions: %w", err)
	}

	return *resp, err
}

func GetPriceInfo(assetID string) (api_models.PriceInfo, error) {
	url := api_models.BASE_URL + api_models.PRICE_INFO_URL + assetID

	resp, err := HTTPGet[api_models.PriceInfo](url, nil)
	if err != nil {
		return *resp, fmt.Errorf("failed to get price info: %w", err)
	}

	return *resp, err
}

func GetHistoricalPrices(assetID, length string) (api_models.HistoricalPrices, error) {
	validLengths := map[string]bool{"ONE_WEEK": true, "ONE_MONTH": true, "THREE_MONTHS": true, "SIX_MONTHS": true, "ONE_YEAR": true}

	if !validLengths[length] {
		return api_models.HistoricalPrices{}, fmt.Errorf("failed to get historical prices: invalid length")
	}

	url := api_models.BASE_URL + api_models.HISTORICAL_PRICES_URL + assetID + "?timePeriod=" + strings.ToLower(length)

	resp, err := HTTPGet[api_models.HistoricalPrices](url, nil)
	if err != nil {
		return *resp, fmt.Errorf("failed to get historical prices: %w", err)
	}

	return *resp, err
}

func GetWarrantInfo(assetID string) (*api_models.WarrantInfo, error) {
	url := api_models.BASE_URL + api_models.WARRANT_INFO_URL + assetID

	resp, err := HTTPGet[api_models.WarrantInfo](url, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to get warrant info: %w", err)
	}

	return resp, err
}

func GetWarrantList(searchInfo api_models.WarrantSearch) (api_models.WarrantList, error) {
	url := api_models.BASE_URL + api_models.WARRANT_LIST_URL

	resp, err := HTTPPost[api_models.WarrantList](url, searchInfo, nil)
	if err != nil {
		return *resp, fmt.Errorf("failed to get warrant list: %w", err)
	}

	return *resp, err
}

func (c *ApiClient) PlaceOrder(orderInfo api_models.NewOrder) (api_models.OrderResponse, error) {
	url := api_models.BASE_URL + api_models.PLACE_ORDER_URL
	fmt.Println(orderInfo)

	resp, err := HTTPPost[api_models.OrderResponse, api_models.NewOrder](url, orderInfo, c.CredHeaders)
	if err != nil {
		return api_models.OrderResponse{}, fmt.Errorf("failed to place warrant order: %w", err)
	}

	return *resp, err
}

func (c *ApiClient) GetMatchingPrice(assetID, side string, volume float64) (float64, error) {
	volString := strconv.Itoa(int(volume))
	url := api_models.BASE_URL + api_models.MATCHING_PRICE_URL + assetID + "/" + side + "?volume=" + volString

	resp, err := HTTPGet[api_models.MatchingPrice](url, c.CredHeaders)
	if err != nil {
		return 0.00, fmt.Errorf("failed to get matching price for asset %s: %w", assetID, err)
	}

	return resp.Price, nil
}

func (c *ApiClient) CheckOrder(accID, orderID string) (api_models.OrderStatus, error) {
	url := api_models.BASE_URL + api_models.CHECK_ORDER_URL + "?orderId=" + orderID + "&cAccountId=" + accID

	resp, err := HTTPGet[api_models.OrderStatus](url, c.CredHeaders)
	if err != nil {
		return *resp, fmt.Errorf("error getting order status for order %s: %w", orderID, err)
	}

	return *resp, nil
}

func (c *ApiClient) ModifyOrder(accID, orderID string, price float64, volume int) (api_models.OrderResponse, error) {
	url := api_models.BASE_URL + api_models.MODIFY_ORDER_URL

	payload := api_models.ModifyOrderInfo{
		AccountID: accID,
		OrderID:   orderID,
		Price:     price,
		Volume:    volume,
	}

	resp, err := HTTPPost[api_models.OrderResponse, api_models.ModifyOrderInfo](url, payload, nil)
	if err != nil {
		return *resp, fmt.Errorf("failed to modify order %s: %w", orderID, err)
	}

	return *resp, nil
}

func (c *ApiClient) GetRequestID() (*string, error) {
	url := api_models.BASE_URL + api_models.REQ_ID_URL

	resp, err := HTTPGet[string](url, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to get requestID: %w", err)
	}

	return resp, nil
}

func (c *ApiClient) ValidateOrder(requestInfo api_models.ValidationRequest) (api_models.ValidationResponse, error) {
	url := api_models.BASE_URL + api_models.VALIDATE_URL

	resp, err := HTTPPost[api_models.ValidationResponse, api_models.ValidationRequest](url, requestInfo, c.CredHeaders)
	if err != nil {
		return *resp, fmt.Errorf("failed to vaildate order: %w", err)
	}

	return *resp, nil
}
