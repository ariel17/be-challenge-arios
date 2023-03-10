FROM golang:alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY . .
RUN go mod tidy
RUN go test -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o app .


FROM alpine:latest
WORKDIR /app
COPY --from=build /build/app .

ENV PORT=8080
ENV FOOTBALL_APIKEY=apikey
ENV FOOTBALL_MAX_REQUESTS_PER_MINUTE=10
ENV DATABASE_DSN=dsn
ENV DATABASE_STATUS_QUERY="SELECT 1"
EXPOSE 8080

CMD ["./app"]