#build stage
FROM golang:1.13.8-alpine as builder
RUN apk add --no-cache git

WORKDIR src/GIG
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/lsflk/gig-sdk
RUN revel version
RUN cd 
RUN revel build "" build prod

FROM alpine:latest
WORKDIR /
COPY --from=builder /go/src/GIG/build /
ENTRYPOINT ["sh", "./build/run.sh"]