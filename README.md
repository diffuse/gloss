# gloss
gloss (Golang open simple service) provides boilerplate routing, database setup, and Docker files to get a minimal 
microservice up and running.  It includes example code to increment and retrieve counter values from a PostgreSQL 
database.  Ideally, one would fork, or clone + mirror push this repository, then edit the handlers + routes, database 
queries, and configurations for their own purposes.  It uses [chi](https://github.com/go-chi/chi) for routing, and 
[pgx](https://github.com/jackc/pgx) for its PostgreSQL database driver.

### Prerequisites
- [Docker](https://www.docker.com/)
- [Golang](https://golang.org/) (for a non-Docker build)

### Running 
#### Using Docker
If you have Docker installed, getting the example stack up and running is as simple as:

`./bootstrap_example.sh`

This will build the image locally with a `gloss:latest` tag, then deploy a service stack with a PostgreSQL instance
exposed on port `5432`, and the example gloss service exposed on port `8080` on the host machine.

#### Local build
If you don't want to use Docker stack, you can build and run the example service locally.  

You must first have a PostgreSQL service exposed on your host machine at port `5432`, with the configuration in 
`test.env`, for example (using Docker cli for brevity):
```bash
docker run -d -p 5432:5432 \
-e POSTGRES_PASSWORD="password" \
-e POSTGRES_USER="test" \
-e POSTGRES_DB="test" \
postgres
```

then you can build 
and run the service locally:
```bash
# get the dependencies
go get -d ./...

# install to ~/go/bin
go install ./...

# export the test environment variables
while read l; do export $l; done < test.env

# run the service (this binds to port 8080 and waits for requests)
~/go/bin/gloss
```

### Interacting with the example counter service
If you have deployed the example successfully, you can interact with the service at `http://localhost:8080/v1`

Example interaction:
```bash
# @URL: http://localhost:8080/v1/counter/{counterId}
# - POST: to increment/create a counter in the database
# - GET: to get the current counter value, if it exists

# create a counter with ID 4
curl -v http://localhost:8080/v1/counter/4 -X POST

# a counter has now been created with ID 4 and value 0

# increment it a few times
for i in $(seq 1 10); do curl -v http://localhost:8080/v1/counter/4 -X POST; done

# get the counter value
curl -v http://localhost:8080/v1/counter/4
# returns 10

# connect to the database and run some other queries
psql -h localhost -p 5432 -U test -d test
```

### Adapting to your own implementation
- Add your business-logic methods to the `gloss.Database` interface, then implement it in a custom package
    - (e.g. the example counter `IncrementCounter` and `GetCounterVal` methods, implemented in the pgsql package)
- Edit the pgsql package if you want to use PostgreSQL, or create your own subpackage with a different database,
then implement the `gloss.Database` interface methods in this package
- Use your `gloss.Database` implementation by assigning an instance of your package's Database to the `Db` var in 
`handlers.go`
- Replace/change the example handler bodies in `handlers.go` to perform your business logic
- Update the routes in `routes.go` to use your handlers

### Disclaimer and considerations for deployment
The deployment scripts, configurations, and any defaults included in this repository are not, under any circumstances, 
to be used in production.  They are here to provide a basic example deployment so you can quickly get up and running in 
a development environment.  For a legitimate deployment, you must first edit them, or create your own configurations 
with secure settings.

As stated above, the responsibility to properly secure your code and deployment is entirely yours, but here are some 
things you may want to consider: 
- Ensure timeouts and size defaults are configured properly for the `http.Server` in `cmd/gloss/main.go`
- Use SSL for communication with the database
- Choose routing middleware to fit your needs
- Properly manage credentials used/shared by the microservice and database

### License
This project is licensed under the terms of the [MIT license](LICENSE).
