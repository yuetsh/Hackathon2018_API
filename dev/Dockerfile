FROM golang:1.11.1-stretch

COPY ./sources.list /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y ttf-wqy-microhei ffmpeg

RUN go get -u github.com/lib/pq

RUN go get -u github.com/pilu/fresh

WORKDIR /go/src/github.com/yuetsh/Hackathon2018_API

ADD . .
