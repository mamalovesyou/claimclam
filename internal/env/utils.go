package env

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const ENVIRONMENT_KEY = "ENVIRONMENT"

// IsProductionEnv return true if environment variable ENV is set to prod | production
func IsProductionEnv() bool {
	env := os.Getenv(ENVIRONMENT_KEY)
	value := strings.ToUpper(env)
	return value == PROD || value == PRODUCTION
}

// IsDevelopmentEnv return true if environment variable ENV is set to dev | development
func IsDevelopmentEnv() bool {
	env := os.Getenv(ENVIRONMENT_KEY)
	value := strings.ToUpper(env)
	return value == DEV || value == DEVELOPMENT
}

// IsSandboxEnv return true if environment variable ENV is set to sandbox
func IsSandboxEnv() bool {
	env := os.Getenv(ENVIRONMENT_KEY)
	value := strings.ToUpper(env)
	return value == SANDBOX
}

// ReadEnv read environment variables and populate a struct with env tags
func ReadEnv(cfg interface{}, requiredIfNoDef bool) error {
	// Overload will replace existing env
	err := godotenv.Overload()
	if err != nil && !os.IsNotExist(err) {
		return errors.New("Error loading .env file")
	}

	if err := env.Parse(cfg, env.Options{RequiredIfNoDef: requiredIfNoDef}); err != nil {
		return errors.New(fmt.Sprintf("Failed to load config. %v", err))
	}

	return nil
}
