#!/usr/bin/env bash
source ./.env/dev
FUNCTION=$(basename $(pwd))
awslocal lambda invoke --function-name ${FUNCTION} test.txt
