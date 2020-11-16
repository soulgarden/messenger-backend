package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//nolint: gochecknoglobals, exhaustivestruct
var rootCmd = &cobra.Command{
	Use:   "messenger",
	Short: "Messenger provides functionality of creating group chats and messaging between users",
}

func Execute() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
