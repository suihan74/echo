#!/bin/sh

if [ -x "$HOME/backend" ]; then
	"$HOME/backend"
else
	go run "*.go"
fi
