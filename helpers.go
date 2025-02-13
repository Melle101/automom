package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"

	avanza_api "github.com/Melle101/mom-bot-v3/avanza-api"
	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
	"github.com/Melle101/mom-bot-v3/internal/models"
	"github.com/Melle101/mom-bot-v3/internal/trading"
)

func getAccountPositions(accURL string, allPositions api_models.PositionsRaw) (models.AccountPositions, error) {
	log.Printf("Getting positions for account %s", accURL)
	accountPositions := models.AccountPositions{
		AccountURL: accURL,
	}

	for _, cashPos := range allPositions.CashPositions {
		if cashPos.Account.URLParameterID == accURL {
			accountPositions.AccountID = cashPos.Account.ID

			accountPositions.CashPosition = models.CashPosition{
				Value:    cashPos.TotalBalance.Value,
				Currency: cashPos.TotalBalance.Unit,
			}

		}
	}

	for _, pos := range allPositions.WithOrderbook {
		if pos.Account.URLParameterID == accURL {
			accountPositions.AccountID = pos.Account.ID
			var underlying string

			if pos.Instrument.Type == "WARRANT" {
				warrantInfo, err := avanza_api.GetWarrantInfo(pos.Instrument.Orderbook.ID)
				if err != nil {
					return models.AccountPositions{}, fmt.Errorf("failed to get warrant info for %s: %v", pos.Instrument.Orderbook.ID, err)
				}

				underlying = warrantInfo.Underlying.OrderbookID
			}

			accountPositions.Positions = append(accountPositions.Positions, models.Position{
				ID:                      pos.ID,
				AssetType:               pos.Instrument.Type,
				AssetName:               pos.Instrument.Name,
				OrderbookID:             pos.Instrument.Orderbook.ID,
				Currency:                pos.Instrument.Currency,
				Quantity:                pos.Volume.Value,
				CurrentValue:            pos.Value.Value,
				AcquisitionValue:        pos.AcquiredValue.Value,
				AverageAcquisitionPrice: pos.AverageAcquiredPrice.Value,
				UnderlyingID:            underlying,
			})

		}
	}
	log.Println("Account positions fetched:")
	log.Println(accountPositions)
	return accountPositions, nil
}

func findUpcomingAssets(cfg models.Config) ([]models.Asset, error) {
	log.Println("Getting upcoming holdings")
	type assetInfo struct {
		underlyingName   string
		underlyingID     string
		assetType        string
		percentageChange float64
		relativeSMA      float64
		targetLev        int
	}

	assetInfos := make([]assetInfo, len(cfg.Assets))

	for i, asset := range cfg.Assets {
		percentageChange, err := trading.GetPercentageChange(asset.UnderlyingID, cfg.Settings.LookbackPeriod)
		relativeSMA, err := trading.GetRelativeSMA(asset.UnderlyingID, cfg.Settings.SMAFilterLength)
		if err != nil {
			return nil, fmt.Errorf("failed to get asset info (%s) for upcoming assets: %w", asset.UnderlyingID, err)
		}

		assetInfos[i] = assetInfo{
			underlyingName:   asset.UnderlyingName,
			underlyingID:     asset.UnderlyingID,
			assetType:        asset.AssetType,
			percentageChange: percentageChange,
			relativeSMA:      relativeSMA,
			targetLev:        asset.TargetLev,
		}
	}

	//Sort by percentage change
	sort.Slice(assetInfos, func(i, j int) bool {
		return assetInfos[i].percentageChange > assetInfos[j].percentageChange
	})

	upcomingHoldings := make([]models.Asset, cfg.Settings.AGG)

	for i := range cfg.Settings.AGG {
		if assetInfos[i].relativeSMA > 1 {
			upcomingHoldings[i] = models.Asset{
				UnderlyingName: assetInfos[i].underlyingName,
				UnderlyingID:   assetInfos[i].underlyingID,
				AssetType:      assetInfos[i].assetType,
				TargetLev:      assetInfos[i].targetLev,
			}
		} else {
			upcomingHoldings[i] = models.Asset{
				UnderlyingName: cfg.Settings.BackupAssetID,
				UnderlyingID:   cfg.Settings.BackupAssetID,
				AssetType:      cfg.Settings.BackupType,
				TargetLev:      1,
			}
		}
	}

	log.Println("Upcoming holdings fetched:")
	log.Println(upcomingHoldings)
	return upcomingHoldings, nil
}

func getTrades(accountPositions models.AccountPositions, upcomingHoldings []models.Asset, cfg models.Config) (models.TradesInfo, error) {
	log.Println("Getting trades")
	var sells []models.OrderInfo
	var buys []models.OrderInfo

	fmt.Println(accountPositions)
	fmt.Println(upcomingHoldings)

	for _, currentPos := range accountPositions.Positions {
		inUpcomingHoldings := false
		for _, upcomingHolding := range upcomingHoldings {
			if currentPos.UnderlyingID == upcomingHolding.UnderlyingID {
				inUpcomingHoldings = true
				break
			}
		}

		if !inUpcomingHoldings {
			sells = append(sells, models.OrderInfo{
				UnderlyingName: currentPos.AssetName,
				UnderlyingID:   currentPos.UnderlyingID,
				AssetID:        currentPos.OrderbookID,
				OrderType:      "SELL",
				Quantity:       currentPos.Quantity,
				AssetType:      currentPos.AssetType,
				AccountID:      cfg.Settings.AccountID,
			})
		}
	}

	BackupPositions := 0

	for _, upcomingHolding := range upcomingHoldings {
		if upcomingHolding.UnderlyingID == cfg.Settings.BackupAssetID {
			BackupPositions++
		}

		inCurrentPositions := false
		for _, currentPos := range accountPositions.Positions {
			if currentPos.UnderlyingID == upcomingHolding.UnderlyingID {
				inCurrentPositions = true
				break
			}
		}

		if !inCurrentPositions {

			var suitableAsset string
			var err error
			if upcomingHolding.AssetType == "WARRANT" {
				suitableAsset, err = findSuitableAsset(upcomingHolding.UnderlyingID, upcomingHolding.TargetLev)
				if err != nil {
					return models.TradesInfo{}, fmt.Errorf("failed to find suitable asset for %s: %w", upcomingHolding.UnderlyingID, err)
				}
			} else if upcomingHolding.AssetType == "CASH" {
				continue
			} else {
				suitableAsset = upcomingHolding.UnderlyingID
			}

			buys = append(buys, models.OrderInfo{
				UnderlyingName: upcomingHolding.UnderlyingName,
				UnderlyingID:   upcomingHolding.UnderlyingID,
				AssetID:        suitableAsset,
				OrderType:      "BUY",
				AssetType:      upcomingHolding.AssetType,
				AccountID:      cfg.Settings.AccountID,
			})
		}
	}

	trades := models.TradesInfo{
		Sells:           sells,
		Buys:            buys,
		BackupPositions: BackupPositions,
	}

	log.Println("Trades fetched:")
	log.Println(trades)
	return trades, nil
}

func findSuitableAsset(underlying string, targetLev int) (string, error) {
	log.Printf("Finding suitable asset for underlyingID: %s", underlying)
	searchInfo := api_models.WarrantSearch{
		Filter: api_models.Filter{
			Directions: []string{"long"},
			Issuers:    []string{},
			SubTypes:   []string{"mini_future"},
			EndDates:   []string{},
		},
		Limit:  20,
		Offset: 0,
		SortBy: api_models.SortBy{
			Field: "leverage",
			Order: "asc",
		},
	}

	if underlying == "155458" {
		log.Printf("UnderlyingID: %s requires search parameter, adding it", underlying)
		searchInfo.Filter.NameQuery = "SP500"
	} else {
		searchInfo.Filter.UnderlyingInstruments = []string{underlying}
	}

	fmt.Println(searchInfo.Filter.UnderlyingInstruments)
	warrantList, err := avanza_api.GetWarrantList(searchInfo)
	if err != nil {
		return "", fmt.Errorf("failed to get warrant list: %w", err)
	}

	sort.Slice(warrantList.Warrants, func(i, j int) bool {
		// Calculate the absolute difference from the target
		distI := math.Abs(float64(warrantList.Warrants[i].Leverage - float64(targetLev)))
		distJ := math.Abs(float64(warrantList.Warrants[j].Leverage - float64(targetLev)))
		return distI < distJ
	})

	for _, warrant := range warrantList.Warrants {
		if warrant.TotalValueTraded > 0 && warrant.Leverage < float64(targetLev)+1 && warrant.BuyPrice <= 3000 {
			return warrant.OrderbookID, nil
		}
	}

	if len(warrantList.Warrants) == 0 {
		return "", fmt.Errorf("failed to find suitable asset for underlying %s.", underlying)

	}
	return warrantList.Warrants[0].OrderbookID, nil
}

func executeTrades(client *avanza_api.ApiClient, trades models.TradesInfo) error {

	log.Println("Executing trades")
	for _, trade := range trades.Sells {
		if trade.AssetType == "WARRANT" {

			log.Printf("Executing sell of assetID: %s (underlyingID: %s)", trade.AssetID, trade.UnderlyingID)
			_, err := trading.WarrantMarketOrder(client, trade)
			if err != nil {
				return fmt.Errorf("Failed to execute sell of asset %s. Error: %w", trade.AssetID, err)
			}
		}
		log.Println("Pausing for a minute to avoid rate limit")
		time.Sleep(60 * time.Second)
	}

	if len(trades.Buys) == 0 {
		return nil
	}

	positionInfo, err := client.GetPositions()
	if err != nil {
		return fmt.Errorf("Error getting cash position: %w", err)
	}

	var availableCash float64
	for _, cashPos := range positionInfo.CashPositions {
		if cashPos.Account.ID == trades.Buys[0].AccountID {
			availableCash = cashPos.TotalBalance.Value
		}
	}

	log.Printf("Calculating cash per buy. Available cash: %f", availableCash)
	cashPerBuy := availableCash / float64(len(trades.Buys)+trades.BackupPositions)
	log.Printf("Calculated cash per buy; %f", cashPerBuy)

	for _, trade := range trades.Buys {

		matchingPrice, err := client.GetMatchingPrice(trade.AssetID, trade.OrderType, 1)
		if err != nil {
			return fmt.Errorf("Error getting matching price for buy of asset: %s", trade.AssetID)
		}

		trade.Quantity = math.Floor(cashPerBuy / matchingPrice)
		if trade.Quantity == 0 {
			return fmt.Errorf("Not enough cash to buy atleast one share of %s", trade.AssetID)
		}

		if trade.AssetType == "WARRANT" {

			log.Printf("Executing buy of assetID: %s (underlyingID: %s)", trade.AssetID, trade.UnderlyingID)
			_, err := trading.WarrantMarketOrder(client, trade)
			if err != nil {
				return fmt.Errorf("Failed to execute buy of asset %s. Error: %w", trade.AssetID, err)
			}
		}

		log.Println("Pausing for a minute to avoid rate limit")
		time.Sleep(60 * time.Second)
	}
	log.Println("Trades executed sucessfully")
	return nil
}

func logPeriod(past models.AccountPositions, upcoming []models.Asset) {
	logFile, err := os.OpenFile("trades.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	fmt.Fprintf(logFile, "\n\n########### %v ##########", time.Now().Format("2006-01-02"))
	fmt.Fprintln(logFile, "\nPast period's holdings:")
	fmt.Fprintf(logFile, "\n%-30s\t%-20s\t%-10s\t%5s\n", "Asset Name", "ID", "Type", "Change [%]")

	for _, holding := range past.Positions {
		change := ((holding.CurrentValue / holding.AcquisitionValue) - 1) * 100
		fmt.Fprintf(logFile, "%-30s\t%-20s\t%-10s\t%5.2f\n", holding.AssetName, holding.ID, holding.AssetType, change)
	}

	fmt.Fprintln(logFile, "\nUpcoming holdings:")
	fmt.Fprintf(logFile, "\n%-20s\t%-20s\n", "Underlying", "Type")

	for _, next := range upcoming {
		fmt.Fprintf(logFile, "%-20s\t%-20s\n", next.UnderlyingName, next.AssetType)
	}
}

func getNextTradeDay() (*time.Time, error) {
	today := time.Now()
	weekday := today.Weekday() // Get today's weekday

	// Calculate days until next Monday
	daysUntilMonday := (8 - int(weekday)) % 7
	if daysUntilMonday == 0 {
		daysUntilMonday = 7 // If today is Monday, get next week's Monday
	}

	nextMonday := today.AddDate(0, 0, daysUntilMonday)
	nextMonday = time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 11, 0, 0, 0, nextMonday.Location())
	for true {
		tradeDayFound, err := isTradeDay(nextMonday.Format("2006-01-02"), "SE")
		if err != nil {
			return nil, err
		}

		if tradeDayFound {
			return &nextMonday, nil
		} else {
			nextMonday.Add(time.Hour * 24)
		}
	}

	return nil, errors.New("failed to find trade day")
}

func isTradeDay(dateString, market string) (bool, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return false, fmt.Errorf("error parsing date: %w", err)
	}

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false, nil
	}

	irrDates, err := avanza_api.GetIrregularDates()
	if err != nil {
		return false, err
	}

	for _, date := range irrDates {
		if date.Date == dateString && date.CountryCode == market {
			return false, nil
		}
	}

	return true, nil
}
