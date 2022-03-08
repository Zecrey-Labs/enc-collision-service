## Service Introduction
* Service provide `getEncCollision` and `getEncCollisionBatches` api to decrypt the collision table
* When service is running firstly, it will call `InitCrypto` func to check whether there  is enough data for decryption. The number of records is `1000000` default.





## Docker
```shell
#1 install docker
sudo apt-get -y install docker.io

#2 login github docker repo
export PAT="${GITHUB_TOKEN}"
echo $PAT | docker login ghcr.io --username ${username} --password-stdin

#3 pull postgres image
docker pull ghcr.io/tt-loveslife/collisiondatabase:0.0.2

#4 run postgres in docker
docker run --name collision -p 5432:5432 -e POSTGRES_DB=crypto -e POSTGRES_USER=<USER> -e POSTGRES_PASSWORD=<PASSWORD>  -d ghcr.io/tt-loveslife/collisiondatabase:0.0.2

#5 pull redis image
docker pull ghcr.io/tygavinzju/redis-sample:1.0.0

#6 run redis in docker
docker run -itd --name redis-test -p 6379:6379 ghcr.io/tygavinzju/redis-sample:1.0.0 --requirepass <PASS>
```

## Environmet
- golang
```shell
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
```

- go-module
```shell
go env -w GO111MODULE="on"
go env -w GOPROXY=https://goproxy.cn
```

- protoc
```shell
sudo apt update
sudo apt install protobuf-compiler
```

- private repo: zecrey-crypto
```shell
go env -w GOPRIVATE=github.com/Zecrey-Labs
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
go mod tidy
```

- start service
```shell
# pwd : zecrey-collision
cd service/appService/api/
go run appservice.go -f etc/appservice-api.yaml
```

## script
- svcConf.bash
This script brings together all the bash commands for installing the development environment
- post.bash, post.lua
These scripts is based on [wrk](https://github.com/wg/wrk). You need to generate a txt file called `enc_data.txt`，which contains enc_data in every single line.
