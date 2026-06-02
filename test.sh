#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"

./dist/server & SERVER_PID=$!

sleep 1

./dist/gen & GEN_PID=$!

sleep 1

LD_LIBRARY_PATH=$(pwd)/dist ./dist/"${UUT:-uut-go}" & UUT_PID=$!

while kill -0 $GEN_PID >/dev/null 2>&1; do sleep 1; done
kill $UUT_PID $SERVER_PID
