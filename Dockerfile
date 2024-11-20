#
# Build stage
#

FROM golang:1.23-alpine as compiler
WORKDIR /go/src
COPY main.go main.go
COPY go.mod go.mod
RUN go mod tidy
RUN go build -o service

FROM alpine
COPY --from=compiler /go/src/service /go/src/service
RUN apk add --no-cache curl
WORKDIR /go/src
CMD ./service
EXPOSE 8080