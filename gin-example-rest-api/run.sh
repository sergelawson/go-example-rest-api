#!/bin/bash
npx nodemon --watch . --ext go  --exec "go build -o server && pkill -f server || true && ./server"