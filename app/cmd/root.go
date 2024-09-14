package cmd

import (
	"chat-api/app/config"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
)

var (
	EnvFilePath string
	rootCmd     = &cobra.Command{
		Use:   "cobra-cli",
		Short: "dummy-server",
	}
)
var (
	rootConfig *config.Root
	database   *gorm.DB
)

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&EnvFilePath, "env", "e", ".env", ".env file to read from")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Cannot Run CLI. err > ", err)
		os.Exit(1)
	}
}
func configReader() {
	log.Infof("Initialize ENV")
	rootConfig = config.Load(EnvFilePath)
}
func init() {
	cobra.OnInitialize(func() {
		configReader()
		initPosgresDb()
	})
}
func initPosgresDb() {
	log.Infof("Initialize postgres")
	database = config.OpenPostgresDatabaseConnection(config.Postgres{
		Host:                  rootConfig.Postgres.Host,
		Port:                  rootConfig.Postgres.Port,
		User:                  rootConfig.Postgres.User,
		Password:              rootConfig.Postgres.Password,
		Dbname:                rootConfig.Postgres.Dbname,
		MaxConnectionLifetime: rootConfig.Postgres.MaxConnectionLifetime,
		MaxOpenConnection:     rootConfig.Postgres.MaxOpenConnection,
		MaxIdleConnection:     rootConfig.Postgres.MaxIdleConnection,
	})
}
