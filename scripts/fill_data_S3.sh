#!/usr/bin/env bash
printf "Fill S3 storage data..."

LOCALSTACK_S3_URL=$2

# Создаем bucket для фильмов
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket films

# Создаем bucket для стандартных картинок
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket default

# Создаем bucket для пользователей
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket users

# Создаем bucket для персон
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket persons

# Создаем bucket для коллекций
aws --endpoint-url="$LOCALSTACK_S3_URL" s3api create-bucket --bucket collections

HOR='hor'
VER='ver'
DEFAULT='default'
USER_AVATAR='users_avatars'
PERSON_AVATAR='persons_avatars'
PERSON_IMAGES='persons_images/'
FILM_IMAGES='films_images/'
COLLECTION_POSTERS='collections_posters/'

for file in $(find "$1" -type f -name "*"); do
  if [[ "$file" == *"$HOR"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://films/posters/hor/

    continue
  fi

  if [[ "$file" == *"$VER"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://films/posters/ver/

    continue
  fi

  if [[ "$file" == *"$DEFAULT"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://default/

    continue
  fi

  if [[ "$file" == *"$USER_AVATAR"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://users/avatars/

    continue
  fi

  if [[ "$file" == *"$PERSON_AVATAR"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://persons/avatars/

    continue
  fi

  if [[ "$file" == *"$PERSON_IMAGES"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://persons/images/"${file##*$PERSON_IMAGES}"

    continue
  fi

  if [[ "$file" == *"$FILM_IMAGES"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://persons/images/"${file##*$FILM_IMAGES}"

    continue
  fi

  if [[ "$file" == *"$COLLECTION_POSTERS"* ]]; then
    aws --endpoint-url="$LOCALSTACK_S3_URL" s3 cp "$file" s3://collections/posters/

    continue
  fi

done
