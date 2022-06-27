package types

type StationMetadata struct {
	Stations []struct {
		LocationID            int     `json:"location_id"`
		StationID             int     `json:"station_id"`
		Name                  string  `json:"name"`
		PublicName            string  `json:"public_name"`
		Latitude              float64 `json:"latitude"`
		Longitude             float64 `json:"longitude"`
		Timezone              string  `json:"timezone"`
		TimezoneOffsetMinutes int     `json:"timezone_offset_minutes"`
		StationMeta           struct {
			ShareWithWf bool    `json:"share_with_wf"`
			ShareWithWu bool    `json:"share_with_wu"`
			Elevation   float64 `json:"elevation"`
		} `json:"station_meta"`
		LastModifiedEpoch int `json:"last_modified_epoch"`
		CreatedEpoch      int `json:"created_epoch"`
		Devices           []struct {
			DeviceID     int    `json:"device_id"`
			SerialNumber string `json:"serial_number"`
			LocationID   int    `json:"location_id"`
			DeviceMeta   struct {
				Agl             float64 `json:"agl"`
				Name            string  `json:"name"`
				Environment     string  `json:"environment"`
				WifiNetworkName string  `json:"wifi_network_name"`
			} `json:"device_meta"`
			DeviceType       string `json:"device_type"`
			HardwareRevision string `json:"hardware_revision"`
			FirmwareRevision string `json:"firmware_revision"`
			DeviceSettings   struct {
				ShowPrecipFinal bool `json:"show_precip_final"`
			} `json:"device_settings,omitempty"`
		} `json:"devices"`
		StationItems []struct {
			LocationItemID int    `json:"location_item_id"`
			LocationID     int    `json:"location_id"`
			DeviceID       int    `json:"device_id"`
			Item           string `json:"item"`
			Sort           int    `json:"sort"`
			StationID      int    `json:"station_id"`
			StationItemID  int    `json:"station_item_id"`
		} `json:"station_items"`
		IsLocalMode bool `json:"is_local_mode"`
	} `json:"stations"`
	Status struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"status"`
}
