#!/usr/bin/env bash
printf "Fill S3 storage data..."

HOR='hor'
VER='ver'
DEFAULT='default'
AVATAR='avatar'

for file in $(find "$1" -type f -name "*"); do
  if [[ "$file" == *"$HOR"* ]]; then
    awslocal s3 cp "$file" s3://films/posters/hor/
  fi

  if [[ "$file" == *"$VER"* ]]; then
    awslocal s3 cp "$file" s3://films/posters/ver/
  fi

  if [[ "$file" == *"$DEFAULT"* ]]; then
    awslocal s3 cp "$file" s3://default/
  fi

  if [[ "$file" == *"$AVATAR"* ]]; then
      awslocal s3 cp "$file" s3://users/avatar/
    fi
done
