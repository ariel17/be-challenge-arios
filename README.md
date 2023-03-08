# Backend challenge by Ariel Gerardo Ríos (that's me!)

## Start services with docker-compose
It is REQUIRED to change `FOOTBALL_APIKEY` environment variable value in
`docker-compose.yml` file. Once started, it can take a few moments up be ready
since it builds + tests + waits for database to be ready to accept connections.

```
docker-compose up
```

## Example usage
```
# enqueues data import process for Premier League competition
curl -X POST --data '{"code":"PL"}' http://localhost:8080/importer

# Get players in Prime League that belongs to a Manches-like team (teamName optional)
curl -X GET "http://localhost:8080/competitions/pl/players?teamName=Manches"

# Get Manchester United team with players (showPlayers optional)
curl -X GET "http://localhost:8080/teams/MUN?showPlayers=true"

# Get Manchester United team persons (players and coaches if present)
curl -X GET "http://localhost:8080/teams/MUN/persons"

# Utility endpoints
curl -X GET "http://localhost:8080/status"
```

## Swagger documentation
Served on http://localhost:8080/swagger/index.html