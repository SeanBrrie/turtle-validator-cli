# Turtle Validator CLI

A command-line tool for validating dataset metadata files against European metadata standards. This tool helps data stewards ensure their dataset descriptions comply with agreed-upon metadata models before submission to national and European data catalogues.

## What This Tool Does

This tool checks if your dataset description files are correctly formatted and contain all required information according to European standards. It's like a spell-checker, but for dataset metadata.

## Prerequisites

- Go 1.24 or later installed on your system

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/SeanBrrie/turtle-validator-cli.git
   cd turtle-validator-cli
   ```

2. Install dependencies:
   ```bash
   go mod init github.com/SeanBrrie/turtle-validator-cli.git
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build ./cmd/turtle-validator-cli/main.go
   ```

## Usage

### Running the Tool

Running the program directly with Go:

```bash
go run cmd/turtle-validator-cli/main.go
```

### CLI

The tool will ask you for the following information:

1. **Domain**: The type of metadata standard you want to validate against
    - `dcat-ap` - For general European data catalogue standards
    - `healthri` - For health research data standards

2. **Content file name**: The name of your metadata file (should be in the `data/` folder)
    - Example: `valid_dcat_ap.ttl`

3. **Context syntax**: The format of your metadata file
    - `Turtle` - For .ttl files
    - `XML` - For .xml files
    - `JSONLD` - For .json files

4. **Validation type**: The specific version of the standard validation type (differs per domain)
    - `V3Full1`
    - `V200`

### Example Session

```
--- New Validation Request ---
Type 'exit' to quit.
Domain: (e.g., dcat-ap, healthri): dcat-ap
Content file name: valid_dcat_ap.ttl
Context syntax (e.g., XML, JSONLD, Turtle): turtle
Validation type (e.g., V3Full1, V200): v3full1
Validation result: true
```

## Test Files

The `data/` folder contains example files you can use to test the validator:

- `valid_dcat_ap.ttl` - A correctly formatted general dataset description
- `valid_health_ri.ttl` - A correctly formatted health research dataset description
- `invalid_dcat_ap.ttl` - An incorrectly formatted general dataset
- `invalid_health_ri.ttl` - An incorrectly formatted health research dataset

## Understanding the Results

- **Validation result: true** - Your metadata file is correctly formatted and complete
- **Validation result: false** - Your metadata file has errors or is missing required information

## Project Structure

```
turtle-validator-cli/
├── cmd/
│   └── turtle-validator-cli/
│       └── main.go                 # Main application entry point
├── data/                           # Example test files
│   ├── valid_dcat_ap.ttl
│   ├── valid_health_ri.ttl
│   ├── invalid_dcat_ap.ttl
│   └── invalid_health_ri.ttl
├── internal/
│   ├── clients/
│   │   ├── enums/                  # Data type definitions
│   │   └── itb_ec_europa_client.go # API client
│   └── services/
│       └── itb_ec_europa_services.go # Validation logic
├── go.mod
├── go.sum
└── README.md
```

## License

N/A
