package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/JackKCWong/gedi/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gedi <expr>",
	Args:  cobra.ExactArgs(1),
	Short: "gedi is a simple file editing tool using a expression",
	RunE: func(cmd *cobra.Command, args []string) error {
		var input io.Reader = os.Stdin
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		if file != "" {
			input, err = os.Open(file)
			if err != nil {
				return err
			}
		}

		filetype, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		if filetype == "auto" {
			if strings.HasSuffix(file, ".csv") {
				filetype = "csv"
			} else if strings.HasSuffix(file, ".jsonl") {
				filetype = "jsonl"
			} else if strings.HasSuffix(file, ".json") {
				filetype = "json"
			} else {
				filetype = "line"
			}
		}
		var reader internal.RecordReader
		switch filetype {
		case "line":
			reader = &internal.LineReader{}
		case "csv":
			reader = &internal.CsvReader{}
		case "jsonl":
			reader = &internal.JsonLReader{}
		case "json":
			reader = &internal.JsonReader{}
		default:
			return fmt.Errorf("unknown file type: %s", filetype)
		}

		var process internal.RecordProcessor
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}

		switch mode {
		case "auto":
			process, err = internal.InferProcess(args[0])
			if err != nil {
				return err
			}
		case "f":
			fallthrough
		case "filter":
			process = internal.Filter{Expr: args[0]}
		case "m":
			fallthrough
		case "map":
			process = internal.Mapper{Expr: args[0]}
		case "r":
			fallthrough
		case "reduce":
			process = internal.Reducer{Expr: args[0]}
		default:
			return fmt.Errorf("unknown mode: %s", mode)
		}

		g := internal.New(reader, process)
		err = g.Run(input)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringP("type", "t", "auto", "file type of the input file, can be line|csv|jsonl|json")
	rootCmd.Flags().StringP("file", "f", "", "path to the input file. If not specified, stdin will be used.")
	rootCmd.Flags().StringP("mode", "m", "auto", "operation mode, can be auto|f[ilter]|m[ap]|r[educe]")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
