#! /bin/bash

set -e

if [ ! -n "$1" ]; then
  echo "No argument, set build deafult"
  go build main.go
elif [[ "$1" = "build" ]]; then
  go build main.go
elif [[ "$1" = "run"  ]]; then
  go run main.go
else
  echo "wrong argument!"
fi
