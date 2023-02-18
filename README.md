# Vector mongodb sink

MongoDB sink for [Vector](https://vector.dev/) based on HTTP server

Written on Go using [Fiber](https://docs.gofiber.io/)

## Run

`go run ./cmd/vector-mongodb-sink/main.go`

## Build

`go build ./cmd/vector-mongodb-sink`

## Env Vars Config

* `APP_HOST` - API server host
* `APP_PORT` - API server port
* `APP_MONGODBURI` - (Required) URI to MongoDB
* `APP_MONGODBNAME` - (Required) DB name
* `APP_AUTHUSERNAME` - User username for Basic Auth
* `APP_AUTHPASSWORD` - User password for Basic Auth