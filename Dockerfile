FROM golang:1.10.3-stretch AS build

WORKDIR /go/src/github.com/yuetsh/Hackathon2018_API

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o zhenxiang .

FROM jrottenberg/ffmpeg:3.3

#COPY ./sources.list /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y ttf-wqy-microhei

COPY --from=build /go/src/github.com/yuetsh/Hackathon2018_API/zhenxiang /

ADD ./templates /

EXPOSE 3010

ENTRYPOINT ["/zhenxiang"]
