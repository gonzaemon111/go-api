#!/bin/bash

set -e

host="$1"
shift
cmd="$@"

echo "Waiting for postgresql"
until psql -h "$host" -U "postgres" -c '\q'; do
  echo -n "."
  sleep 1
done

echo "Postgres is up - executing command"
exec $cmd
