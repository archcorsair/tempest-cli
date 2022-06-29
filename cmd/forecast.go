/*
Copyright © 2022 Daniel Shneyder <archcorsair@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tempest-cli/types"
	"tempest-cli/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const BASEURL = "https://swd.weatherflow.com/swd/rest"

var (
	stationId      int    = 0
	scale          string = ""
	maxDays        int    = 10
	todayOnly      bool   = false
	showConditions bool   = true
)

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().IntVarP(&maxDays, "max", "m", 10, "Maximum days to display up to 10")
	forecastCmd.Flags().BoolVarP(&todayOnly, "today", "t", false, "Display only today's forecast")
	forecastCmd.Flags().IntVarP(&stationId, "station", "s", 0, "Display forecast for a specific owned station")
	forecastCmd.Flags().BoolVarP(&showConditions, "conditions", "c", true, "Whether to display conditions")
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get forecast from weather station",
	Long:  "Get forecast from default weather station or station with [station id]. Will use default station if no station id is provided",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scale = viper.GetString("scale")
		apiKey := viper.GetString("api_key")

		if stationId == 0 {
			stationId = viper.GetInt("station_id")
		}

		if scale == "" {
			scale = "F"
		}

		body, err := util.Fetch(BASEURL + "/better_forecast?station_id=" + strconv.Itoa(stationId) + "&token=" + apiKey)
		if err != nil {
			fmt.Println(err)
		}

		var result types.Forecast
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Cannot unmarshal JSON")
		}
		// Current Conditions
		fmt.Println(renderConditions(result, showConditions))
		// Forecast
		daily := result.Forecast.Daily
		// TODO: hourly := result.Forecast.Hourly
		fmt.Printf("--------------------------\nForecast\n--------------------------\n")
		var daysToDisplay int
		if maxDays > len(daily) {
			daysToDisplay = len(daily)
		} else if maxDays < len(daily) || maxDays <= 0 {
			daysToDisplay = maxDays
		} else if todayOnly {
			daysToDisplay = 1
		} else {
			daysToDisplay = maxDays
		}

		for i := 0; i < daysToDisplay; i++ {
			fmt.Println(renderDailyForecast(daily[i]))
		}
	},
}

func renderConditions(result types.Forecast, show bool) string {
	if !show {
		return ""
	}
	currentConditions := result.CurrentConditions
	units := result.Units

	conditionsString := "--------------------------"
	conditionsString += fmt.Sprintf("Station ID: %d @ %s \n", stationId, result.LocationName)
	conditionsString += fmt.Sprintf("%s\n", util.FormatTime(currentConditions.Time, ""))
	conditionsString += fmt.Sprintf("--------------------------\nCurrent Conditions: %s %s\n--------------------------\n", util.GetWeatherIcon(currentConditions.Icon), currentConditions.Conditions)
	conditionsString += fmt.Sprintf("Temp: %.2f°%s\n", util.CurrentScale(currentConditions.AirTemperature, scale), scale)
	conditionsString += fmt.Sprintf("Feels Like: %.2f°%s\n", util.CurrentScale(currentConditions.FeelsLike, scale), scale)
	conditionsString += fmt.Sprintf("Rel Humidity: %d%%\n", currentConditions.RelativeHumidity)
	conditionsString += fmt.Sprintf("Dew Point: %.2f°%s\n", util.CurrentScale(currentConditions.DewPoint, scale), scale)
	conditionsString += fmt.Sprintf("Avg Wind Speed: %.2f %s\n", currentConditions.WindAvg, units.UnitsWind)
	conditionsString += fmt.Sprintf("Wind Direction: %s\n", currentConditions.WindDirectionCardinal)
	conditionsString += fmt.Sprintf("Wind Gust: %.2f %s\n", currentConditions.WindGust, units.UnitsWind)
	conditionsString += fmt.Sprintf("Pressure: %.2f %s\n", currentConditions.StationPressure, units.UnitsPressure)
	conditionsString += fmt.Sprintf("Pressure Trend: %s\n", currentConditions.PressureTrend)
	conditionsString += fmt.Sprintf("Solar Radiation: %d %s\n", currentConditions.SolarRadiation, units.UnitsSolarRadiation)
	conditionsString += fmt.Sprintf("UV Index: %d\n", currentConditions.Uv)
	conditionsString += fmt.Sprintf("Brightness: %d %s\n", currentConditions.Brightness, units.UnitsBrightness)

	return conditionsString
}

func renderDailyForecast(daily types.Daily) string {
	precipProbability := strconv.Itoa(daily.PrecipProbability)
	if precipProbability == "0" {
		precipProbability = ""
	}

	dailyForecast := fmt.Sprint(util.FormatTime(daily.DayStartLocal, "forecast"), "\n")
	dailyForecast += fmt.Sprint("🌡️  High ", util.CurrentScale(daily.AirTempHigh, scale), "°", scale, " -> Low ", util.CurrentScale(daily.AirTempLow, scale), "°", scale, "\n")
	dailyForecast += fmt.Sprint(util.GetWeatherIcon(daily.Icon), " ", daily.Conditions, "\n")
	dailyForecast += fmt.Sprint("🌧️  ", precipProbability, "%\n")
	dailyForecast += fmt.Sprint("🌅 ", util.FormatTime(daily.Sunrise, "sun"), "\n")
	dailyForecast += fmt.Sprint("🌇 ", util.FormatTime(daily.Sunset, "sun"), "\n")

	return dailyForecast
}
