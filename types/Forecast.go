package types

type Forecast struct {
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	Timezone              string  `json:"timezone"`
	TimezoneOffsetMinutes int     `json:"timezone_offset_minutes"`
	LocationName          string  `json:"location_name"`
	CurrentConditions     struct {
		Time                            int     `json:"time"`
		Conditions                      string  `json:"conditions"`
		Icon                            string  `json:"icon"`
		AirTemperature                  float64 `json:"air_temperature"`
		SeaLevelPressure                float64 `json:"sea_level_pressure"`
		StationPressure                 float64 `json:"station_pressure"`
		PressureTrend                   string  `json:"pressure_trend"`
		RelativeHumidity                int     `json:"relative_humidity"`
		WindAvg                         float64 `json:"wind_avg"`
		WindDirection                   int     `json:"wind_direction"`
		WindDirectionCardinal           string  `json:"wind_direction_cardinal"`
		WindGust                        float64 `json:"wind_gust"`
		SolarRadiation                  int     `json:"solar_radiation"`
		Uv                              int     `json:"uv"`
		Brightness                      int     `json:"brightness"`
		FeelsLike                       float64 `json:"feels_like"`
		DewPoint                        float64 `json:"dew_point"`
		WetBulbTemperature              float64 `json:"wet_bulb_temperature"`
		WetBulbGlobeTemperature         float64 `json:"wet_bulb_globe_temperature"`
		DeltaT                          float64 `json:"delta_t"`
		AirDensity                      float64 `json:"air_density"`
		LightningStrikeCountLast1Hr     int     `json:"lightning_strike_count_last_1hr"`
		LightningStrikeCountLast3Hr     int     `json:"lightning_strike_count_last_3hr"`
		PrecipAccumLocalDay             int     `json:"precip_accum_local_day"`
		PrecipAccumLocalYesterday       int     `json:"precip_accum_local_yesterday"`
		PrecipMinutesLocalDay           int     `json:"precip_minutes_local_day"`
		PrecipMinutesLocalYesterday     int     `json:"precip_minutes_local_yesterday"`
		IsPrecipLocalDayRainCheck       bool    `json:"is_precip_local_day_rain_check"`
		IsPrecipLocalYesterdayRainCheck bool    `json:"is_precip_local_yesterday_rain_check"`
	} `json:"current_conditions"`
	Forecast struct {
		Daily []struct {
			DayStartLocal     int     `json:"day_start_local"`
			DayNum            int     `json:"day_num"`
			MonthNum          int     `json:"month_num"`
			Conditions        string  `json:"conditions"`
			Icon              string  `json:"icon"`
			Sunrise           int     `json:"sunrise"`
			Sunset            int     `json:"sunset"`
			AirTempHigh       float64 `json:"air_temp_high"`
			AirTempLow        float64 `json:"air_temp_low"`
			PrecipProbability int     `json:"precip_probability"`
			PrecipIcon        string  `json:"precip_icon"`
			PrecipType        string  `json:"precip_type"`
		} `json:"daily"`
		Hourly []struct {
			Time                  int     `json:"time"`
			Conditions            string  `json:"conditions"`
			Icon                  string  `json:"icon"`
			AirTemperature        float64 `json:"air_temperature"`
			SeaLevelPressure      float64 `json:"sea_level_pressure"`
			RelativeHumidity      int     `json:"relative_humidity"`
			Precip                int     `json:"precip"`
			PrecipProbability     int     `json:"precip_probability"`
			PrecipType            string  `json:"precip_type"`
			PrecipIcon            string  `json:"precip_icon"`
			WindAvg               float64 `json:"wind_avg"`
			WindDirection         int     `json:"wind_direction"`
			WindDirectionCardinal string  `json:"wind_direction_cardinal"`
			WindGust              float64 `json:"wind_gust"`
			Uv                    float64 `json:"uv"`
			FeelsLike             float64 `json:"feels_like"`
			LocalHour             int     `json:"local_hour"`
			LocalDay              int     `json:"local_day"`
		} `json:"hourly"`
	} `json:"forecast"`
	Status struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"status"`
	Units struct {
		UnitsTemp           string `json:"units_temp"`
		UnitsWind           string `json:"units_wind"`
		UnitsPrecip         string `json:"units_precip"`
		UnitsPressure       string `json:"units_pressure"`
		UnitsDistance       string `json:"units_distance"`
		UnitsBrightness     string `json:"units_brightness"`
		UnitsSolarRadiation string `json:"units_solar_radiation"`
		UnitsOther          string `json:"units_other"`
		UnitsAirDensity     string `json:"units_air_density"`
	} `json:"units"`
	SourceIDConditions int `json:"source_id_conditions"`
}

type Daily struct {
	DayStartLocal     int     `json:"day_start_local"`
	DayNum            int     `json:"day_num"`
	MonthNum          int     `json:"month_num"`
	Conditions        string  `json:"conditions"`
	Icon              string  `json:"icon"`
	Sunrise           int     `json:"sunrise"`
	Sunset            int     `json:"sunset"`
	AirTempHigh       float64 `json:"air_temp_high"`
	AirTempLow        float64 `json:"air_temp_low"`
	PrecipProbability int     `json:"precip_probability"`
	PrecipIcon        string  `json:"precip_icon"`
	PrecipType        string  `json:"precip_type"`
}

type Hourly struct {
	Time                  int     `json:"time"`
	Conditions            string  `json:"conditions"`
	Icon                  string  `json:"icon"`
	AirTemperature        float64 `json:"air_temperature"`
	SeaLevelPressure      float64 `json:"sea_level_pressure"`
	RelativeHumidity      int     `json:"relative_humidity"`
	Precip                int     `json:"precip"`
	PrecipProbability     int     `json:"precip_probability"`
	PrecipType            string  `json:"precip_type"`
	PrecipIcon            string  `json:"precip_icon"`
	WindAvg               float64 `json:"wind_avg"`
	WindDirection         int     `json:"wind_direction"`
	WindDirectionCardinal string  `json:"wind_direction_cardinal"`
	WindGust              float64 `json:"wind_gust"`
	Uv                    float64 `json:"uv"`
	FeelsLike             float64 `json:"feels_like"`
	LocalHour             int     `json:"local_hour"`
	LocalDay              int     `json:"local_day"`
}
