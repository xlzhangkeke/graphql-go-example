#!/bin/bash
set -e
cd "$(dirname "$0")"/
migrate -database=postgres://postgres:pandora@10.95.84.99:25432/graphql -source=./migrations "$@"
