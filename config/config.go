package config

import (
	"fmt"
	"jsol-ask/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabasePath     string
	TelegramBotToken string
	GroupID          uint64
	GeneralThread    int
	SuperAdminID     uint64
}

// LoadConfig loads the configuration from the .env file and returns Config with validation of all required values
func LoadConfig() *Config {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		utils.Log(utils.PanicLog, utils.SystemModule, "Error loading .env file: file not found!")
	}

	config := &Config{
		DatabasePath:     getEnvOrPanic("DATABASE_PATH"),
		TelegramBotToken: getEnvOrPanic("TELEGRAM_BOT_TOKEN"),
		GroupID:          parseUintEnv("GROUP_ID", 64),
		GeneralThread:    parseIntEnv("GENERAL_THREAD", 32),
		SuperAdminID:     parseUintEnv("SUPER_ADMIN_ID", 64),
	}

	return config
}

// getEnvOrPanic retrieves an environment variable or terminates execution if the variable is not set
func getEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		utils.Log(utils.PanicLog, utils.SystemModule, fmt.Sprintf("Environment variable %s is not set or empty", key))
	}
	return value
}

// parseUintEnv converts an environment variable to uint64 or terminates execution on error
func parseUintEnv(key string, bitSize int) uint64 {
	valueStr := getEnvOrPanic(key)
	value, err := strconv.ParseUint(valueStr, 10, bitSize)
	if err != nil {
		utils.Log(utils.PanicLog, utils.SystemModule, fmt.Sprintf("Error converting environment variable %s to uint64: %v", key, err))
	}
	return value
}

// parseIntEnv converts an environment variable to int or terminates execution on error
func parseIntEnv(key string, bitSize int) int {
	valueStr := getEnvOrPanic(key)
	value, err := strconv.ParseInt(valueStr, 10, bitSize)
	if err != nil {
		utils.Log(utils.PanicLog, utils.SystemModule, fmt.Sprintf("Error converting environment variable %s to int: %v", key, err))
	}
	return int(value)
}
