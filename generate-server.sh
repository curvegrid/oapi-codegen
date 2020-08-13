#!/bin/sh

set -e

go run cmd/oapi-codegen/oapi-codegen.go -package spec -o ../multibaas/server/app/web/api/spec/generated.go -generate types,server /home/natsukagami/curvegrid/multibaas-openapi/reference/MultiBaas-API.v1.yaml

goimports -w ../multibaas/server
