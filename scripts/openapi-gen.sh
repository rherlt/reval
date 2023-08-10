#!/bin/bash
echo download dependencies...
go mod tidy
echo done

echo generating types...
~/go/bin/oapi-codegen \
  -config ../configs/debug/evaluationapi.cfg.yaml \
    ../api/evaluationapi.yaml
echo done

echo download dependencies again...
go mod tidy
echo done