#!/bin/sh

set -e

echo "start migration"
make migrateup
echo "Starting server"
make server