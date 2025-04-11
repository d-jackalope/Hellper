package config

import (
	"errors"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// This package is designed for centralized storage of configurations required for service operation
// Use the `Get***()` functions from the `config` package to retrieve values from environment variables.

// Don’t forget to declare any new environment variables in the `config` package if needed.
type Config struct {
	AI     *AIConfig
	Bot    *BotConfig
	DBLink string `env:"DB_LINK,required,notEmpty"`
	TGKey  string `env:"TG_KEY,required,notEmpty"`
}

type AIConfig struct {
	ModelsListEndpoint       string `env:"DB_LINK,required,notEmpty"`
	ImageGenerationModel     string `env:"DB_LINK,required,notEmpty"`
	ImageGenerationEndpoint  string `env:"DB_LINK,required,notEmpty"`
	ImageRecognitionModel    string `env:"DB_LINK,required,notEmpty"`
	ImageRecognitionEndpoint string `env:"DB_LINK,required,notEmpty"`
	VoiceRecognitionModel    string `env:"DB_LINK,required,notEmpty"`
	VoiceRecognitionEndpoint string `env:"DB_LINK,required,notEmpty"`
}

type BotConfig struct {
	Admin *Admin
}

type Admin struct {
	Login    string
	Password string
}

var cfg *Config

// Environment variable loading.
// The first function is for initializing environment variables.
// Ensures that no variable is left empty.
func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("failed to load env")
	}
	if err := env.Parse(&cfg); err != nil {
		return errors.New("failed to load config")
	}
	log.Info().Msg("Environment variables loaded successfully.")
	return nil
}

func GetConfig() *Config {
	return cfg
}

func GetAI() *AIConfig {
	return GetConfig().AI
}

func GetBot() *BotConfig {
	return GetConfig().Bot
}

func GetDBLink() string {
	return GetConfig().DBLink
}

func GetTGKEY() string {
	return GetConfig().TGKey
}
