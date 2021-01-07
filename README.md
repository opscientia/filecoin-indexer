# filecoin-indexer

Data indexer and API service for the Filecoin network

## Requirements

- Go 1.14+
- PostgreSQL 10+

## Installation

```bash
$ go get -u github.com/figment-networks/filecoin-indexer
```

## Usage

```bash
$ ./filecoin-indexer --help

Usage of ./filecoin-indexer:
  -cmd string
    	Command to run
  -config string
    	Path to a config file
```

Executing commands:

```bash
$ filecoin-indexer -config path/to/config.json -cmd=COMMAND
```

Available commands:

| Name       | Description                                   |
|------------|-----------------------------------------------|
| `migrate`  | Migrates the database                         |
| `rollback` | Rolls the schema back to the previous version |
| `sync`     | Runs the synchronization process              |
| `server`   | Starts the API server                         |

## Configuration

You can configure the indexer using a config file or environment variables.

### Config file

Example:

```json
{
  "rpc_endpoint": "5.6.7.8:1234",
  "database_dsn": "dbname=filecoin-indexer",
  "server_addr": "localhost",
  "server_port": 8080,
  "initial_height": 0,
  "batch_size": 100,
  "debug": true
}
```

### Environment variables

| Name             | Description               | Default Value | Required |
|------------------|---------------------------|---------------|----------|
| `RPC_ENDPOINT`   | Lotus RPC endpoint        | —             | Yes      |
| `DATABASE_DSN`   | PostgreSQL database URL   | —             | Yes      |
| `SERVER_ADDR`    | HTTP server address       | `0.0.0.0`     | No       |
| `SERVER_PORT`    | HTTP server port          | `8080`        | No       |
| `INITIAL_HEIGHT` | Initial sync height       | `0`           | No       |
| `BATCH_SIZE`     | Limit of heights per sync | —             | No       |
| `DEBUG`          | Debug mode                | `false`       | No       |

## Running Application

Once you have created a database and specified all configuration options, you need to migrate the database:

```bash
$ filecoin-indexer -config config.json -cmd=migrate
```

The synchronization process can be initiated with the command below:

```bash
$ filecoin-indexer -config config.json -cmd=sync
```

To start the API server, you have to run the following command:

```bash
$ filecoin-indexer -config config.json -cmd=server
```

## API Reference

| Method | Path                              | Description                    | Parameters          |
|--------|-----------------------------------|--------------------------------|---------------------|
| GET    | `/miners`                         | List of all storage miners     | `height` (optional) |
| GET    | `/miners/:address`                | Storage miner details          | `height` (optional) |
| GET    | `/miners/:address/events`         | List of storage miner events   | `height` (optional) |
| GET    | `/top_miners`                     | List of top 100 storage miners | `height` (optional) |
| GET    | `/transactions`                   | List of all transactions       | `height` (optional) |
| GET    | `/accounts/:address`              | Account details                | —                   |
| GET    | `/accounts/:address/transactions` | List of account transactions   | `height` (optional) |
| GET    | `/events`                         | List of all events             | `height` (optional) |

## Score Calculation

The reputation of storage miners is calculated with the following formula:

![Score formula](assets/score-formula.svg)

Where:

![Symbol description](assets/symbol-description.svg)

### Variables

| Name        | Description                                                      | Weight |
|-------------|------------------------------------------------------------------|--------|
| Slashes     | Reciprocal of the number of miner's deals that have been slashed | 100    |
| Faults      | Reciprocal of the total number of miner's faults                 | 100    |
| Power       | Miner's quality-adjusted power divided by network power          | 100    |
| Sector Size | Miner's sector size divided by 32 GiB                            | 10     |

## License

The application is dual-licensed under:

- MIT License
- Apache License 2.0
