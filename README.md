# SDCS
分布式课程作业

运行流程：
```shell
# 程序打包
docker build -t sdcs .
# 运行容器
docker-compose up -d
# 关闭容器
docker-compose down
```
注意：需确保能网络环境能联通外网，否则docker build大概率失败。
