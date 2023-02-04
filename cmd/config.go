/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage Jenkins client configuration",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
