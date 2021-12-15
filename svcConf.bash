#!/bin/bash
# docker
sudo apt-get -y install docker.io
# PAT login 
export PAT="${GITHUB_TOKEN}"
echo $PAT | docker login ghcr.io --username tt-loveslife --password-stdin
# postgres
sudo apt install postgresql-client
docker pull ghcr.io/tt-loveslife/collisiondatabase:0.0.2
docker run --name collision -p 5432:5432 -e POSTGRES_DB=crypto -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456  -d ghcr.io/tt-loveslife/collisiondatabase:0.0.2
# redis
sudo apt install redis-tools
docker pull ghcr.io/tygavinzju/redis-sample:1.0.0
docker run -itd --name redis-test -p 6379:6379 ghcr.io/tygavinzju/redis-sample:1.0.0
# go
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
# go-module
go env GO111MODULE
go env -w GO111MODULE="on"
go env -w GOPROXY=https://goproxy.cn
go env GOMODCACHE
go env -w GOMODCACHE=$GOPATH/pkg/mod
# protobuf
sudo apt update
sudo apt install protobuf-compiler
# go mod tidy
go env -w GOPRIVATE=github.com/Zecrey-Labs
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
# pwd : enc-collision-service/service/appService/api/
go mod tidy
# server, pwd : zecrey-collision
cd service/appService/api/
go run appservice.go -f etc/appservice-api.yaml