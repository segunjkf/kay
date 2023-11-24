/*
Copyright © 2023 SAKA-AIYEDUN SEGUN sege.timz12@gmail.com
*/
package tool

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	fileName     string
	compressType string
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress a given file.",
	Long:  `Compress command enables the compression of a specified file into either Gzip or Zlib format. It requires a file path as an argument and generates a compressed file in the same directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := OpenFile(fileName)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer f.Close()

		switch compressType {
		case "gzip":
			_, err = compressGzip(f)
			if err != nil {
				logrus.Error(err)
			}
		case "zlib":
			_, err = compressZlib(f)
			if err != nil {
				logrus.Error(err)
			}
		default:
			logrus.Errorf("Unknown compression type '%s'. Please select either 'gzip' or 'zlib'.", compressType)
		}
	},
}

func OpenFile(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func compressGzip(file *os.File) (*os.File, error) {
	compressFileName := fmt.Sprintf("%s.gz", file.Name())
	compressedFile, err := os.Create(compressFileName)
	if err != nil {
		return nil, err
	}
	gzFile := gzip.NewWriter(compressedFile)
	defer gzFile.Close()

	_, err = io.Copy(gzFile, file)
	if err != nil {
		return nil, err
	}

	if err := gzFile.Close(); err != nil {
		return nil, err
	}

	return compressedFile, nil
}

func compressZlib(file *os.File) (*os.File, error) {
	compressFileName := fmt.Sprintf("%s.zlib", file.Name())
	compressedFile, err := os.Create(compressFileName)
	if err != nil {
		return nil, err
	}
	zlibWriter := zlib.NewWriter(compressedFile)
	defer zlibWriter.Close()

	_, err = io.Copy(zlibWriter, file)
	if err != nil {
		return nil, err
	}

	if err := zlibWriter.Close(); err != nil {
		return nil, err
	}

	return compressedFile, nil
}

func init() {
	ToolCmd.AddCommand(compressCmd)

	compressCmd.Flags().StringVarP(&fileName, "filename", "f", "", "Path to the file to be compressed")
	if err := compressCmd.MarkFlagRequired("filename"); err != nil {
		logrus.Error(err)
	}

	compressCmd.Flags().StringVarP(&compressType, "type", "t", "gzip", "Type of compression (gzip or zlib)")
}
