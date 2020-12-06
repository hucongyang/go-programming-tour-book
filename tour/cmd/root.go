package cmd

import "github.com/spf13/cobra"

/**
 * 用于放置根命令
 */
var rootCmd = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init()  {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}