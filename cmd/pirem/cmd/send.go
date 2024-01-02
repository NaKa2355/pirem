/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	pirem "github.com/NaKa2355/pirem/api/gen/go/api/v1"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send infrared from selected device",
	Long: `send infrared from selected device 
	to send data, use binary protobuf file with redirect`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileFlag := cmd.Flag("file")
		devFlag := cmd.Flag("device")
		var irData = &pirem.IrData{}

		if !devFlag.Changed {
			return errors.New("no input device id")
		}
		if !fileFlag.Changed {
			return errors.New("choose input file")
		}

		buf, err := os.ReadFile(fileFlag.Value.String())
		if err != nil {
			return err
		}

		err = proto.Unmarshal(buf, irData)
		if err != nil {
			return err
		}

		conn, client, err := MakeConnection(Protocol, DomainSocketPath)
		if err != nil {
			return err
		}
		defer conn.Close()

		_, err = client.SendIr(context.Background(), &pirem.SendIrRequest{DeviceId: devFlag.Value.String(), IrData: irData})
		if err != nil {
			return err
		}

		fmt.Println("sent IR signal!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendCmd.Flags().StringP("file", "f", "", "input file to send")
	sendCmd.Flags().StringP("device", "d", "", "device id")
}
