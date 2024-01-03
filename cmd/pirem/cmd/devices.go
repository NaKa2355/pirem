/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pirem "github.com/NaKa2355/pirem-proto/gen/go/api/v1"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "get device(s) information",
	Long:  `get device(s) information`,
	RunE: func(cmd *cobra.Command, args []string) error {
		devFlag := cmd.Flag("device")
		if devFlag.Changed {
			return getDeviceInfo(devFlag.Value.String())
		}
		return getAllDevsinfo()
	},
}

func getAllDevsinfo() error {
	conn, client, err := MakeConnection(Protocol, DomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.ListDevices(context.Background(), &pirem.ListDevicesRequest{})
	if err != nil {
		return err
	}

	result, _ := MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func getDeviceInfo(deviceID string) error {
	conn, client, err := MakeConnection(Protocol, DomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetDevice(context.Background(), &pirem.GetDeviceRequest{DeviceId: deviceID})
	if err != nil {
		return err
	}

	result, _ := MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(devicesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	devicesCmd.Flags().StringP("device", "d", "", "device id")
}
