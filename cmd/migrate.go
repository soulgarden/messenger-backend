package cmd

import "github.com/spf13/cobra"

//nolint: gochecknoglobals,exhaustivestruct
var migrateCmd = &cobra.Command{
	Use:   "migrates",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
