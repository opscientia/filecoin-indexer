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
  -mode string
    	Fetcher mode (default "worker")
```

Executing commands:

```bash
$ filecoin-indexer -config path/to/config.json -cmd COMMAND -mode MODE
```

Available commands:

| Name       | Description                                   |
|------------|-----------------------------------------------|
| `migrate`  | Migrates the database                         |
| `rollback` | Rolls the schema back to the previous version |
| `fetcher`  | Starts the fetching process                   |
| `indexer`  | Starts the indexing process                   |
| `server`   | Starts the API server                         |

## Configuration

You can configure the indexer using a config file or environment variables.

### Config file

Example:

```json
{
  "rpc_endpoint": "5.6.7.8:1234",
  "rpc_timeout": "30s",
  "database_dsn": "dbname=filecoin-indexer",
  "initial_height": 0,
  "batch_size": 100,
  "sync_interval": "1s",
  "worker_addr": "localhost",
  "worker_port": "7000",
  "workers": "localhost:7000 localhost:7001",
  "server_addr": "0.0.0.0",
  "server_port": 8080,
  "rollbar_token": "a672d00ae7967e33e6f07a9cbdbefb2a",
  "rollbar_env": "staging",
  "metrics_addr": "localhost",
  "metrics_port": "8090",
  "debug": true
}
```

### Environment variables

| Name                  | Description                      | Default Value    | Required |
|-----------------------|----------------------------------|------------------|----------|
| `RPC_ENDPOINT`        | Lotus RPC endpoint               | —                | Yes      |
| `RPC_TIMEOUT`         | RPC client timeout               | `30s`            | No       |
| `DATABASE_DSN`        | PostgreSQL database URL          | —                | Yes      |
| `INITIAL_HEIGHT`      | Initial sync height              | `0`              | No       |
| `BATCH_SIZE`          | Number of heights per sync       | —                | No       |
| `SYNC_INTERVAL`       | Interval between sync jobs       | `1s`             | No       |
| `WORKER_ADDR`         | Worker server address            | `127.0.0.1`      | No       |
| `WORKER_PORT`         | Worker server port               | `7000`           | No       |
| `WORKERS`             | Space-separated worker endpoints | `127.0.0.1:7000` | No       |
| `SERVER_ADDR`         | API server address               | `0.0.0.0`        | No       |
| `SERVER_PORT`         | API server port                  | `8080`           | No       |
| `ROLLBAR_TOKEN`       | Rollbar token                    | —                | No       |
| `ROLLBAR_ENV`         | Rollbar environment              | `development`    | No       |
| `METRICS_ADDR`        | Metrics server address           | `127.0.0.1`      | No       |
| `METRICS_PORT`        | Metrics server port              | `8090`           | No       |
| `DEBUG`               | Debug mode                       | `false`          | No       |

## Running Application

Once you have created a database and specified all configuration options, you need to migrate the database:

```bash
$ filecoin-indexer -config config.json -cmd migrate
```

The indexing process can be started with the command below:

```bash
$ filecoin-indexer -config config.json -cmd indexer
```

To start the API server, you have to run the following command:

```bash
$ filecoin-indexer -config config.json -cmd server
```

## API Reference

| Method | Path                              | Description                    |
|--------|-----------------------------------|--------------------------------|
| GET    | `/miners`                         | List of all storage miners     |
| GET    | `/miners/:address`                | Storage miner details          |
| GET    | `/miners/:address/events`         | List of storage miner events   |
| GET    | `/top_miners`                     | List of top 100 storage miners |
| GET    | `/transactions`                   | List of all transactions       |
| GET    | `/accounts/:address`              | Account details                |
| GET    | `/accounts/:address/transactions` | List of account transactions   |
| GET    | `/events`                         | List of all events             |
| GET    | `/health`                         | Health check                   |
| GET    | `/status`                         | Synchronization status         |
| GET    | `/metrics`                        | Prometheus metrics             |

For more information see the [Miner Reputation System API](https://learn.figment.io/network-documentation/filecoin/rpc-and-rest-api/miner-reputation-system-api) on Figment Learn.

## Score Calculation

The reputation of storage miners is calculated with the following formula:

![Score formula](assets/score-formula.svg)

Where:

![Symbol description](assets/symbol-description.svg)

### Variables

| Name        | Description                                       | Weight |
|-------------|---------------------------------------------------|--------|
| Slashes     | Reciprocal of the number of slashed deals squared | 100    |
| Faults      | Reciprocal of the total number of faults          | 100    |
| Power       | Quality-adjusted power divided by network power   | 100    |
| Sector Size | Sector size divided by 32 GiB                     | 10     |

## License

This application is dual-licensed under:

- MIT License
- Apache License 2.0
