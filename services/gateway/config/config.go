package config

var DefaultGatewayConfig = Config{
	Environment:         "dev",
	Port:                "3000",
	PodcastsServiceAddr: "accounts:5003",
}

type Config struct {
	Environment         string `env:"ENVIRONMENT"`
	Port                string `env:"PORT"`
	PodcastsServiceAddr string `env:"PODCAST_SERVICE_ADDR",required"`
}
