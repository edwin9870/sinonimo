package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sinonimo",
	Short: "Sinonimo is a cli app for searching spanish synonyms",
	Long:  "A fast and easy way to search spanish synonyms words",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
