# 基础镜像
FROM ubuntu:20.04

# 将编译后的linux文件 webook 复制(打包进来)到镜像，放到工作目录 /app 目录下
COPY webook-linux-arm-v-k8s /app/webook
# 设定我的工作目录
WORKDIR /app

# CMD 是执行命令，可以镜像启动之前执行一些命令
# ENTRYPOINT 进来就执行这条命令
ENTRYPOINT ["/app/webook"]
