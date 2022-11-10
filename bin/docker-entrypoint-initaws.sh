#!/usr/bin/env bash
printf "Configuring localstack components..."

readonly LOCALSTACK_S3_URL=http://localstack:4566

sleep 5

set -x

echo "INFO"
echo "$AWS_PROFILEee"

echo "[default]" > ~/.aws/config
echo "region = us-east-1" >> ~/.aws/config
echo "output = json" >> ~/.aws/config

set +x
