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

**Sample Usage**:

`tempest forecast -c false -t`

Omit current conditions and show today's forecast only

`tempest forecast -m 5 -s 12345`

Show current conditions from station id 12345 and a 5 day forecast (including today)

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
