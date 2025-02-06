package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Melle101/mom-bot-v3/internal/models"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Msg)
}

func Load(path string) (*models.Config, error) {
	var cfg models.Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validate(cfg *models.Config) error {
	if cfg.Settings.AGG <= 0 {
		return &ValidationError{
			Field: "AGG",
			Msg:   "must be greater than 0",
		}
	}

	if cfg.Settings.SMAFilterLength <= 0 {
		return &ValidationError{
			Field: "SMAFilterLength",
			Msg:   "must be greater than 0",
		}
	}

	if cfg.Settings.AccountURL == "" {
		return &ValidationError{
			Field: "AccountURL",
			Msg:   "cannot be empty",
		}
	}

	validPeriods := map[string]bool{"ONE_WEEK": true, "ONE_MONTH": true, "THREE_MONTHS": true, "ONE_YEAR": true}

	if validPeriods[cfg.Settings.LookbackPeriod] {
	} else {
		return &ValidationError{
			Field: "LookbackPeriod",
			Msg:   "must be one of ONE_WEEK, ONE_MONTH, THREE_MONTHS, ONE_YEAR",
		}
	}

	if len(cfg.Assets) == 0 {
		return &ValidationError{
			Field: "Universe",
			Msg:   "must contain at least one asset",
		}
	}

	for i, asset := range cfg.Assets {
		if asset.UnderlyingID == "" {
			return &ValidationError{
				Field: fmt.Sprintf("Universe[%d].AssetID", i),
				Msg:   "cannot be empty",
			}
		}
	}

	return nil
}

func InitLogger() {
	// Open log file
	logFile, err := os.OpenFile("momBot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("❌ Failed to open log file:", err)
	}

	// Log to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)                           // Apply globally
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Add timestamp and file info
}
