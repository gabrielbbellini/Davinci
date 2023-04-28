# Davinci

## Run project

- Create a file called settings.yml on the root of the project with the following configuration:

```yaml
# server settings
server:
  # your ip address where the server will run
  host: "localhost"

  # port where the server will listen
  port: "8000"

# database settings
database:
  # the database host (domain/ip)
  host: "localhost"

  # the port where the database is listening (usually 3306)
  port: "3306"

  # the database user
  user: "root"

  # the database user's password
  password: "root"

  # the database name
  name: "databasename"
```

- Finally in the root of the project execute the following code:

```lang-none
go run main.go
```

If all the settings where applied successfully the output is: "[main] Server is running on localhost:8000". Enjoy it!
