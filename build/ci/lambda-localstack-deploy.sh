#!/usr/bin/env bash
source ./.env/dev
FUNCTION=$(basename $(pwd))
rm -rf handler.zip && zip handler.zip ./main
awslocal lambda create-function \
  --function-name ${FUNCTION} \
  --role arn:aws:iam::1234:role/role_name \
  --memory 128 \
  --runtime go1.x \
  --zip-file fileb://$(pwd)/handler.zip \
  --handler ${HANDLER:-main} || echo "Failed to deploy - is awslocal installed or is localstack down?"
