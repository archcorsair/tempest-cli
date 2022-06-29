/*
Copyright © 2022 Daniel Shneyder <archcorsair@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homePath string
var configChanged bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tempest [command]",
	Short: "tempest-cli interact with your tempest devices",
	Long:  "tempest-cli is a simple CLI tool for getting weather data from your Tempest Weather Devices",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, error := os.UserHomeDir()
	cobra.CheckErr(error)
	homePath = home

	// Search config in home directory with name ".tempest-cli" (without extension).
	viper.SetConfigName(".tempest-cli")
	viper.SetConfigType("json")
	viper.AddConfigPath(homePath)
	viper.AddConfigPath(".") // Fallback to looking in the current directory
	viper.AutomaticEnv()     // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		// No config file found
		firstLaunch()
	}
	// Validate config file
	if !hasValidApiKey() {
		promptForApiKey()
		configChanged = true
	}

	if !hasValidStationId() {
		promptForStationId()
		configChanged = true
	}

	if configChanged {
		configPath := path.Join(homePath, ".tempest-cli")
		viper.WriteConfigAs(configPath)
		fmt.Println("New config written to: ", configPath)
	}
}

func hasValidApiKey() bool {
	return viper.GetString("api_key") != ""
}

func hasValidStationId() bool {
	return viper.GetString("station_id") != ""
}

func promptForStationId() {
	label := "A valid station ID is required to use Tempest-CLI. Please enter your station ID"
	errorString := "station ID cannot be empty"
	result, err := Prompt(label, errorString)
	if err != nil {
		fmt.Printf("Setting station ID failed %v\n", err)
		os.Exit(1)
	}
	id, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Setting station ID failed %v\n", err)
		os.Exit(1)
	}
	viper.Set("station_id", id)
}

func promptForApiKey() {
	label := "A valid API key is required to use Tempest-CLI. Please enter your API key"
	errorString := "api key cannot be empty"
	fmt.Println("You can find your personal token at https://tempestwx.com/settings/tokens")
	result, err := Prompt(label, errorString)
	if err != nil {
		fmt.Printf("Setting API key failed %v\n", err)
		os.Exit(1)
	}
	viper.Set("api_key", result)
}

func Prompt(label string, errorString string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			if input == "" {
				return errors.New(errorString)
			}
			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func firstLaunch() {
	fmt.Println("Welcome to Tempest-CLI!\nFirst-run Setup")
	fmt.Println("You can find your personal token at https://tempestwx.com/settings/tokens")
	key, err := Prompt("Please enter your API key", "api key cannot be empty")
	if err != nil {
		fmt.Printf("Setting API key failed %v\n", err)
		os.Exit(1)
	}
	viper.Set("api_key", key)

	stationId, err := Prompt("Set your home station ID", "station id cannot be empty")
	if err != nil {
		fmt.Printf("Setting station id failed %v\n", err)
		os.Exit(1)
	}
	viper.Set("station_id", stationId)

	scalePrompt := promptui.Select{
		Label: "Set your preferred temperature scale",
		Items: []string{"Fahrenheit", "Celsius"},
	}
	_, scale, err := scalePrompt.Run()
	if err != nil {
		fmt.Printf("Setting scale failed %v\n", err)
		os.Exit(1)
	}
	var chosenScale string
	if scale == "Fahrenheit" {
		chosenScale = "F"
	} else {
		chosenScale = "C"
	}
	viper.Set("scale", chosenScale)

	configPath := path.Join(homePath, ".tempest-cli")
	viper.SafeWriteConfigAs(configPath)
	fmt.Println("Setup Complete! Saved to: ", configPath)
}
