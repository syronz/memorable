#!/usr/bin/env bash

set -e
echo "mode: atomic" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
	go test -race -coverprofile=profile.out -covermode=atomic "$d"
	if [ -f profile.out ]; then
		tail --line=+2 profile.out >> coverage.txt
		rm profile.out
	fi
done
