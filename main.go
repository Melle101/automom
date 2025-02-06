package main

import (
	"log"

	"github.com/Melle101/mom-bot-v3/internal/config"
	"github.com/Melle101/mom-bot-v3/internal/trading"
)

func main() {

	config.InitLogger()

	client, err := trading.NewClient()
	if err != nil {
		log.Fatalf("Error starting client: %v", err)
	}

	//sessionInfo, err := client.ApiClient.GetSessionInfo()
	//accounts, err := client.ApiClient.GetAccounts()

	allPositions, err := client.ApiClient.GetPositions()
	if err != nil {
		log.Fatalf("Failed to get all positions. Exiting. Error: %v", err)
	}

	currentPositions, err := getAccountPositions(client.Cfg.Settings.AccountURL, allPositions)
	if err != nil {
		log.Fatalf("Failed to get account positions. Exiting. Error: %v", err)
	}

	upcomingHoldings, err := findUpcomingAssets(client.Cfg)
	if err != nil {
		log.Fatalf("Failed to get upcoming holdings. Exiting. Error: %v", err)
	}

	trades, err := getTrades(currentPositions, upcomingHoldings, client.Cfg)
	if err != nil {
		log.Fatalf("Failed to trades. Exiting. Error: %v", err)
	}

	err = executeTrades(&client.ApiClient, trades)
	if err != nil {
		log.Fatalf("Failed to execute trades. Exiting. Error: %v", err)
	}

}
