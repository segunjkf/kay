/*
Copyright © 2023 SAKA-AIYEDUN SEGUN sege.timz12@gmail.com
*/
package tool

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	file       string
	shaType    string
	shaFileLoc string
)

var shasumCmd = &cobra.Command{
	Use:   "shasum [file] [shaType]",
	Short: "Generates SHA checksum of a file",
	Long:  `The shasum command calculates and prints the SHA checksum (SHA1 or SHA256) for a specified file. It requires a file path and the type of SHA sum to generate (sha256 or sha1).`,
	Run: func(cmd *cobra.Command, args []string) {

		f, err := openfile(file)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer f.Close()

		if shaType == "sha256" {
			sha256sum, err := gensha256sum(f)
			if err != nil {
				logrus.Errorf("Failed to generate SHA256 sum for '%s': %v", file, err)
			} else {
				fmt.Println("SHA256:", sha256sum)
			}
		} else if shaType == "sha1" {
			sha1sum, err := gensha1sum(f)
			if err != nil {
				logrus.Errorf("Failed to generate SHA1 sum for '%s': %v", file, err)
			} else {
				fmt.Println("SHA1:", sha1sum)
			}
		} else {
			logrus.Errorf("Unknown hash algorithm '%s'. Please select either 'sha256' or 'sha1'.", shaType)
		}
		// if
	},
}

// Function for Opening file from the user input
func openfile(file string) (*os.File, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to open the file: %v", err)
	}

	return f, nil
}

// generate a shasum form the file
func gensha1sum(file *os.File) (string, error) {

	h := sha1.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("unable to generate create new hash: %v", err)
	}

	fmt.Printf("%x", h.Sum(nil))
	return fmt.Sprintf("sha256sum:%x\n", h.Sum(nil)), nil
}

func gensha256sum(file *os.File) (string, error) {

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("unable to generate create new hash: %v", err)
	}

	fmt.Printf("%x", h.Sum(nil))
	return fmt.Sprintf("sha256sum:%x\n", h.Sum(nil)), nil
}

func init() {
	ToolCmd.AddCommand(shasumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	shasumCmd.Flags().StringVarP(&file, "file", "f", "", "Path to the file")
	if err := shasumCmd.MarkFlagRequired("file"); err != nil {
		fmt.Println(err)
	}
	shasumCmd.PersistentFlags().StringVarP(&shaType, "type", "t", "sha256", "Type of SHA sum to generate (sha256 or sha1)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shasumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
