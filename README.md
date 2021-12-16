# enc-collision-service开发入门
Windows下开发环境使用`Goland + WSL2 + Docker`体验较好

## Docker环境
```shell
#1 安装docker
sudo apt-get -y install docker.io

#2 登陆github提供的 docker 仓库
export PAT="${GITHUB_TOKEN}"
echo $PAT | docker login ghcr.io --username ${username} --password-stdin

#3 拉取 postgres 镜像
docker pull ghcr.io/tt-loveslife/collisiondatabase:0.0.2

#4 启动 postgres 镜像
docker run --name collision -p 5432:5432 -e POSTGRES_DB=crypto -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456  -d ghcr.io/tt-loveslife/collisiondatabase:0.0.2

#5 拉取 redis 镜像
docker pull ghcr.io/tygavinzju/redis-sample:1.0.0

#6 启动 redis 镜像
docker run -itd --name redis-test -p 6379:6379 ghcr.io/tygavinzju/redis-sample:1.0.0
```

## 开发环境
- go语言安装
```shell
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
# 最好持久化到环境变量
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
```

- go-module配置
```shell
go env -w GO111MODULE="on"
go env -w GOPROXY=https://goproxy.cn
```

- 安装protoc
```shell
sudo apt update
sudo apt install protobuf-compiler
```

- zecrey-crypto私有仓库依赖
```shell
go env -w GOPRIVATE=github.com/Zecrey-Labs
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
go mod tidy
```

- 项目启动
```shell
# pwd : zecrey-collision
cd service/appService/api/
go run appservice.go -f etc/appservice-api.yaml
```

## 脚本介绍
- svcConf.bash
该脚本汇集了安装开发环境的所有指令
- post.bash, post.lua
这两个脚本运行依赖于压测脚本 [wrk](https://github.com/wg/wrk)， 只需要编写一个 `enc_data.txt`的文本文件，并在每行放置一个 `加密余额` 参数值， 该脚本即可对指定地址压力访问