#!/bin/bash
echo update ent schema...
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./../ent/schema
echo done.

echo download dependencies again...
go mod tidy
go mod vendor
echo done.

