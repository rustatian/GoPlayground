# Image page: <https://hub.docker.com/_/golang>
FROM golang:1.16.3 as builder

# arguments to pass on each go tool link invocation
ENV LDFLAGS="-s"

RUN mkdir /src

WORKDIR /src

COPY ./go.mod ./go.sum ./

# Burn modules cache
RUN set -x \
    && go version \
    && go mod download \
    && go mod verify

COPY . .

# compile binary file without local pathes and debug information
RUN CGO_ENABLED=0 go build -trimpath -ldflags "$LDFLAGS" -o ./push-notification-service ./cmd/linux

# Image page: <https://hub.docker.com/_/alpine>
FROM alpine:3.13

# copy required files from builder image
COPY --from=builder /src/push-notification-service /usr/bin/push-notification-service
COPY --from=builder /src/tests/appspace.json /usr/bin/appspace.json

ENTRYPOINT ["/usr/bin/push-notification-service"]