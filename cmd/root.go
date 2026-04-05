package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoVault",
	Short: "Password Manager tool with good encryption",
	Long: `GoVault is a a self hostable secure password manager tool using good encryption
		to store your passwords and that of other users on your own laptop securly`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
