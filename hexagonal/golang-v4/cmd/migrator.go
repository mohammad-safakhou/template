package cmd

import (
	"fmt"
	"git.siz-tel.com/charging/template/config"
	"git.siz-tel.com/charging/template/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"git.siz-tel.com/charging/template/internal/repository/database"
	"git.siz-tel.com/charging/template/logger"
	"github.com/spf13/cobra"
)

func getSourceURL(path string) string {
	if path != "" {
		return path
	}
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error: can not find project root path", err)
	}

	var sourceURL string

	pattern := regexp.MustCompile(`^.*\.(down|up)\.sql$`)

	err = filepath.Walk(workingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		match := pattern.MatchString(info.Name())
		if match {
			sourceURL = filepath.Dir(path)
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	if sourceURL == "" {
		log.Println("Can not find the SourceURL of database schemes directory.")
		log.Fatalln("Please ensure you are in the root of the project when run or build the project.")
	}
	return sourceURL
}

var migratorCmd = &cobra.Command{
	Use:   "migrator",
	Short: "Manages your database migrations",
	Long: `The migrator command allows you to manage your database migrations.
You can apply all up migrations using 'migrator up' or apply all down migrations using 'migrator down'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use 'migrator up' to apply all up migrations or 'migrator down' to apply all down migrations.")
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all up migrations",
	Long: `The 'migrator up' command applies all up migrations to your database.
	This will update your database schema to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("Applying all up migrations...")
		path := cmd.Flag("path").Value.String()
		database.MigrateUp("file://"+getSourceURL(path),
			utils.PostgresUrl(
				config.AppConfig.Databases.Postgres.Host,
				config.AppConfig.Databases.Postgres.Port,
				config.AppConfig.Databases.Postgres.User,
				config.AppConfig.Databases.Postgres.Pass,
				config.AppConfig.Databases.Postgres.DatabaseName,
				config.AppConfig.Databases.Postgres.SslMode,
			),
		)
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Apply all down migrations",
	Long: `The 'migrator down' command applies all down migrations to your database.
	This will revert your database schema to the previous version.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("Applying all down migrations...")
		path := cmd.Flag("path").Value.String()
		database.MigrateDown("file://"+getSourceURL(path),
			utils.PostgresUrl(
				config.AppConfig.Databases.Postgres.Host,
				config.AppConfig.Databases.Postgres.Port,
				config.AppConfig.Databases.Postgres.User,
				config.AppConfig.Databases.Postgres.Pass,
				config.AppConfig.Databases.Postgres.DatabaseName,
				config.AppConfig.Databases.Postgres.SslMode,
			),
		)
	},
}

func init() {
	upCmd.PersistentFlags().StringP("path", "p", "", "path to database schemes directory")
	migratorCmd.AddCommand(upCmd)
	migratorCmd.AddCommand(downCmd)
}
