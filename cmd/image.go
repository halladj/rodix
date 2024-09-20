/*
Copyright © 2024 NAME HERE Hamza.Halladj.cs@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image called")
	},
}

func init() {
	//rootCmd.AddCommand(imageCmd)
}
