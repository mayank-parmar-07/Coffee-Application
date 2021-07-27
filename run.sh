#!/bin/bash  
if [[ "$1" == "true" ]]; then
    go test ./... --parallel=4
else
    go run main.go
fi