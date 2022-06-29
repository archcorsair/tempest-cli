/*
Copyright © 2022 Daniel Shneyder <archcorsair@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tempest-cli/types"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const BASEURL = "https://swd.weatherflow.com/swd/rest"

var (
	stationId string = ""
	scale     string = ""
)

var forecastCmd = &cobra.Command{
	Use:   "forecast [station id]",
	Short: "Get forecast from weather station",
	Long:  "Get forecast from default weather station or station with [station id]. Will use default station if no station id is provided",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stationId = viper.GetString("station_id")
		scale = viper.GetString("scale")
		if scale == "" {
			scale = "F"
		}

		body, err := BasicFetch(BASEURL + "/better_forecast?station_id=" + stationId + "&token=" + viper.GetString("api_key"))
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
		fmt.Printf("%s\n", formatTime(currentConditions.Time, ""))
		fmt.Printf("--------------------------\nCurrent Conditions: %s %s\n--------------------------\n", getForecastIcon(currentConditions.Icon), currentConditions.Conditions)
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
		// TODO: hourly := result.Forecast.Hourly

		fmt.Printf("--------------------------\nForecast\n--------------------------\n")
		for i := 0; i < len(daily); i++ {
			fmt.Println(renderDailyForecast(daily[i]))
		}

	},
}

func renderDailyForecast(daily types.Daily) string {
	precipProbability := strconv.Itoa(daily.PrecipProbability)
	if precipProbability == "0" {
		precipProbability = ""
	}

	dailyForecast := fmt.Sprint(formatTime(daily.DayStartLocal, "forecast"), "\n")
	dailyForecast += fmt.Sprint("🌡️  High ", ConvertToScale(daily.AirTempHigh), "°", scale, " -> Low ", ConvertToScale(daily.AirTempLow), "°", scale, "\n")
	dailyForecast += fmt.Sprint(getForecastIcon(daily.Icon), " ", daily.Conditions, "\n")
	dailyForecast += fmt.Sprint("🌧️  ", precipProbability, "%\n")
	dailyForecast += fmt.Sprint("🌅 ", formatTime(daily.Sunrise, "sun"), "\n")
	dailyForecast += fmt.Sprint("🌇 ", formatTime(daily.Sunset, "sun"), "\n")

	return dailyForecast
}

func getForecastIcon(iconString string) string {
	switch iconString {
	case "possibly-rainy-day":
		return "☔"
	case "partly-cloudy-day":
		return "⛅"
	case "clear-day":
		return "☀️ "
	default:
		return ""
	}
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}

func formatTime(unixTime int, format string) string {
	t := time.Unix(int64(unixTime), 0)
	if format == "sun" {
		return t.Format("3:04PM")
	}
	if format == "forecast" {
		return t.Format("Mon, Jan 2")
	}
	return t.Format("Monday Jan 2 03:04:05PM 2006")
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
