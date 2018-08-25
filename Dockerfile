FROM golang:1.10.3-stretch AS builder

COPY ./sources.list /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y locales locales-all ttf-wqy-microhei ffmpeg

RUN mkdir -p /go/src/golang.org/x/ \
    && cd /go/src/golang.org/x/ \
    && git clone https://github.com/golang/crypto.git crypto \
    && go install crypto

RUN go get -v github.com/labstack/echo

RUN go get -v github.com/labstack/echo/middleware

RUN go get -v github.com/pilu/fresh

ARG project=/go/src/github.com/yuetsh/Hackathon2018_API

WORKDIR ${project}

ADD . .

ENTRYPOINT ["fresh"]

#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o release main.go
#
#FROM scratch
#
#RUN apt-get update \
#    && apt-get install -y locales locales-all ttf-wqy-microhei
#
#ARG project=/go/src/github.com/yuetsh/Hackathon2018_API
#
#COPY --from=builder ${project}/release /
#
#EXPOSE 3010
#
#ENTRYPOINT ["/release"]
