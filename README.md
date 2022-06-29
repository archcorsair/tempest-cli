# tempest-cli

A simple CLI utility for interacting with your Tempest Weather devices

# Getting Started

### Prerequisites

- Go: https://go.dev/

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
