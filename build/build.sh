#/bin/sh
set -x

CI_COMMIT_TAG=$(git describe --always --tags)

docker build -t linclaus/stock-exportor:$CI_COMMIT_TAG -f build/package/Dockerfile .
docker push linclaus/stock-exportor:$CI_COMMIT_TAG