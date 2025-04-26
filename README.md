# GORM SQL Server Example

This project demonstrates how to use GORM with SQL Server in Go.

## Project Structure

- `cmd/` - Contains the main application
- `sqlwrap/` - SQL library for interacting with the database
- `docker-compose.yml` - Docker Compose configuration for SQL Server and the Go application
- `init-db.sql` - SQL script to initialize the database

## Running the Application

### Using Docker Compose

1. Make sure you have Docker and Docker Compose installed
2. Run the following command:

```bash
docker-compose up
```

This will:
- Start a SQL Server container
- Initialize the database with sample data
- Build and start the Go application

### Running Locally

1. Start SQL Server (you can use Docker):

```bash
docker-compose up db
```

2. Run the Go application:

```bash
cd cmd
go run main.go
```

## Connection String

The default connection string is:

```
sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=testdb&encrypt=true&trustServerCertificate=true
```

You can override it by setting the `DB_CONNECTION_STRING` environment variable. 