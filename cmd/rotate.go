/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// TODO: add a switch case to handle the different degrees.
var degree float32

// rotateCmd represents the rotate command
var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "rotates media files by +/- 90, 180, 270 degrees",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if input file exists
		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			return fmt.Errorf("input file does not exist: %s", inputFile)
		}

		outputFile := fmt.Sprintf("rotated_%s%s", "1", ".png")

		// Use ffmpeg-go to rotate the image
		fmt.Printf("Rotating image %s by %f degrees\n", inputFile, degree)

		// TODO: This code is trash check out docs.
		err := ffmpeg.Input(inputFile).
			Output(outputFile, ffmpeg.KwArgs{"vf": "transpose=2"}).
			ErrorToStdOut().
			Run()

		if err != nil {
			return fmt.Errorf("FFmpeg failed: %v", err)
		}

		fmt.Println("Image rotation complete!")
		return nil

	},
}

func init() {

	rotateCmd.PersistentFlags().Float32VarP(
		&degree,
		"degree",
		"d",
		0,
		"Degree to be rotated by (required).",
	)

	rotateCmd.MarkPersistentFlagRequired("degree")
}
