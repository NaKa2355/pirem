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
var remotesCmd = &cobra.Command{
	Use:   "remotes",
	Short: "get remote(s)",
	Long:  `get remote(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		remoteFalg := cmd.Flag("remote")
		if remoteFalg.Changed {
			return getRemote(remoteFalg.Value.String())
		}
		return listRemotes()
	},
}

func listRemotes() error {
	conn, client, err := utils.MakeConnection(utils.Protocol, utils.GrpcDomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.ListRemotes(context.Background(), &pirem.ListRemotesRequest{})
	if err != nil {
		return err
	}

	result, _ := utils.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func getRemote(remoteID string) error {
	conn, client, err := utils.MakeConnection(utils.Protocol, utils.GrpcDomainSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetRemote(context.Background(), &pirem.GetRemoteRequest{RemoteId: remoteID})
	if err != nil {
		return err
	}

	result, _ := utils.MarshalToString(resp)
	fmt.Println(result)
	return nil
}

func init() {
	rootCmd.AddCommand(remotesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	remotesCmd.Flags().StringP("remote", "r", "", "remote id")
}
