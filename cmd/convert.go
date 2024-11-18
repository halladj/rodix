/*
Copyright © 2024 NAME HERE Hamza.halladj.cs@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	outputFile string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert media file into other file type.",
	Long:  `convert the media file extention { either video or images } into a diseared other extention.`,
	RunE: func(cmd *cobra.Command, args []string) error {

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
		&outputFile,
		"output",
		"o",
		"",
		"Output file extention(required)",
	)

	convertCmd.MarkPersistentFlagRequired("output")
}

func cleanString(s string) (string, error) {

	// TODO: use regix to verify that it has a proper form.

	// verify if this class support empty string, if it does change it.
	r, _ := regexp.Compile("[[:alpha:]].[[:alpha:]]")
	b := r.MatchString(s)
	if !b {
		return "", errors.New("file does not have a valid name")
	}

	newString := strings.Split(s, ".")
	if len(newString) < 1 {
		return "", errors.New("error")
	}
	// TODO: add regix for the supported extention
	// complete the rest of extentions list
	r, _ = regexp.Compile("jpeg|png|jpg")

	return newString[0], nil
}

func checkOutExtention(s string) string {

	// TODO implement using regix
	return s
}
