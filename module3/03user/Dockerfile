# 从alpine建立镜像
FROM alpine
# 将服务编译出来的二进制文件复制到根目录里
ADD user /user
# 设置容器的运行命令为/user
ENTRYPOINT ["/user"]