# 已迁移
由于 :

1. 接口定位不符 ( 预期发布的接口的内容不仅限于营销类 ) 
2. Serverless Components 发布了 V2 版本, 迁移到 V2 需要对项目整体结构进行大幅改动 
3. 不希望造成文件丢失或干扰 ( 主要是被 gitignore 的日志、配置、二进制程序或测试文件 )

综合以上, 决定将本项目整体迁移到: 

> [serverless-apis](https://github.com/offcn-jl/serverless-apis)

# marketing-apis
Marketing APIs

部署步骤:

1. 运行 sls， 扫码登陆

...

> 配置 NPM 代理 ： https://blog.csdn.net/qq_31945977/article/details/81537917
