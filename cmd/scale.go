/*
Copyright © 2022 archcorsair <archcorsair@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale [F/C]",
	Short: "Set the default temperature scale",
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if input != "F" && input != "C" {
				return errors.New("scale must only be 'F' or 'C'")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Select a default temp scale [F/C]",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		viper.Set("scale", result)
		viper.WriteConfigAs(viper.ConfigFileUsed())
	},
}

func init() {
	setCmd.AddCommand(scaleCmd)
}
