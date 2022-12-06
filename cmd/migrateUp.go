package cmd

import (
	"fmt"

	"github.com/danh996/go-custom-migrate/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate cmd up",
	Long:  "migrate command to handle sql migration up",
	Run: func(cmd *cobra.Command, agrs []string) {
		fmt.Println("Running migrate up command")
		db := database.Open()
		dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			fmt.Printf("error when create instance %v \n", err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://./migrations",
			"postgres", dbDriver)

		if err != nil {
			fmt.Printf("error when NewWithDatabaseInstance %v \n", err)
		}

		err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			fmt.Printf("error when Up %v \n", err)
		}
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
