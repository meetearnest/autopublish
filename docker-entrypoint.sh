#!/usr/bin/env bash
set -e

export PATH=${PATH}:/usr/src/app/node_modules/.bin

# NOTE: I'm leaving this here as a warning. Do not put anything here that could
# echo output.
# echo "npm install"
# npm install
exec $@
