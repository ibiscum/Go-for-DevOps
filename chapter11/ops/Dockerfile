# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0
FROM golang:1.22 as build
WORKDIR /app/
COPY . .
RUN go env -w GOPROXY=direct
RUN CGO_ENABLED=0 go build -o ops ops.go
FROM alpine:latest
COPY --from=build /app/ops  /app/ops
CMD ["/app/main"]
