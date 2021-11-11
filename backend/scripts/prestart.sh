#!/bin/bash

while ! nc -z $POSTGRES_SERVER $POSTGRES_PORT; do
  echo "Wating for Database !!!!!!!!!!!!"
  sleep 0.1
done