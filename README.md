# tempest-cli

A simple CLI utility for interacting with your Tempest Weather devices

# Getting Started

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
// MacOS
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
