/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	int_cmd "github.com/NaKa2355/pirem/internal/app/pirem/cmd"
	"github.com/NaKa2355/pirem/internal/app/pirem/daemon"

	"github.com/spf13/cobra"
)

const ConfigFilePath = "/etc/piremd.json"

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "execute as daemon",
	Long: `execute as daemon. 
	config file: /etc/piremd.json
	service file: /lib/systemd/system/piremd.service
	socket file: /tmp/pirem.sock`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startDaemon()
	},
}

func startDaemon() error {
	d, err := daemon.New(int_cmd.ConfigFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faild to start daemon: %s\n", err)
		return err
	}
	d.Start(int_cmd.DomainSocketPath)
	return nil
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
