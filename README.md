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

| Name                    | Description                                   |
|-------------------------|-----------------------------------------------|
| `migrate`, `migrate:up` | Migrates the database                         |
| `migrate:down`          | Rolls the schema back to the previous version |
| `sync`                  | Runs the synchronization process              |
| `server`                | Starts the API server                         |

## Configuration

You can configure the indexer using a config file or environment variables.

### Config File

Example:

```json
{
  "rpc_endpoint": "5.6.7.8:1234",
  "database_dsn": "dbname=filecoin-indexer",
  "server_addr": "localhost",
  "debug": true
}
```

### Environment Variables

| Name                 | Description             | Default Value | Required |
|----------------------|-------------------------|---------------|----------|
| `RPC_ENDPOINT`       | Lotus RPC endpoint      | —             | **Yes**  |
| `DATABASE_DSN`       | PostgreSQL database URL | —             | **Yes**  |
| `SERVER_ADDR`        | HTTP server address     | `0.0.0.0`     | No       |
| `SERVER_PORT`        | HTTP server port        | `8080`        | No       |
| `DEBUG`              | Debug mode              | `false`       | No       |

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

| Method | Path          | Description                    |
|--------|---------------|--------------------------------|
| GET    | `/miners`     | List of all storage miners     |
| GET    | `/top_miners` | List of top 100 storage miners |
