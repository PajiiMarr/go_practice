#!/bin/sh
set -u
for bin in *_bin; do
    echo "Running $bin..." >&2
    ./"$bin" &
done
wait
