/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addPswdCmd represents the addPswd command
var addPswdCmd = &cobra.Command{
	Use:   "addPswd",
	Short: "Command to add Account/Password",
	Long: `This command makes you able to add accounts usernames and passwords to your database
		encrypted ofcourse :)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addPswd called")
	},
}

func init() {
	rootCmd.AddCommand(addPswdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPswdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPswdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
