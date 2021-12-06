#!/bin/bash
# docker
sudo apt-get -y install docker.io
# PAT login 
export PAT="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"
echo $PAT | docker login ghcr.io --username tt-loveslife --password-stdin
# postgres
docker pull ghcr.io/zecrey-labs/zecrey-postgres-sample:0.0.3
docker run --name zecreypostgres -p 5432:5432 -e POSTGRES_DB=crypto -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456  -d ghcr.io/zecrey-labs/zecrey-postgres-sample:0.0.3
# redis
docker pull ghcr.io/tygavinzju/redis-sample:1.0.0
docker run -itd --name redis-test -p 6379:6379 ghcr.io/tygavinzju/redis-sample:1.0.0
# server, pwd : zecrey-collision
cd service/appService/api/
go run appservice.go -f etc/appservice-api.yaml