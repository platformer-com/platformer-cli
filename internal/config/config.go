package config

import (
	"strings"

	"github.com/spf13/viper"
)

// SaveToken saves the given permanent token to the local config
func SaveToken(token string) {
	viper.Set("auth.token", strings.TrimSpace(token))
	_ = viper.WriteConfig()
}

// RemoveToken removes the locally saved token
func RemoveToken() {
	viper.Set("auth.token", "")
	_ = viper.WriteConfig()
}

// GetToken retrieves the locally stored perm.token
// Returns an empty string if the value is not set.
func GetToken() string {
	return viper.GetString("auth.token")
}

// SetDefaultOrg saves a given organization name to the local config
func SetDefaultOrg(orgName string) {
	viper.Set("context.organization", strings.TrimSpace(orgName))
	_ = viper.WriteConfig()
}

// GetDefaultOrg retrieves the saved default org name from the local config.
// Returns an empty string if the value is not set.
func GetDefaultOrg() string {
	return viper.GetString("context.organization")
}

// SetDefaultProject saves a given project name to the local config
func SetDefaultProject(projectName string) {
	viper.Set("context.project", strings.TrimSpace(projectName))
	_ = viper.WriteConfig()
}

// GetDefaultProject retrieves the saved default project name from the local config.
// Returns an empty string if the value is not set.
func GetDefaultProject() string {
	return viper.GetString("context.project")
}

func GetDefaultContextURL() string {
	defaultContextURL := "https://beta.api.platformer.com"

	if viper.GetString("context.url") != "" {
		return viper.GetString("context.url")
	}

	return defaultContextURL

}

func SetDefaultContextURL(url string) {
	viper.Set("context.url", strings.TrimSpace(url))
}
