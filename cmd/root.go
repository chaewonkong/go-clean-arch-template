package cmd

import (
	"log"
	app "nass-admin-go/internal/_app"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const version = "2.0.0"

var rootCmd = &cobra.Command{
	Use:     "nass-admin-api",
	Version: version,
	Short:   "nass-admin-api CLI",
	Run: func(cmd *cobra.Command, args []string) {
		// fx를 이용해 Dependency Injection을 하며 서버 애플리케이션 구동
		fx.New(app.Modules).Run()
	},
}

// 서버 실행
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}