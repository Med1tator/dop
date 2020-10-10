package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time cli",
	Long:  desc,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "now cli",
	Long:  `get time now`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		fmt.Printf("Local Time：\t%s\n", now.Format("2006-01-02 15:04:05"))
		fmt.Printf("UTC Time：\t%s\n", now.UTC().Format("2006-01-02 15:04:05"))
	},
}

func init() {
	timeCmd.AddCommand(timeNowCmd)
}
