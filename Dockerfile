# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /coba-golang

FROM scratch

COPY --from=build /coba-golang /coba-golang

ENTRYPOINT ["/coba-golang"]