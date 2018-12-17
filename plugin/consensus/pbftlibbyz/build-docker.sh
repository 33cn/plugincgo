#!/bin/sh

dir=$(cd "$(dirname "$0")" || "exit -1"; pwd)
cd "${dir}/../../../" || "exit -1"
docker build -t pbftlibbyz -f "$dir/Dockerfile" .
