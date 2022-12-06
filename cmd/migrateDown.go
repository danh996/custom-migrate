package cmd

import (
	"fmt"

	"github.com/danh996/go-custom-migrate/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate cmd up",
	Long:  "migrate command to handle sql migration down",
	Run: func(cmd *cobra.Command, agrs []string) {
		fmt.Println("Running migrate down command")
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

		err = m.Down() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			fmt.Printf("error when Down %v \n", err)
		}
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
