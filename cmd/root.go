package cmd

import (
	"log"
	"os"

	app "github.com/chaewonkong/go-clean-arch-template/internal/_app"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const version = "2.0.0"

var rootCmd = &cobra.Command{
	Use:     "go-clean-arch-template",
	Version: version,
	Short:   "go-clean-arch-template",
	Run: func(cmd *cobra.Command, args []string) {
		// Run server with dependencies injected by fx
		fx.New(app.Modules).Run()
	},
}

// Start server
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
