/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"

	int_cmd "github.com/NaKa2355/pirem/internal/app/pirem/cmd"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"

	"github.com/spf13/cobra"
)

// isbusyCmd represents the isbusy command
var isbusyCmd = &cobra.Command{
	Use:   "isbusy",
	Short: "get if device is busy",
	Long:  `get id device is busy`,
	RunE: func(cmd *cobra.Command, args []string) error {
		devFlag := cmd.Flag("device")
		if !devFlag.Changed {
			return errors.New("no input device id")
		}
		conn, client, err := int_cmd.MakeConnection(int_cmd.Protocol, int_cmd.DomainSocketPath)
		if err != nil {
			return err
		}
		defer conn.Close()

		resp, err := client.IsBusy(context.Background(), &apiremv1.IsBusyRequest{DeviceId: devFlag.Value.String()})
		if err != nil {
			return err
		}
		fmt.Println(resp.IsBusy)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(isbusyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isbusyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	isbusyCmd.Flags().StringP("device", "d", "", "device id")
}
