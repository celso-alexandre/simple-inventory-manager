#!/bin/bash
echo -e
# Reference: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
