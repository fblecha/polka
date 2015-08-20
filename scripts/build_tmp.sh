#!/bin/sh
rm -rf ./tmp/
./polka new tmp/todo
./polka generate endpoint items
