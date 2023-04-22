/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pipelineCmd represents the pipeline command
var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "manage Jenkins pipelines",
}

func init() {
	rootCmd.AddCommand(pipelineCmd)
}
