package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Melle101/mom-bot-v3/internal/config"
	"github.com/Melle101/mom-bot-v3/internal/trading"
)

func main() {

	log.Println("Starting service.")

	for {
		nextRun, err := getNextTradeDay()
		if err != nil {
			log.Fatalf("Error getting next trade date: %v", err)
		}

		log.Println("Next run scheduled for: ", nextRun.Format("2006-01-02 15:04"))

		time.Sleep(time.Until(*nextRun))

		log.Println("Executing sequence at ", time.Now().Format("2006-01-02 15:04"))
		excuteHoldingsSwap()
		if err != nil {
			log.Fatalf("Errors: %v", err)
		}
	}
}

func excuteHoldingsSwap() error {
	config.InitLogger()

	client, err := trading.NewClient()
	if err != nil {
		return fmt.Errorf("Error starting client: %v", err)
	}

	allPositions, err := client.ApiClient.GetPositions()
	if err != nil {
		return fmt.Errorf("Failed to get all positions. Exiting. Error: %v", err)
	}

	currentPositions, err := getAccountPositions(client.Cfg.Settings.AccountURL, allPositions)
	if err != nil {
		return fmt.Errorf("Failed to get account positions. Exiting. Error: %v", err)
	}

	upcomingHoldings, err := findUpcomingAssets(client.Cfg)
	if err != nil {
		return fmt.Errorf("Failed to get upcoming holdings. Exiting. Error: %v", err)
	}

	trades, err := getTrades(currentPositions, upcomingHoldings, client.Cfg)
	if err != nil {
		return fmt.Errorf("Failed to trades. Exiting. Error: %v", err)
	}

	err = executeTrades(&client.ApiClient, trades)
	if err != nil {
		return fmt.Errorf("Failed to execute trades. Exiting. Error: %v", err)
	}

	logPeriod(currentPositions, upcomingHoldings)

	return nil
}
