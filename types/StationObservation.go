package types

type StationObservation struct {
	StationID   int     `json:"station_id"`
	StationName string  `json:"station_name"`
	PublicName  string  `json:"public_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Timezone    string  `json:"timezone"`
	Elevation   float64 `json:"elevation"`
	IsPublic    bool    `json:"is_public"`
	Status      struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"status"`
	StationUnits struct {
		UnitsTemp      string `json:"units_temp"`
		UnitsWind      string `json:"units_wind"`
		UnitsPrecip    string `json:"units_precip"`
		UnitsPressure  string `json:"units_pressure"`
		UnitsDistance  string `json:"units_distance"`
		UnitsDirection string `json:"units_direction"`
		UnitsOther     string `json:"units_other"`
	} `json:"station_units"`
	OutdoorKeys []string `json:"outdoor_keys"`
	Obs         []struct {
		Timestamp                        int     `json:"timestamp"`
		AirTemperature                   float64 `json:"air_temperature"`
		BarometricPressure               float64 `json:"barometric_pressure"`
		StationPressure                  float64 `json:"station_pressure"`
		SeaLevelPressure                 float64 `json:"sea_level_pressure"`
		RelativeHumidity                 int     `json:"relative_humidity"`
		Precip                           float64 `json:"precip"`
		PrecipAccumLast1Hr               float64 `json:"precip_accum_last_1hr"`
		PrecipAccumLocalDay              float64 `json:"precip_accum_local_day"`
		PrecipAccumLocalDayFinal         float64 `json:"precip_accum_local_day_final"`
		PrecipAccumLocalYesterday        float64 `json:"precip_accum_local_yesterday"`
		PrecipAccumLocalYesterdayFinal   float64 `json:"precip_accum_local_yesterday_final"`
		PrecipMinutesLocalDay            int     `json:"precip_minutes_local_day"`
		PrecipMinutesLocalYesterday      int     `json:"precip_minutes_local_yesterday"`
		PrecipMinutesLocalYesterdayFinal int     `json:"precip_minutes_local_yesterday_final"`
		PrecipAnalysisTypeYesterday      int     `json:"precip_analysis_type_yesterday"`
		WindAvg                          float64 `json:"wind_avg"`
		WindDirection                    int     `json:"wind_direction"`
		WindGust                         float64 `json:"wind_gust"`
		WindLull                         float64 `json:"wind_lull"`
		SolarRadiation                   int     `json:"solar_radiation"`
		Uv                               float64 `json:"uv"`
		Brightness                       int     `json:"brightness"`
		LightningStrikeCount             int     `json:"lightning_strike_count"`
		LightningStrikeCountLast1Hr      int     `json:"lightning_strike_count_last_1hr"`
		LightningStrikeCountLast3Hr      int     `json:"lightning_strike_count_last_3hr"`
		FeelsLike                        float64 `json:"feels_like"`
		HeatIndex                        float64 `json:"heat_index"`
		WindChill                        float64 `json:"wind_chill"`
		DewPoint                         float64 `json:"dew_point"`
		WetBulbTemperature               float64 `json:"wet_bulb_temperature"`
		WetBulbGlobeTemperature          float64 `json:"wet_bulb_globe_temperature"`
		DeltaT                           float64 `json:"delta_t"`
		AirDensity                       float64 `json:"air_density"`
		PressureTrend                    string  `json:"pressure_trend"`
	} `json:"obs"`
}
