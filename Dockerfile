# syntax=docker/dockerfile:1.5
FROM golang:1.22.2-alpine3.19 as build

# hadolint ignore=DL3018
RUN apk update && apk add --no-cache \
    git ca-certificates tzdata build-base && \
    update-ca-certificates

ENV USER=appuser
ENV UID=1000

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/mypackage/app/

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=. \
    GOCACHE=/root/.cache/go-build \
    GOMODCACHE=/go/pkg/mod \
    GOOS=linux GOARCH=amd64 \
    CGO_ENABLED=0 GOGC=off \
    go build \
    -ldflags='-w -s -extldflags "-static"' \
    -o /go/bin/app ./cmd/.

FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

COPY --from=build /go/bin/app /go/bin/app

USER appuser:appuser

ENTRYPOINT ["/go/bin/app"]
