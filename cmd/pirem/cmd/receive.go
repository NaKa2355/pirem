/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"time"

	pirem "github.com/NaKa2355/pirem/api/gen/go/api/v1"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "receive infrared from selected device",
	Long:  `receive infrared from selected device`,
	RunE:  cmd,
}

func receiveRawIr(deviceId string) (*pirem.IrData, error) {
	var irData *pirem.IrData
	conn, client, err := MakeConnection(Protocol, DomainSocketPath)
	if err != nil {
		return irData, err
	}
	defer conn.Close()

	irData, err = client.ReceiveIr(context.Background(), &pirem.ReceiveIrRequest{DeviceId: deviceId})
	if err != nil {
		return irData, err
	}
	return irData, nil
}

func progress(progChan chan any) {
	t := time.NewTicker(1 * time.Second)
	fmt.Print("receiving infrared data")
	for {
		select {
		case <-progChan:
			return
		case <-t.C:
			fmt.Print(".")
		}
	}
}

func cmd(cmd *cobra.Command, args []string) error {
	var result []byte
	var err error = nil
	progChan := make(chan any)
	outFlag := cmd.Flag("out")
	jsonFlag := cmd.Flag("json")
	progFlag := cmd.Flag("progress")
	devFlag := cmd.Flag("device")

	if !devFlag.Changed {
		return errors.New("no input device id")
	}

	if progFlag.Changed {
		go progress(progChan)
		defer close(progChan)
	}

	irData, err := receiveRawIr(devFlag.Value.String())
	if err != nil {
		return err
	}
	if progFlag.Changed {
		progChan <- new(any)
	}

	if jsonFlag.Changed {
		jsonStr, _ := MarshalToString(irData)
		jsonStr += "\n"
		result = []byte(jsonStr)
	} else {
		result, _ = proto.Marshal(irData)
	}

	if outFlag.Changed {
		f, err := os.Create(outFlag.Value.String())
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(result)
		if err != nil {
			return err
		}
	} else {
		if progFlag.Changed {
			fmt.Println("")
		}
		binary.Write(os.Stdout, binary.LittleEndian, result)
	}

	if progFlag.Changed {
		fmt.Println("Received IR signal!")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(receiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	receiveCmd.Flags().BoolP("progress", "p", false, "display progress and result")
	receiveCmd.Flags().BoolP("json", "j", false, "output as json format")
	receiveCmd.Flags().StringP("out", "o", "", "output to file path")
	receiveCmd.Flags().StringP("device", "d", "", "device id")
}
