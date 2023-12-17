package config

var DefaultGatewayConfig = Config{
	Environment: "dev",
	Port:        "3000",
}

type Config struct {
	Environment string `env:"ENVIRONMENT"`
	Port        string `env:"PORT"`
}
