FROM alpine:v3.17

RUN apk add libc6-compat
RUN ln -s /lib/libc.so.6 /usr/lib/libresolv.so.2

ENV TZ=Europe/Moscow

RUN apk add -U --no-cache tzdata && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
