#!/bin/sh

go test $(go list ./... | grep -v /vendor/) -cover
