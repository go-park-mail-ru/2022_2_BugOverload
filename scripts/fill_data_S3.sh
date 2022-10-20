#!/usr/bin/env bash
printf "Fill S3 storage data..."

for file in $(find "$1" -type f -name "*"); do
  if grep -q hor "$file"; then
    awslocal s3 cp "$file" s3://films/posters/hor/
  fi
done
