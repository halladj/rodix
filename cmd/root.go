/*
Copyright © 2024 NAME HERE Hamza.Halladj.cs@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var inputFile string

var rootCmd = &cobra.Command{
	Use:   "rodix",
	Short: "",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&inputFile,
		"input",
		"i",
		"",
		"Input file name (required)",
	)
	rootCmd.AddCommand(rotateCmd)
	rootCmd.AddCommand(convertCmd)
	rootCmd.MarkPersistentFlagRequired("input")
}
