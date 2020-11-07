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

| Method | Path          | Description                    | Parameters          |
|--------|---------------|--------------------------------|---------------------|
| GET    | `/miners`     | List of all storage miners     | `height` (optional) |
| GET    | `/top_miners` | List of top 100 storage miners | `height` (optional) |

## Score Calculation

The reputation of storage miners is calculated with the following formula:

<img src="https://latex.codecogs.com/svg.latex?\LARGE%20S_{\mathcal{M}_i}%20=%20\sum_j%20v_{j_{\mathcal{M}_i}}%20w_{v_j}" title="\LARGE S_{\mathcal{M}_i} = \sum_j v_{j_{\mathcal{M}_i}} w_{v_j}">

Where:

<img src="https://latex.codecogs.com/svg.latex?\\%20S_{\mathcal{M}_i}%20-%20\textrm{reputation%20score%20of%20a%20miner%20$\mathcal{M}_i$}%20\\%20\\%20v_{j_{\mathcal{M}_i}}%20-%20\textrm{variable%20$v_j$%20calculated%20for%20the%20miner%20$\mathcal{M}_i$}%20\\%20w_{v_j}%20-%20\textrm{weight%20of%20the%20variable%20$v_j$}" title="\\ S_{\mathcal{M}_i} - \textrm{reputation score of a miner $\mathcal{M}_i$} \\ \\ v_{j_{\mathcal{M}_i}} - \textrm{variable $v_j$ calculated for the miner $\mathcal{M}_i$} \\ w_{v_j} - \textrm{weight of the variable $v_j$}">

### Variables

| Name        | Weight | Description                                             |
|-------------|--------|---------------------------------------------------------|
| Faults      | 100    | Reciprocal of the total number of miner's faults        |
| Power       | 100    | Miner's quality-adjusted power divided by network power |
| Sector Size | 10     | Miner's sector size divided by 32 GiB                   |
