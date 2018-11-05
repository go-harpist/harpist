#!/usr/bin/env bash
source ./.env/dev
GOOS=linux go build -o ${HANDLER:-main}
