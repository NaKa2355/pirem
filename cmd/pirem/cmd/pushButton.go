/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pirem "github.com/NaKa2355/pirem-proto/gen/go/api/v1"
	"github.com/NaKa2355/pirem/cmd/pirem/utils"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pushButtonCmd = &cobra.Command{
	Use:   "push",
	Short: "push button",
	Long:  `push button`,
	RunE: func(cmd *cobra.Command, args []string) error {
		buttonFlag := cmd.Flag("button")
		if buttonFlag.Changed {
			return pushButton(buttonFlag.Value.String())
		}
		return fmt.Errorf("need argument")
	},
}

func pushButton(buttonID string) error {
	conn, client, err := utils.MakeConnection(utils.Protocol, utils.GrpcDomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.PushButton(context.Background(), &pirem.PushButtonRequest{
		ButtonId: buttonID,
	})
	if err != nil {
		return err
	}

	result, _ := utils.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(pushButtonCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pushButtonCmd.Flags().StringP("button", "b", "", "button id")
}
