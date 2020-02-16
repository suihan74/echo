#!/bin/sh

if [ "$NODE_ENV" = "development" ]; then
	npm run dev
elif [ "$NODE_ENV" = "production" ]; then
	# Describe what you need for deployment in production
	exit 0
fi
