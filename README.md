# Backend challenge from Ariel Gerardo Ríos (that's me!)

## Usage
TODO

### Build Docker image
```
docker build . -t be-challenge-arios
```

### Using environment variables file
Add keys to `.env` file:
```
MY_SECRET_KEY1=v4lu3!#
```

Make Docker pick them as follows:
```
docker run --env-file .env be-challenge-arios
```

### Build Swagger documentation
```
swag init -o api
```

* Served on http://localhost:8080/swagger/index.html
* Swaggo docs: https://github.com/swaggo/swag#getting-started
