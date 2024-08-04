package main

import (
	"fmt"
	"io"
	"os"

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

		var reader internal.RecordReader
		filetype, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		switch filetype {
		case "line":
			reader = internal.LineReader{}
		case "csv":
			reader = &internal.CsvReader{}
		default:
			return fmt.Errorf("unknown file type: %s", filetype)
		}

		g := internal.New(reader, internal.Filter{Expr: args[0]})

		return g.Run(input)
	},
}

func init() {
	rootCmd.Flags().StringP("type", "t", "line", "file type of the input file, can be line|csv")
	rootCmd.Flags().StringP("file", "f", "", "path to the input file. If not specified, stdin will be used.")
	// rootCmd.Flags().StringP("mode", "m", "filter", "operation mode, can be filter|map|reduce")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
