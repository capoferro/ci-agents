#!bash

set -euo pipefail

##
# since this script is meant to run for the user inside the container, there's
# a bit of hackery we have to do around the fact we're changing to the host
# user's UID. This creates a temporary home dir for storing the module cache
# and a gopath for storing go downloads.
##

if [ ! -w "${GOPATH}" ]
then
  mkdir -p /tmp/gopath /tmp/home
  HOME=/tmp/home
  GOPATH=/tmp/gopath
fi

rm -rf gen/client gen/svc/uisvc/models gen/svc/uisvc/restapi/operations
mkdir -p gen/client/uisvc

/go/bin/swagger generate client --template-dir swagger/templates -f swagger/uisvc/swagger.yml -t gen/client/uisvc
/go/bin/swagger	generate server --template-dir swagger/templates -f swagger/uisvc/swagger.yml -t gen/svc/uisvc

GO111MODULE=off go get github.com/golang/mock/...
${GOPATH}/bin/mockgen -package github github.com/tinyci/ci-agents/clients/github Client > mocks/github/mock.go

protoc -I/go/src /go/src/github.com/tinyci/ci-agents/grpc/types/*.proto --go_out=plugins=grpc:/go/src

for i in $(find grpc/services -maxdepth 1 -type d -name '*' | tail -n +2)
do 
  SPEC=$(basename $i .proto)
  protoc -I/go/src ${PWD}/grpc/services/${SPEC}/server.proto --go_out=plugins=grpc:/go/src
  protoc -I/usr/include/google/protobuf:${PWD}:/go/src ${PWD}/grpc/services/${SPEC}/server.proto --go_out=plugins=grpc:/go/src
done
