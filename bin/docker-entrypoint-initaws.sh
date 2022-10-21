#!/usr/bin/env bash
printf "Configuring localstack components..."

readonly LOCALSTACK_S3_URL=http://localstack:4566

sleep 5

set -x

aws configure set aws_access_key_id foo
aws configure set aws_secret_access_key bar
echo "[default]" > ~/.aws/config
echo "region = us-east-1" >> ~/.aws/config
echo "output = json" >> ~/.aws/config

# Создаем bucket для фильмов
aws --endpoint-url=$LOCALSTACK_S3_URL s3api create-bucket --bucket films

# Создаем bucket для стандартных картинок
aws --endpoint-url=$LOCALSTACK_S3_URL s3api create-bucket --bucket default

set +x
