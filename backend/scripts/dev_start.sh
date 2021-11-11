#!/bin/bash


bash /build/scripts/prestart.sh


# install swaggo/swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generate OpenApi docs
swag init

# Start the app
air