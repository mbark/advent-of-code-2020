#!/usr/bin/env bash
set -euo pipefail

while [[ $# -gt 0 ]]
do
key="$1"
case $key in
    -d|--day)
    DAY="$2"
    shift # past argument
    shift # past value
    ;;
esac
done

perf() {
    echo "=== DAY $DAY ==="
    if [ -f "day$DAY/main.go" ]; then
        go build "day$DAY/main.go"
        hyperfine './main' --style basic --warmup 10
    else
        echo "No code written..."
    fi
}

if [ -n "${DAY-}" ]; then
    perf
    exit 0
fi

for DAY in {1..25} .. N
do
    perf
done
