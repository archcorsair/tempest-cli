# tempest-cli

A simple CLI utility for interacting with your Tempest Weather devices

# Under Construction :construction:

- [x] Scaffold out initial version to allow commands to be added
- [x] Implement [forecast](https://weatherflow.github.io/Tempest/api/swagger/#/forecast) support
- [ ] Implement [Observable station](https://weatherflow.github.io/Tempest/api/swagger/#!/stations/getStations) lookup
- [ ] Implement [Station metadata](https://weatherflow.github.io/Tempest/api/swagger/#!/stations/getStationById) lookup
- [ ] Implement [Device observation](https://weatherflow.github.io/Tempest/api/swagger/#!/observations/getObservationsByDeviceId)
- [ ] Implement [Station observation](https://weatherflow.github.io/Tempest/api/swagger/#!/observations/getStationObservation)
- [ ] Implement interactive persistent dashboard using [termui](https://github.com/gizak/termui)

# Getting Started

## Usage

```
tempest [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  forecast    Get forecast from weather station
  help        Help about any command

Flags:
  -h, --help     help for tempest
  -t, --toggle   Help message for toggle

Use "tempest [command] --help" for more information about a command.
```

### Commands

`forecast`

```
Usage:
  tempest forecast [flags]

Flags:
  -c, --conditions    Whether to display conditions (default true)
  -h, --help          help for forecast
  -m, --max int       Maximum days to display up to 10 (default 10)
  -s, --station int   Display forecast for a specific owned station
  -t, --today         Display only today's forecast
```

**Usage Examples**:

Omit current conditions and show today's forecast only using configured default station id

`tempest forecast -c false -t`

```
--------------------------
Wednesday Jun 01 08:30:20AM 2022
--------------------------
Station ID: 12345 @ Cityville
--------------------------
Current Conditions: ☁️ Cloudy
--------------------------
Temp: 53.60°F
Feels Like: 53.60°F
Rel Humidity: 93%
Dew Point: 51.80°F
Avg Wind Speed: 1.00 mps
Wind Direction: WNW
Wind Gust: 3.00 mps
Pressure: 1008.00 mb
Pressure Trend: rising
Solar Radiation: 149 w/m2
UV Index: 2
Brightness: 17871 lux

--------------------------
10 Day Forecast
--------------------------
Wed, Jun 01
🌡️  High 62.6°F -> Low 53.6°F
🌂 Rain Possible
🌧️  20%
🌅 5:51AM
🌇 8:35PM
```

Show current conditions from station id 12345 and a 2 day forecast (including today)

`tempest forecast -m 2 -s 12345`

```
--------------------------
Wednesday Jun 01 08:57:20AM 2022
--------------------------
Station ID: 12345 @ Cityville
--------------------------
Current Conditions: ⛅️ Partly Cloudy
--------------------------
Temp: 53.60c°F
Feels Like: 53.60°F
Rel Humidity: 91%
Dew Point: 51.80°F
Avg Wind Speed: 1.00 mps
Wind Direction: WNW
Wind Gust: 2.00 mps
Pressure: 1008.20 mb
Pressure Trend: rising
Solar Radiation: 210 w/m2
UV Index: 3
Brightness: 25241 lux

--------------------------
2 Day Forecast
--------------------------
Wed, Jun 01
🌡️  High 62.6°F -> Low 53.6°F
🌂 Rain Possible
🌧️  20%
🌅 5:51AM
🌇 8:35PM

Thu, Jun 30
🌡️  High 62.6°F -> Low 53.6°F
🌂 Rain Possible
🌧️  20%
🌅 5:51AM
🌇 8:35PM
```

### Prerequisites

- Go: https://go.dev/
- A Tempest Weather System: https://weatherflow.com/tempest-weather-system/
- A Tempest API Key (Token): https://tempestwx.com/settings/tokens

### Clone the repository

```
git clone git@github.com:archcorsair/tempest-cli.git
cd tempest-cli/
```

### Download dependencies

```
go mod tidy
```

### Compile binary

```
// MacOS or Linux
go build -o bin/tempest .

// Windows
go build -o bin/tempest.exe .
```

### Running

```
cd bin/

// MacOS
./tempest --help

// Windows
./tempest.exe --help
```

### First Launch

You will be prompted to input your API key and Station ID

These values are required because the tempest API only allows you access to data from your own devices.

### Config File

tempest-cli will generate a config file if none exists: `$HOME/.tempest-cli`
