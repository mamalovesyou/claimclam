package cmd

import (
	"github.com/mamalovesyou/getclaim/internal/logging"
	"github.com/mamalovesyou/getclaim/internal/server"
	"github.com/mamalovesyou/getclaim/services/gateway/app"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start gateway service",
		Long:  `Start gateway service`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &server.Config{
				Port:           gatewayCfg.Port,
				Logger:         logging.NewLogger(),
				AllowedOrigins: []string{"*"},
			}
			app := app.NewApp(cfg)
			app.Run()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
