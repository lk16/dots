#!/bin/bash

# disable orphan warnings
export COMPOSE_IGNORE_ORPHANS=True

docker-compose -f docker-compose-test.yml up --build --abort-on-container-exit

docker cp dots_test:/app/coverage.out /tmp/dots_coverage.out
go tool cover -html="/tmp/dots_coverage.out" 2>&1 >/dev/null &

docker-compose -f docker-compose-test.yml down --volumes