FROM alpine:latest

RUN apk add libc6-compat

ENV TZ=Europe/Moscow

RUN apk add -U --no-cache tzdata && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
