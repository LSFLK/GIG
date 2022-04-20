#build stage
FROM golang:1.14-alpine AS builder
WORKDIR /root/go/src/GIG/
COPY . .
ENV GOPATH=/root/go/
ENV PATH="/root/go/bin:${PATH}"

# RUN export PATH="$PATH:/root/go/bin/"
RUN echo $PATH
RUN echo $GOPATH
RUN apk add --no-cache git
RUN go get -v github.com/lsflk/gig-sdk 
RUN go get -v github.com/revel/revel 
RUN go get -v github.com/revel/cmd/revel 
RUN revel build "" build -m prod

#running stage
EXPOSE 9000

ENTRYPOINT ["sh", "./build/run.sh"]
