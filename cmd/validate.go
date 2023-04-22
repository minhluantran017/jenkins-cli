/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/minhluantran017/jenkins-cli/internal/server"
	"github.com/spf13/cobra"
)

var JenkinsfilePath string

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate Jenkinsfile",
	Run: func(cmd *cobra.Command, args []string) {
		server.ValidateJenkinsfile(JenkinsfilePath)
	},
}

func init() {
	pipelineCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVarP(&JenkinsfilePath, "file", "f", "", "Jenkinsfile path to validate")
	validateCmd.MarkFlagRequired("file")
}
