/*
Copyright © 2023 Saka-Aiyedin Segun sege.timz12@gmail.com
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	inputUrl string
	// custom client logic
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response",
	Long:  `This pings a remote URL and returns the response`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

		// use the ping function
		if response, err := ping(inputUrl); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(response)
		}
	},
}

func ping(domain string) (int, error) {
	url := "http://" + domain
	newRequest, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, fmt.Errorf("unable to create new request: %v", err)
	}

	requestResponse, err := client.Do(newRequest)
	if err != nil {
		return 0, fmt.Errorf("unable to ping url: %v", err)
	}
	defer requestResponse.Body.Close()

	return requestResponse.StatusCode, nil
}

func init() {
	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	pingCmd.PersistentFlags().StringVarP(&inputUrl, "url", "u", "", "custom Url to ping")
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
