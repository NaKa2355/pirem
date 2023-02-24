/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	int_cmd "github.com/NaKa2355/pirem/internal/app/pirem/cmd"
	apirem_v1 "github.com/NaKa2355/pirem/pkg/apirem.v1"

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
		var irData = &apirem_v1.RawIrData{}

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

		conn, client, err := int_cmd.MakeConnection(int_cmd.Protocol, int_cmd.DomainSocketPath)
		if err != nil {
			return err
		}
		defer conn.Close()

		_, err = client.SendRawIr(context.Background(), &apirem_v1.SendRawIrRequest{DeviceId: devFlag.Value.String(), IrData: irData})
		if err != nil {
			return err
		}

		fmt.Printf("sent %d pluses\n", len(irData.OnOffPluseNs))
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
