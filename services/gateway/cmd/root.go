package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/mamalovesyou/claimclam/internal/env"
	"github.com/mamalovesyou/claimclam/internal/logging"
	"github.com/mamalovesyou/claimclam/services/gateway/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	gatewayCfg = config.DefaultGatewayConfig

	rootCmd = &cobra.Command{
		Use:   "gateway",
		Short: "API Gateway",
		Long:  `API Gateway is a rest proxy server that proxy graphql requests to the right service.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := env.ReadEnv(&gatewayCfg, false)
			if err != nil {
				logging.Fatalf(ctx, "Unable to parse config: %v", err)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
