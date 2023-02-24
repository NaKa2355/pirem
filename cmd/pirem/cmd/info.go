/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	int_cmd "pirem/internal/app/pirem/cmd"

	apiremv1 "github.com/NaKa2355/pirem_pkg/apirem.v1"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
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
	conn, client, err := int_cmd.MakeConnection(int_cmd.Protocol, int_cmd.DomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetAllDeviceInfo(context.Background(), &apiremv1.GetAllDeviceInfoRequest{})
	if err != nil {
		return err
	}

	result, _ := int_cmd.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func getDeviceInfo(deviceID string) error {
	conn, client, err := int_cmd.MakeConnection(int_cmd.Protocol, int_cmd.DomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetDeviceInfo(context.Background(), &apiremv1.GetDeviceInfoRequest{DeviceId: deviceID})
	if err != nil {
		return err
	}

	result, _ := int_cmd.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	infoCmd.Flags().StringP("device", "d", "", "device id")
}
