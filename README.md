# Golang base

This is a project template for future Golang projects of my own.

## Features

* Simple environment variables picking and setting.
* Basic Docker configuration file to build images for production.
* Testing frameworks added.
* Gin-Gonic HTTP framework port 8080 (configurable through env), with status
  handler.

## Usage

### Build Docker image
```
docker build . -t ariel17/base
```

### Using environment variables file
Add keys to `.env` file:
```
MY_SECRET_KEY1=v4lu3!#
```

Make Docker pick them as follows:
```
docker run --env-file .env ariel17/base
```

### Build Swagger documentation
```
swag init -o api
```

* Served on http://localhost:8080/swagger/index.html
* Swaggo docs: https://github.com/swaggo/swag#getting-started