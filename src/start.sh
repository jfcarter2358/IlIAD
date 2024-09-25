#! /usr/bin/env bash

sleep_amount="${SLEEP:-"0"}"
echo "Sleeping for ${sleep_amount} seconds"
sleep "${sleep_amount}"

run_mode="${RUN_MODE:normal}"

RUN_DIR=$(dirname $0)

pushd "${RUN_DIR}"

echo "Starting iad"
if [[ "${run_mode}" == "normal" ]]; then
    ./iad
else 
    while true; do
        echo "Service started in coverage mode"
        ./iad -test.coverprofile=cover.out "$@" || exit 1;
        echo "Server restarting.."
    done
fi

popd
