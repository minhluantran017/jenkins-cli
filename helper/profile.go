package helper

import (
	"github.com/spf13/viper"
)

type Profile struct {
	Name     string `mapstructure:"name"` 
	Url 		string `mapstructure:"url"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Get the Profile by name
func GetProfile(profileName string) Profile {
	var profile Profile
	viper.UnmarshalKey(profileName, &profile)
	return profile
}
