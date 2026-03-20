# ---- 构建阶段 ----
FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go build -trimpath -ldflags="-s -w" -o server .

# ---- 运行阶段 ----
FROM alpine:3.20

# 时区
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# 非 root 用户运行，提升安全性
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

COPY --from=builder /build/server ./server
COPY --from=builder /build/resources ./resources/
COPY --from=builder /build/docs ./docs/

EXPOSE 7888

# 默认使用 docker 配置；K8s 部署时通过 command 覆盖为 /etc/dilu/config.prod.yaml
CMD ["./server", "start", "-c", "./resources/config.docker.yaml"]
