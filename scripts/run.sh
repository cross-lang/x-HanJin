#!/bin/sh

# x-HanJin 启动脚本
# 同时启动 Go 服务和 Nginx 反向代理

# 启动 Go 服务（后台运行）
./server &

# 等待 Go 服务启动
sleep 5

# 启动 Nginx（前台运行）
nginx -t && nginx -g 'daemon off;'

echo "run successfully!"
