FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler