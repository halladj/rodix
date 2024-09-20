/*
Copyright © 2024 NAME HERE Hamza.halladj.cs@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	inputFile  string
	outputFile string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert media file into other file type.",
	Long:  `convert the media file extention { either video or images } into a diseared other extention.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//if inputFile == "" {
		//return errors.New("input file is required")
		//}

		//if outputFile == "" {
		//return errors.New("output file is required")
		//}

		// Check if input file exists
		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			return fmt.Errorf("input file does not exist: %s", inputFile)
		}

		// Use ffmpeg-go to convert the video
		fmt.Printf("Converting video from %s to %s\n", cleanString(inputFile), outputFile)

		err := ffmpeg.Input(inputFile).
			Output(cleanString(inputFile) + "_converted." + outputFile).
			Run()

		if err != nil {
			return fmt.Errorf("FFmpeg failed: %v", err)
		}

		fmt.Println("Video conversion complete!")
		return nil
	},
}

func init() {

	convertCmd.PersistentFlags().StringVarP(
		&inputFile,
		"input",
		"i",
		"",
		"Input file name (required)",
	)
	convertCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		"",
		"Output file extention(required)",
	)

	convertCmd.MarkPersistentFlagRequired("input")
	convertCmd.MarkPersistentFlagRequired("output")
}

func cleanString(s string) string {
	newString := strings.Split(s, ".")
	if len(newString) < 1 {
		return ""
	}

	// TODO: use regix to verify that it has a proper form.
	// TODO: add regix for the supported extention

	return newString[0]
}

func checkOutExtention(s string) string {

	// TODO implement using regix
	return s
}
