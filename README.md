# ProxyPool


## 说明

基于Go开发的爬虫代理IP池。

网上很多提供免费代理的网站，但是每个网站提供的IP不多，并且很多无法使用。

该项目通过定时爬取各网站的代理IP， 验证可用之后再入库, 然后提供API给客户端使用。

同时欢迎参与扩展代理源以增加代理池IP的质量和数量。

## 安装

- 使用代码方式运行:

```shell
git clone https://github.com/ClassmateLin/proxy_pool.git 
cd proxy_pool && go mod download
go run manage.go
```

- 使用docker-compose构建运行:
```shell
python -m pip install docker-compose
docker-compose up -d
```

- 使用docker镜像运行:

```shell

```

## 使用


