#/bin/sh
set -x

CI_COMMIT_TAG=$(git describe --always --tags)

docker build -t linclaus/stock-exportor:latest -f build/package/Dockerfile .