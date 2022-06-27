package types

type DeviceObservation struct {
	Status struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"status"`
	DeviceID int    `json:"device_id"`
	Type     string `json:"type"`
	Source   string `json:"source"`
	Summary  struct {
		PressureTrend                  string  `json:"pressure_trend"`
		StrikeCount1H                  int     `json:"strike_count_1h"`
		StrikeCount3H                  int     `json:"strike_count_3h"`
		PrecipTotal1H                  float64 `json:"precip_total_1h"`
		PrecipAccumLocalYesterday      float64 `json:"precip_accum_local_yesterday"`
		PrecipAccumLocalYesterdayFinal float64 `json:"precip_accum_local_yesterday_final"`
		PrecipAnalysisTypeYesterday    int     `json:"precip_analysis_type_yesterday"`
		FeelsLike                      float64 `json:"feels_like"`
		HeatIndex                      float64 `json:"heat_index"`
		WindChill                      float64 `json:"wind_chill"`
	} `json:"summary"`
	Obs [][]float64 `json:"obs"`
}
