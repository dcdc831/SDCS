FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list &&\
    sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y curl
RUN curl -O https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
RUN tar xvf go1.21.3.linux-amd64.tar.gz
RUN mv go /usr/local
ENV PATH=$PATH:/usr/local/go/bin

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
WORKDIR /go/SDCS
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY ./hash ./hash
COPY ./httpagent ./httpagent
COPY ./jsontrans ./jsontrans
COPY ./node ./node
COPY ./proto ./proto
RUN go build -o SDCS
ENTRYPOINT [ "./SDCS" ]