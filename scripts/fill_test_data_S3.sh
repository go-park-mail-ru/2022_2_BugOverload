#!/usr/bin/env bash
printf "Fill S3 storage tests data..."

readonly LOCALSTACK_S3_URL=http://localhost:4566

aws configure set AWS_ACCESS_KEY_ID foo
aws configure set AWS_SECRET_ACCESS_KEY bar

awslocal s3 cp test/data/test.jpeg s3://films/posters/hor/
