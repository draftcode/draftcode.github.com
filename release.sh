#!/bin/bash -e
cd "$(dirname "$0")"
rm -rf docs
(cd devserve && go build && ./devserve -build)
