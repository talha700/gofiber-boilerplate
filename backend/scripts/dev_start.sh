#!/bin/bash


bash /build/scripts/prestart.sh

# install air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# install swaggo/swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generate OpenApi docs
swag init

# Start the app
air