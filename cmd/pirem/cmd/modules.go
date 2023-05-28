/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NaKa2355/pirem/build"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "get modules name which loaed",
	Long:  `get module names which loaed. You can write module names to json config file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		for k := range build.Modules {
			fmt.Printf("%s\n", k)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modulesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
