package cmd

import (
	"dop/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_TOLOWER = iota + 1
	MODE_TOUPPER
)

var desc = strings.Join([]string{
	"Subcommand for word conversion mode:",
	"1: ToLower",
	"2：ToUpper",
}, "\n")

var w string
var m int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word cli",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var res string
		switch m {
		case MODE_TOLOWER:
			res = word.ToLower(w)
		case MODE_TOUPPER:
			res = word.ToUpper(w)
		default:
			log.Fatalf("Unsupported conversion mode")
		}
		log.Printf("result：\n%s", res)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&w, "word", "w", "", "word content")
	wordCmd.Flags().Int8VarP(&m, "mode", "m", 1, "word conversion mode")
}
