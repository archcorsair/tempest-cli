package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
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

func FormatTime(unixTime int, format string) string {
	t := time.Unix(int64(unixTime), 0)
	if format == "sun" {
		return t.Format("3:04PM")
	}
	if format == "forecast" {
		return t.Format("Mon, Jan 2")
	}
	return t.Format("Monday Jan 2 03:04:05PM 2006")
}

func CurrentScale(c float64, scale string) float64 {
	if scale == "F" {
		return c*9/5 + 32
	}
	return c
}

func GetWeatherIcon(iconString string) string {
	switch iconString {
	case "clear-day":
		return "☀️"
	case "clear-night":
		return "🌙"
	case "cloudy":
		return "☁️"
	case "foggy":
		return "🌁"
	case "partly-cloudy-day":
		return "⛅️"
	case "partly-cloudy-night":
		return "☁️"
	case "possibly-rainy-day":
		fallthrough
	case "possibly-rainy-night":
		return "🌂"
	case "possibly-sleet-day":
		fallthrough
	case "possibly-sleet-night":
		fallthrough
	case "sleet":
		return "❄️🌧"
	case "possibly-snow-day":
		fallthrough
	case "snow":
		return "🌨"
	case "possibly-snow-night":
		return "🌨"
	case "possibly-thunderstorm-day":
		fallthrough
	case "possibly-thunderstorm-night":
		fallthrough
	case "thunderstorm":
		return "⛈"
	case "rainy":
		return "🌧"
	case "windy":
		return "💨"
	default:
		return ""
	}
}
