#!/usr/bin/env bash
printf "Fill S3 storage data..."

LOCALSTACK_S3_URL=$2

# Создаем bucket для фильмов
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket films

# Создаем bucket для стандартных картинок
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket default

# Создаем bucket для пользователей
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket users

HOR='hor'
VER='ver'
DEFAULT='default'
AVATAR='avatar'

for file in $(find "$1" -type f -name "*"); do
  if [[ "$file" == *"$HOR"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://films/posters/hor/
  fi

  if [[ "$file" == *"$VER"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://films/posters/ver/
  fi

  if [[ "$file" == *"$DEFAULT"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://default/
  fi

  if [[ "$file" == *"$AVATAR"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://users/avatar/
  fi
done
