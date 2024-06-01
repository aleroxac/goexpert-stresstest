FROM golang:1.22.3 AS base



# Stage 1: Build the binary
FROM base AS builder
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main cmd/main.go



# Stage 2: Compress the binary using UPX
FROM alpine AS upx
RUN apk add --no-cache upx
COPY --from=builder /build/main /upx/main
RUN upx --best --lzma /upx/main -o /upx/main_compressed



# Stage 3: Create the final image
FROM scratch AS main
WORKDIR /app
COPY --from=upx /upx/main_compressed /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/app/main" ]
