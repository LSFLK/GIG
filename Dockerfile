#Compile stage
FROM golang:1.13.8-alpine AS builder

# Add required packages
RUN apk add  --no-cache --update git curl bash

RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/lsflk/gig-sdk

WORKDIR /go/src/GIG
RUN pwd
ADD go.mod go.sum ./
RUN go mod download
ENV CGO_ENABLED 0 \
    GOOS=linux \
    GOARCH=amd64
ADD . .

RUN revel package . -m prod

# Run stage
FROM alpine:3.13
EXPOSE 9000
RUN apk update
WORKDIR /
RUN ls
COPY --from=builder /go/src/GIG/GIG.tar.gz .
RUN tar -xzvf GIG.tar.gz && rm GIG.tar.gz
RUN pwd
ENTRYPOINT /go/src/GIG/build/run.sh
