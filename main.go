package main

import (
	"fmt"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "gedi <expr>",
	Short: "gedi is a simple file editing tool using a expression",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.Flags().StringP("type", "t", "line", "file type of the input file, can be line|csv|jsonl|jsonarray")
	rootCmd.Flags().StringP("file", "f", "", "path to the input file. If not specified, stdin will be used.")
	rootCmd.Flags().StringP("mode", "m", "filter", "operation mode, can be filter|map|reduce")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}


