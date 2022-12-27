ARG PACKAGE_VERSION

## go build
##
FROM golang:1.19 as builder

ARG PACKAGE_VERSION

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY internal internal

RUN go build \
    -o bin/unifi-blueberry-core \
    -ldflags "\
      -X internal/core.BuildVersion=$PACKAGE_VERSION \
      -X internal/core.BuildPlatform=$(go env GOOS)/$(go env GOARCH) \
      -X internal/core.BuildGitCommit=$(git rev-list -1 HEAD) \
      -X internal/core.BuildGoVersion=$(go env GOVERSION)"

## package
##
FROM debian:bullseye as package

ARG PACKAGE_VERSION

RUN apt-get update && apt-get install -y \
      gettext \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /build

COPY build/package.sh .
COPY build/packageroot/ packageroot/

COPY --from=builder /build/bin/* bin/

RUN sh package.sh

## artifact for export
##
FROM scratch as artifact

COPY --from=package /build/out/ /
