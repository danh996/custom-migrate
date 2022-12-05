package cmd

import (
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate cmd",
	Long:  "migrate command to handle sql migration",
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
