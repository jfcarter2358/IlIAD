#! /usr/bin/env bash

combine_coverage() {
    gocovmerge coverage/*.out > coverage/merged.cov

    total_coverage="$(go tool cover -func=coverage/merged.cov | grep -E '^total\:' | sed -E 's/\s+/ /g')"
    echo "Total coverage: ${total_coverage}"

    go tool cover -html=coverage/merged.cov -o coverage/coverage.html
}

start_services() {
    protocol="$1"

    docker compose down --remove-orphans
    docker compose rm -f
    docker compose -f "test.docker-compose.yaml" up -d

    echo "Sleeping for 60 seconds to give everything time to come up"
    sleep 10
    echo "50..."
    sleep 10
    echo "40..."
    sleep 10
    echo "30..."
    sleep 10
    echo "20..."
    sleep 10
    echo "10..."
    sleep 10
    echo "Done!"
}

stop_services() {
    docker compose down --remove-orphans
    docker compose rm -f
}

run_test() {
    start_services

    cd test/tests
    pytest
    cd ..

    rm -rf coverage
    mkdir coverage

    curl "http://localhost:19999"
    curl "http://localhost:29999"

    docker compose cp ${APP_NAME}:/home/${APP_NAME}/cover.out "coverage/${APP_NAME}.cover.out"

    stop_services
}

run_test
combine_coverage
