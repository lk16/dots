#!/bin/bash

# disable orphan warnings
export COMPOSE_IGNORE_ORPHANS=True

docker-compose -f docker-compose-test.yml up --build --abort-on-container-exit

docker-compose -f docker-compose-test.yml down --volumes