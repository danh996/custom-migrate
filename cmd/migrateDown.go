package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate cmd up",
	Long:  "migrate command to handle sql migration down",
	Run: func(cmd *cobra.Command, agrs []string) {
		fmt.Println("Running migrate down command")
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
