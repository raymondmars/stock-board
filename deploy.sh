#!/bin/sh
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -a -o stock-board . && docker-compose build && docker-compose push && rm ./stock-board