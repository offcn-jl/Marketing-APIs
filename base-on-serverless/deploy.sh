#! /bin/bash

# 编译 Golang 二进制文件
sh build.sh

# 拼接 Serverless 配置
# https://blog.csdn.net/Alvin_Lam/article/details/79123882
if [ "$1" == "release" ]; then
    cat serverless-release.yml serverless-config.yml > serverless.yml
else
    cat serverless-dev.yml serverless-config.yml > serverless.yml
fi

# 部署
sls --debug
