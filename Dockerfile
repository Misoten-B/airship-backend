FROM golang:1.21.3-alpine3.17 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod tidy

FROM base AS dev
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest
COPY . .
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
COPY . .
RUN addgroup -g 10001 scratchgroup && \
    adduser -u 10001 -G scratchgroup -D scratch

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go build \
    -ldflags="-s -w" \
    -o golang-app \
    -trimpath

FROM scratch AS runner
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/golang-app golang-app
USER scratch
CMD ["/golang-app"]
