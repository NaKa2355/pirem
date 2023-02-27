/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"

	apiremv1 "github.com/NaKa2355/irdeck-proto/gen/go/pirem/api/v1"
	int_cmd "github.com/NaKa2355/pirem/internal/app/pirem/cmd"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "get device status",
	Long:  `get device status`,
	RunE: func(cmd *cobra.Command, args []string) error {
		devFlag := cmd.Flag("device")
		if !devFlag.Changed {
			return errors.New("no input device id")
		}
		return getDeviceStatus(devFlag.Value.String())
	},
}

func getDeviceStatus(deviceID string) error {
	conn, client, err := int_cmd.MakeConnection(int_cmd.Protocol, int_cmd.DomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetDeviceStatus(context.Background(), &apiremv1.GetDeviceStatusRequest{DeviceId: deviceID})
	if err != nil {
		return err
	}

	result, _ := int_cmd.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")
	statusCmd.Flags().StringP("device", "d", "", "device id")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
