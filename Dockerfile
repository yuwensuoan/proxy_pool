FROM golang:1.16 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn" \
    CGO_ENABLED=0;

WORKDIR /proxy_pool

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o manage .

#FROM golang:alpine
#
#WORKDIR /proxy_pool
#
#COPY --from=builder /proxy_pool/config .
#COPY --from=builder /proxy_pool/manage .



