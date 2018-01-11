#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
  mkdir -p ./reports/$d
  go test -race -coverprofile=profile.out -covermode=atomic $d | go-junit-report > ./reports/${d}_report.xml
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done
