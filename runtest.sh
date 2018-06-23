#!/bin/bash
htmlout=/tmp/coverage.html
coverout=/tmp/coverage.out

go test -v -coverprofile=$coverout
go tool cover -html=/tmp/coverage.out  -o $htmlout
echo "HTML output written to $htmlout"
