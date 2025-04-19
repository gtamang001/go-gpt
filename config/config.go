package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAIAPIKey string
	OpenAIURL    string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, add .env file for loading env variables")
	}
	AppConfig = Config{
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", "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"),
		OpenAIURL:    getEnv("OPENAI_URL", "https://api.openai.com/v1/chat/completions"),
	}

	if AppConfig.OpenAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in environment variables")
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
