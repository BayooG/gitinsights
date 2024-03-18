package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var verbose bool

	rootCmd := &cobra.Command{
		Use:   "gitinsights",
		Short: "A simple CLI tool for git insights",
		Run: func(cmd *cobra.Command, args []string) {
			// Your command logic goes here
			if verbose {
				fmt.Println("Verbose mode enabled")
			}
			fmt.Println("Hello from mycli!")
		},
	}

	// Add help flag
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})

	// Add verbose flag
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
