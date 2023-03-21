/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/minhluantran017/jenkins-cli/internal/client"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view Jenkins client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		client.ViewProfile()
	},
}

func init() {
	configCmd.AddCommand(viewCmd)
}
