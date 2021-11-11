#!/bin/bash


bash /build/scripts/prestart.sh

go build -o api .

./api