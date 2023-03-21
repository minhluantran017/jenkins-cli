/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/minhluantran017/jenkins-cli/internal/client"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update Jenkins client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		client.UpdateProfile()
	},
}

func init() {
	configCmd.AddCommand(updateCmd)
}
