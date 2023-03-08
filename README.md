# Backend challenge by Ariel Gerardo Ríos (that's me!)

## Start services with docker-compose
It is REQUIRED to change `FOOTBALL_APIKEY` environment variable value in
`docker-compose.yml` file.

```
# It can take a few moments since it's building + testing + waiting for db to be
# accepting connections
docker-compose up -d
```

## Swagger documentation
Served on http://localhost:8080/swagger/index.html

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
```