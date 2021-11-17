#/bin/bash
RUN_NAME="gf.bridgx.api"
mkdir -p output/conf output/bin
sudo mkdir -p /tmp/bridgx/logs/

find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/
cp scripts/run_api.sh output/

go fmt ./...
go vet ./...
export GO111MODULE="on"
export GOPRIVATE="code.galaxy-future.com"

go build -o output/bin/${RUN_NAME} ./cmd/api