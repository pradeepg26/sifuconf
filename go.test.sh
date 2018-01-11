#!/usr/bin/env bash

set -e
echo "" > coverage.txt
rm -f $CIRCLE_TEST_REPORTS/report.xml

for d in $(go list ./... | grep -v vendor); do
  mkdir -p $CIRCLE_TEST_REPORTS/$d
  go test -v -race -coverprofile=profile.out -covermode=atomic $d | go-junit-report > ./$CIRCLE_TEST_REPORTS/${d}_report.xml
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done

junit-merge -r -o $CIRCLE_TEST_REPORTS/report.xml $CIRCLE_TEST_REPORTS
