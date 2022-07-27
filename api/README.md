# Install

In order to install all dependencies of project, execute the cmd below.

```go get```

### Dependencies
- Postgres

[Optional] Postgres: If you do not have the database installed yet, you can use the database as container. The cmd below shows as up the container using docke

``` docker run --name postgres -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=@postgres -e POSTGRES_DB=subscriber_dev -d postgres ```

### Environments
Create a new file named ".env" with the content of file .env.example and fill the environments with your correct config

# Exec project
In order to execute do the cmd below and the api will be on :6004 port

```air```