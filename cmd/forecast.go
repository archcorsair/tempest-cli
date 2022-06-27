/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tempest-cli/types"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const BASEURL = "https://swd.weatherflow.com/swd/rest"

var (
	stationId string = ""
	scale     string = "F"
)

var forecastCmd = &cobra.Command{
	Use:   "forecast [station id]",
	Short: "Get forecast from weather station",
	Long:  "Get forecast from default weather station or station with [station id]. Will use default station if no station id is provided",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stationId = viper.GetString("station_id")
		if len(args) == 0 && stationId == "" {
			input := ""
			fmt.Println("Please provide a station id:")
			fmt.Scanln(&input)
			stationId = input
		}

		targetStation := ""
		if len(args) == 1 {
			targetStation = args[0]
			stationId = targetStation
		} else {
			targetStation = stationId
		}

		body, err := BasicFetch(BASEURL + "/better_forecast?station_id=" + targetStation + "&token=" + viper.GetString("api_key"))
		if err != nil {
			fmt.Println(err)
		}

		var result types.Forecast
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Cannot unmarshal JSON")
		}

		currentConditions := result.CurrentConditions
		units := result.Units

		// Current Conditions
		fmt.Println("--------------------------")
		fmt.Printf("Station ID: %s @ %s \n", stationId, result.LocationName)
		fmt.Printf("%s\n", UTCTime(currentConditions.Time))
		fmt.Printf("--------------------------\nCurrent Conditions: %s\n--------------------------\n", currentConditions.Conditions)
		fmt.Printf("Temp: %.2f°%s\n", ConvertToScale(currentConditions.AirTemperature), scale)
		fmt.Printf("Feels Like: %.2f°%s\n", ConvertToScale(currentConditions.FeelsLike), scale)
		fmt.Printf("Rel Humidity: %d%%\n", currentConditions.RelativeHumidity)
		fmt.Printf("Dew Point: %.2f°%s\n", ConvertToScale(currentConditions.DewPoint), scale)
		fmt.Printf("Avg Wind Speed: %.2f %s\n", currentConditions.WindAvg, units.UnitsWind)
		fmt.Printf("Wind Direction: %s\n", currentConditions.WindDirectionCardinal)
		fmt.Printf("Wind Gust: %.2f %s\n", currentConditions.WindGust, units.UnitsWind)
		fmt.Printf("Pressure: %.2f %s\n", currentConditions.StationPressure, units.UnitsPressure)
		fmt.Printf("Pressure Trend: %s\n", currentConditions.PressureTrend)
		fmt.Printf("Solar Radiation: %d %s\n", currentConditions.SolarRadiation, units.UnitsSolarRadiation)
		fmt.Printf("UV Index: %d\n", currentConditions.Uv)
		fmt.Printf("Brightness: %d %s\n", currentConditions.Brightness, units.UnitsBrightness)

		// Forecast
		daily := result.Forecast.Daily
		fmt.Printf("DAILY TYPE: %T", daily)

		fmt.Printf("--------------------------\nForecast\n--------------------------\n")
		fmt.Printf("%d/%d - %s\n", daily[0].MonthNum, daily[0].DayNum, daily[0].Conditions)
		// fmt.Printf("Rain Chance: %s %d%%\n", daily[0].PrecipProbability, getRainChanceIcon(0))
		// TODO: Render this with a for loop over the daily array

	},
}

func renderDailyForecast(daily) {
	for i := 0; i < len(daily); i++ {
		fmt.Printf(daily[i])
	}
}

// func getRainChanceIcon(index int, daily struct) string {
// 	icon := daily[index].PrecipIcon
// 	switch icon {
// 	case "chance-rain":
// 		return "☔"
// 	default:
// 		return ""
// 	}
// }

func init() {
	rootCmd.AddCommand(forecastCmd)
}

// type promptContext struct {
// 	errorMsg string
// 	label    string
// }

// func promptGetInput(pc promptContext) string {
// 	validate := func(input string) error {
// 		if len(input) <= 0 {
// 			return errors.New(pc.errorMsg)
// 		}
// 		return nil
// 	}

// 	templates := &promptui.PromptTemplates{
// 		Prompt:  "{{ . }} ",
// 		Valid:   "{{ . | green }} ",
// 		Invalid: "{{ . | red }} ",
// 		Success: "{{ . | bold }} ",
// 	}

// 	prompt := promptui.Prompt{
// 		Label:     pc.label,
// 		Templates: templates,
// 		Validate:  validate,
// 	}

// 	result, err := prompt.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Input: %s\n", result)
// 	return result
// }

func UTCTime(unixTime int) string {
	t := time.Unix(int64(unixTime), 0)
	return t.Format("Mon Jan 2 03:04:05PM")
}

func BasicFetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, nil
}

func ConvertToScale(c float64) float64 {
	if scale == "F" {
		return c*9/5 + 32
	}
	return c
}
