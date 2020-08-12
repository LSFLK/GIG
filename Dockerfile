#build stage
FROM golang:1.14-alpine AS builder
WORKDIR /go/src/GIG
COPY . .
RUN apk add --no-cache git
RUN git clone https://github.com/LSFLK/GIG-SDK.git /go/src/GIG-SDK
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN revel build "" build -m prod

#running stage
FROM alpine:3.9 
COPY --from=builder /go/src/GIG/build /app/GIG
ENTRYPOINT ["sh", "/app/GIG/run.sh"]
EXPOSE 9000