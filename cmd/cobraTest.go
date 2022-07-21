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
				fmt.Println("RUN fuck")
			},
		}
	)
	rootCmd.AddCommand(serveCmd)

	// 實測 無法一次執行兩個subcommand
	var serveCmd2 = &cobra.Command{
		Use:   "serve2",
		Short: "Short2",
		Long:  "Long2",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("RUN fuck 2")
		},
	}
	rootCmd.AddCommand(serveCmd2)

	// 實作可參考我寫的 chat_room包 的match-live
}
