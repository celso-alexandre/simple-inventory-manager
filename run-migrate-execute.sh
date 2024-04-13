#!/bin/bash
echo -e

source .env
# echo "DB_URL: $DB_URL"
# echo "MIGRATION_DIR: $MIGRATION_DIR"
if [ -z "$DB_URL" ]; then
  echo "DB_URL is not set"
  exit 1
fi

if [ -z "$MIGRATION_DIR" ]; then
  echo "MIGRATION_DIR is not set"
  exit 1
fi

COUNT=$1
# echo $COUNT
migrate -path=$MIGRATION_DIR -database $DB_URL up $COUNT
