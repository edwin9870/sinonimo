package cmd

import (
	"fmt"
	"os"

	"github.com/edwin9870/sinonimo/internal/sinonimo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sinonimo",
	Short: "Sinonimo is a cli app for searching spanish synonyms",
	Long:  "A fast and easy way to search spanish synonyms words",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		w := sinonimo.WordReference{}
		result, err := w.Search(args[0])
		if err != nil {
			return err
		}

		if len(result) == 0 {
			fmt.Println("Not match were found")
			return nil
		}

		for _, v := range result {
			fmt.Println(v)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
