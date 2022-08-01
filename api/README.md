# Install

### Dependencies
- golang 1.18+
- Postgres (not needed for development)
- SQLite (used to run tests and could be used for development)

[Optional] Postgres: If you do not have the database installed yet, you can use the database as container. The cmd below shows as up the container using docke

``` docker run --name postgres -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=@postgres -e POSTGRES_DB=subscriber_dev -d postgres ```

## Run on development

In order to install all dependencies of project, execute the cmd below.

```go get```

### Environments
Create a new file named ".env" with the content of file .env.example and fill the environments with your correct config:
- sub_database - used to inform the connection string of database. Options: "sqlite | sqlite:memory | postgree_connection". If nothing informed the sqlite will be used
- sub_salt_hash - used to inform the salt to create the hash of password
- sub_jwt_key - used to inform the key to generate the JWT token

# Exec project
In order to execute do the cmd below and the api will be on :6004 port

```air```