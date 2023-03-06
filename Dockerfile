FROM golang:alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY . .
RUN go mod tidy
RUN go test -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o app .


FROM alpine:latest
WORKDIR /app

ENV ENVIRONMENT=production
ENV PORT=8080
EXPOSE 8080

CMD ["./app"]