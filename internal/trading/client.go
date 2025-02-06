package trading

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	avanza_api "github.com/Melle101/mom-bot-v3/avanza-api"
	"github.com/Melle101/mom-bot-v3/internal/config"
	"github.com/Melle101/mom-bot-v3/internal/models"
)

func NewClient() (*models.Client, error) {
	var auth struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		TOTP     string `toml:"TOTP"`
	}

	log.Println("Loading auth.toml")
	if _, err := toml.DecodeFile("auth.toml", &auth); err != nil {
		return nil, fmt.Errorf("failed to load auth config: %w", err)
	}

	log.Println("Loading config.toml")
	config, err := config.Load("config.toml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println("Creating ApiClient")
	cl, err := avanza_api.CreateClient(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	client := &models.Client{
		ApiClient: *cl,
		Cfg:       *config,
	}

	return client, nil
}
