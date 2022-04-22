#Compile stage
FROM golang:1.13.8-alpine AS builder

# Add required packages
RUN apk add  --no-cache --update git curl bash

RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/lsflk/gig-sdk

WORKDIR /go/src/GIG
ADD go.mod go.sum ./
RUN go mod download
ENV CGO_ENABLED 0 \
    GOOS=linux \
    GOARCH=amd64
ADD . .
RUN revel build "" build -m prod

# Run stage
FROM alpine:3.15
EXPOSE 9000
COPY --from=builder /go/src/GIG/build /build
RUN mkdir /build/app && mkdir /build/app/cache
WORKDIR /build
ENTRYPOINT ./run.sh