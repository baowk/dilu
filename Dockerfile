# 声明镜像来源为golang:alpine
FROM golang:1.22 as builder

# 声明工作目录
WORKDIR /root/app/cmd

# 拷贝整个server项目到工作目录
COPY . .

# go generate 编译前自动执行代码
# go env 查看go的环境变量
# go build -o server . 打包项目生成文件名为server的二进制文件

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go build -o server .

# ==================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为alpine:latest
FROM alpine:latest

# 修改镜像源为阿里
# 指定时区
RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.4/main\nhttp://mirrors.aliyun.com/alpine/v3.4/community" >  /etc/apk/repositories \
&& apk update && apk add tzdata \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Shanghai/Asia" > /etc/timezone \
&& apk del tzdata

# 声明工作目录
WORKDIR /root/app/cmd

# 从第一个阶段复制二进制文件到第二个阶段
COPY --from=builder /root/app/cmd/server ./server
COPY --from=builder /root/app/cmd/resources ./resources/
COPY --from=builder /root/app/cmd/docs ./docs/

EXPOSE 80

# 运行打包好的二进制并用-c 指定config.docker.yaml配置文件
CMD ["./server", "start", "-c", "./resources/config.docker.yaml"]
