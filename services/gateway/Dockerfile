FROM golang:1.21 as dependencies
WORKDIR /claimclam
COPY go.* .
RUN go mod download

FROM dependencies as builder
COPY . .
RUN make service.gateway

# https://docs.aws.amazon.com/inspector/latest/user/supported.html#supported-os-ecr
FROM alpine:latest as prod
RUN apk add --update --no-cache make ca-certificates git
COPY --from=builder /claimclam/bin/gateway-cli /gateway-cli
ENTRYPOINT ["/gateway-cli", "start"]