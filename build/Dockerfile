FROM golang:1.25-alpine3.22 AS builder

LABEL maintainer="zhangyan <zhangyan@weattech.com>"

ENV GOPROXY https://goproxy.cn

RUN echo 'hosts: files dns' > /etc/nsswitch.conf \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && set -ex \
    && apk update && apk upgrade --no-cache \
    && apk add bash tzdata bind-tools busybox-extras ca-certificates libc6-compat \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 复制、重命名构建脚本并赋予执行权限，然后执行它来编译所有应用
COPY build/build_in_docker.sh /build.sh
RUN chmod +x /build.sh && /build.sh

FROM alpine:3.22

# 定义要运行的应用名，可以被 docker build --build-arg 覆盖
ARG APP_NAME=client

# 从 builder 阶段复制时区信息，确保最终镜像时区正确
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone

WORKDIR /app

# 从构建器阶段复制指定的可执行文件
COPY --from=builder "/app/bin/${APP_NAME}" "/app/${APP_NAME}"

# 运行指定应用
CMD ["/app/${APP_NAME}"]