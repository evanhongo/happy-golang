# Introduction

## This project simply demo
- ## Dependency injection using google wire
- ## Authorization flow using oauth2.0
- ## API using protobuf
- ## Job queue using machinery

---

## Prerequisites
- ## [libvips - image processing library](https://www.libvips.org/)
```
sudo apt install libvips
sudo apt install libvips-tools
```

- ## redis

---

# Folder Structure

## repository
## Responsible for any kind of data fetching or data posting activity. This can be interactions with databases or external API

<br />

## service
## Core business logic implementations are right here

<br />

## other...
## Please refer to golang development convention

---

# Development Procedure
```sh
# Download dependencies in go.mod file
go mod download

# Upgrade dependency and all its dependencies to the latest version
go get -u <dependencies>

# Add any missing modules necessary to build  the current module's packages and dependencies, and remove unused modules
go mod tidy

# Install testing tools
go install github.com/golang/mock/mockgen
go install github.com/onsi/ginkgo/v2/ginkgo

# Generate files for testing purpose
ginkgo bootstrap
ginkgo generate

# Install swagger cli
go install github.com/swaggo/swag/cmd/swag

# Install protobuf cli and related tools
sudo apt update && apt install protobuf-compiler
go install google.golang.org/protobuf
go install github.com/twitchtv/twirp/protoc-gen-twirp

# Install wire cli
go install github.com/google/wire/cmd/wire

# Generate wire_gen.go file
go run github.com/google/wire/cmd/wire
```

---

# Testing
- ## [mocking framework - gomock](https://github.com/golang/mock)
- ## [testing framework - ginkgo](https://onsi.github.io/ginkgo/)