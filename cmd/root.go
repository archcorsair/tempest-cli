/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tempest [command]",
	Short: "tempest-cli interact with your tempest devices",
	Long:  "tempest-cli is a simple CLI tool for getting weather data from your Tempest Weather Devices",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tempest-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tempest-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".tempest-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		apiKey := viper.GetString("api_key")
		stationId := viper.GetString("station_id")
		if apiKey == "" {
			validate := func(input string) error {
				if len(input) != 36 {
					return errors.New("invalid API key\ntry again with format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
				}
				return nil
			}
			prompt := promptui.Prompt{
				Label:    "API Key",
				Validate: validate,
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			viper.Set("api_key", result)
			viper.WriteConfigAs(viper.ConfigFileUsed())
		}
		if stationId == "" {
			validate := func(input string) error {
				if len(input) != 5 {
					return errors.New("invalid station id - must be 5 digits")
				}
				return nil
			}
			prompt := promptui.Prompt{
				Label:    "Station Id",
				Validate: validate,
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			viper.Set("station_id", result)
			viper.WriteConfigAs(viper.ConfigFileUsed())
		}
	}
}
