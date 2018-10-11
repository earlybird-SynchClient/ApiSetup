#!/bin/sh

set -e

if [ "$APP_ENV" = 'production' ]; then
  ApiSetup
else
  go get github.com/pilu/fresh
  fresh
fi
