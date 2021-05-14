#!/bin/bash
set -e

for d in ./databases/*/ ; do
    d=${d:12:${#d}-13}
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" --command="create database ${d}"
done
