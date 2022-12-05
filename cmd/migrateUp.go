package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate cmd up",
	Long:  "migrate command to handle sql migration up",
	Run: func(cmd *cobra.Command, agrs []string) {
		fmt.Println("Running migrate up command")
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
