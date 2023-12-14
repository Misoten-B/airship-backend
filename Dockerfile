FROM golang:1.21.3-alpine3.17 AS base

ARG SA_TYPE
ARG SA_PROJECT_ID
ARG SA_PRIVATE_KEY_ID
ARG SA_PRIVATE_KEY
ARG SA_CLIENT_EMAIL
ARG SA_CLIENT_ID
ARG SA_AUTH_URI
ARG SA_TOKEN_URI
ARG SA_AUTH_PROVIDER_X509_CERT_URL
ARG SA_CLIENT_X509_CERT_URL
ARG SA_UNIVERSE_DOMAIN

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod tidy
RUN <<EOR
tee serviceAccountKey.json <<EOF
{
  "type": "${SA_TYPE}",
  "project_id": "${SA_PROJECT_ID}",
  "private_key_id": "${SA_PRIVATE_KEY_ID}",
  "private_key": "${SA_PRIVATE_KEY}",
  "client_email": "${SA_CLIENT_EMAIL}",
  "client_id": "${SA_CLIENT_ID}",
  "auth_uri": "${SA_AUTH_URI}",
  "token_uri": "${SA_TOKEN_URI}",
  "auth_provider_x509_cert_url": "${SA_AUTH_PROVIDER_X509_CERT_URL}",
  "client_x509_cert_url": "${SA_CLIENT_X509_CERT_URL}",
  "universe_domain": "${SA_UNIVERSE_DOMAIN}"
}
EOF
EOR

FROM base AS dev
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
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
COPY --from=builder /app/serviceAccountKey.json serviceAccountKey.json
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/golang-app golang-app
USER scratch
EXPOSE 8080
CMD ["/golang-app"]
