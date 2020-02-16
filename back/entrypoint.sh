#!/bin/sh

if [ -x backend ]; then
	./backend
else
	go run *.go
fi
