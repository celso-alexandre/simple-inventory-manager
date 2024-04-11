#!/bin/bash
echo -e

echo "Migration name: "
read name
migrate create -ext sql -dir ./migrations -seq $name
