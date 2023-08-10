# syntax=docker/dockerfile:1

# --------------------------------------------
# Build the go backend application from source
# --------------------------------------------
FROM golang:1.20 AS go-build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

#copy all repo files
COPY . .

# create code for OpenApi spec
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.4

# generate open api spec for go server
RUN $GOPATH/bin/oapi-codegen -config configs/docker/evaluationapi.cfg.yaml api/evaluationapi.yaml
# build application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app/ ./... 

# -----------------------------------------
# Run the go backend tests in the container
# -----------------------------------------
FROM go-build-stage AS go-run-test-stage
RUN go test -v ./...


# --------------------------------------------------
# Build the Angular frontend application from source
# --------------------------------------------------
FROM node:20.5-alpine AS node-build-stage

WORKDIR /app

COPY ui/reval-web/package.json ui/reval-web/package-lock.json ./

RUN apk update && apk add openjdk17-jre && npm install @openapitools/openapi-generator-cli && npm install
COPY . .
RUN cd ui/reval-web/ && npm run build

# -----------------------------------------------------------------
# Deploy the build outputs and all required files into a lean image
# -----------------------------------------------------------------
FROM gcr.io/distroless/base-debian11 AS build-release-stage


WORKDIR /app

COPY --from=go-build-stage /usr/local/bin/app/ /app/configs/docker/reval.env /app/
COPY --from=go-build-stage /app/data /app/data
COPY --from=node-build-stage /app/ui/reval-web/dist/reval-web/ /app/www

# make sure to run webserver in release mode
ENV GIN_MODE=release

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./reval"]
