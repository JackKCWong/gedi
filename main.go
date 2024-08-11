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

		parallel, _ := cmd.Flags().GetInt("parallel")
		g := internal.New(reader, internal.Filter{
			Expr:     args[0],
			Parallel: parallel,
		})

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
	rootCmd.Flags().IntP("parallel", "p", 1, "number of parallelism")
	// rootCmd.Flags().StringP("mode", "m", "filter", "operation mode, can be filter|map|reduce")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
