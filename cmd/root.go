/*
Copyright © 2023 Luan Tran <minhluantran017@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/minhluantran017/jenkins-cli/internal/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	profile string
)

var rootCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Jenkins commandline utilities",
	Long: `Jenkins commandline utilities.
This is the mini CLI application, written in Golang,
for interacting with Jenkins server.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "", "Jenkins profile to use")
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	configFile := home + "/.jenkins/config.toml"
	viper.SetConfigType("toml")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv() 
	viper.SetEnvPrefix("JENKINS")
	helper.HandleError(viper.BindEnv("profile"))
	
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file: ", viper.ConfigFileUsed())
		if profileName := viper.GetString("profile"); profileName == "" {
			log.Debug("Profile is not set. Using 'default'.")
			viper.Set("profile", "default")
		}
		log.Debug("Using config profile: ", viper.GetString("profile"))
	}
}
