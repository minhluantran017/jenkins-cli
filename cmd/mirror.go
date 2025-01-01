/*
Copyright © 2025 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"github.com/minhluantran017/jenkins-cli/internal/server"
	"github.com/spf13/cobra"
)

var JenkinsfilePath string

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "mirror Jenkinsfile",
	Run: func(cmd *cobra.Command, args []string) {
		server.MirrorJenkinsfile(SourceProfile, DestinationProfile, PipelineName)
	},
}

func init() {
	pipelineCmd.AddCommand(mirrorCmd)
	mirrorCmd.Flags().StringVarP(&SourceProfile, "src", "s", "", "Jenkins profile to use as mirror source")
	mirrorCmd.Flags().StringVarP(&DestinationProfile, "dst", "d", "", "Jenkins profile to use as mirror destination")
	mirrorCmd.Flags().StringVarP(&PipelineName, "pipeline", "p", "", "Pipeline name to mirror")
	mirrorCmd.MarkFlagRequired("src")
	mirrorCmd.MarkFlagRequired("dst")
	mirrorCmd.MarkFlagRequired("pipeline")
}
