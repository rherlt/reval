# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app/ ./... 
# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

ENV GIN_MODE=release

WORKDIR /app

COPY --from=build-stage /usr/local/bin/app/ /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./reval"]
