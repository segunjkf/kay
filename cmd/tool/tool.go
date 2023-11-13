/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package tool

import (
	"github.com/spf13/cobra"
)

// toolCmd represents the tool command
var ToolCmd = &cobra.Command{
	Use:   "tool",
	Short: "A sub pallet for kay tool based commands",
	Long:  `A sub pallet for kay tool based commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Custom logic to generate shasum

func init() {
	// rootCmd.AddCommand(toolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
