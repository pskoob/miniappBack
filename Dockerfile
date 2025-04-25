FROM golang:1.23 AS build

WORKDIR /src

RUN apt-get update && apt-get install -y git build-essential

COPY go.mod go.sum ./
COPY vendor vendor/
COPY . .

RUN go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/mini-app ./cmd/mini-app

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    libssl3 && rm -rf /var/lib/apt/lists/*

COPY --from=build /bin/mini-app /bin/mini-app
COPY --from=build /src/pkg/nearffi/libnear_bridge.so /usr/lib/

RUN ldconfig
ENV LD_LIBRARY_PATH=/usr/lib

ENTRYPOINT ["/bin/mini-app"]
