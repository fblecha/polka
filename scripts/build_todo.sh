#!/bin/sh
rm -rf ./tmp/
./polka new tmp/todo
cd tmp/todo
../../polka generate endpoint items
#./polka local server
#./polka deploy
