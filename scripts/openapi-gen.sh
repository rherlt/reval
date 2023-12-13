#!/bin/bash
echo download dependencies...
go mod tidy
echo done.

#go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

echo generating types...
~/go/bin/oapi-codegen \
  -config ../configs/debug/evaluationapi.cfg.yaml \
    ../api/evaluationapi.yaml
echo done.

echo download dependencies again...
go mod tidy
go mod vendor
echo done.