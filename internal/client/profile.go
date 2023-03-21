package client

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type Profile struct {
	Url      string `toml:"url"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

// Get the Profile by name
func GetProfile(profileName string) Profile {
	var profile Profile
	if err := viper.UnmarshalKey(profileName, &profile); err != nil {
		log.Fatal("Error unmarshaling profile:", err)
	}
	return profile
}

// View Jenkins client profile
func ViewProfile() {
	profile := GetProfile(viper.GetString("profile"))
	log.Info("Jenkins URL: ", profile.Url)
	log.Info("Jenkins username: ", profile.UserName)
}

// Update Jenkins profile value(s) from user input
func UpdateProfile() {
	profileName := viper.GetString("profile")
	log.Infof("Updating Jenkins profile '%s'...", profileName)
	profile := GetProfile(profileName)
	
	// Create new Viper instance for configuration file update only
	new_viper := viper.New()
	new_viper.SetConfigFile(viper.ConfigFileUsed())
	new_viper.ReadInConfig()

	// Read in URL
	fmt.Printf("Jenkins URL [%s]: ", profile.Url)
	var url string
	fmt.Scanln(&url)
	if url == "" {
		url = profile.Url
	}

	// Read in username
	fmt.Printf("Jenkins Username [%s]: ", profile.UserName)
	var username string
	fmt.Scanln(&username)
	if username == "" {
		username = profile.UserName
	}

	// Read in password
	fmt.Print("Password: ")
	password, err := term.ReadPassword(0)
	if err != nil {
		log.Fatal("Error reading password:", err)
	}
	if string(password) == "" {
		password = []byte(profile.Password)
	}
	
	// Write to config file
	fmt.Println()
	profile.Url = url
	profile.UserName = username
	profile.Password = string(password)
	new_viper.Set(profileName, profile)
	if err := new_viper.WriteConfig(); err != nil {
		log.Fatal("Error writing configuration file:", err)
	}
	log.Infof("Profile '%s' updated successfully.", profileName)
}
