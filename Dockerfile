#build stage
FROM golang:1.14-alpine AS builder
WORKDIR /root/go/src/GIG
COPY . .
RUN apk add --no-cache git
RUN go get github.com/lsflk/gig-sdk
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN revel build "" build -m prod

#running stage
EXPOSE 9000

ENTRYPOINT ["sh", "./build/run.sh"]
