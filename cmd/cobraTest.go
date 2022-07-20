package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	// cobra.
	var (
		serveCmd = &cobra.Command{
			Use:   "serve",
			Short: "Short",
			Long:  "Long",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("RUN func")
			},
		}
	)
	rootCmd.AddCommand(serveCmd)
}
