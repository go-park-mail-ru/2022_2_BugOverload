FROM alpine:latest

RUN apk add libc6-compat
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime

ENV TZ=Europe/Moscow

RUN apk del tzdata
