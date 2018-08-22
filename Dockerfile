FROM golang:alpine AS builder

RUN cp /etc/apk/repositories /etc/apk/repositories.bak

RUN echo "http://mirrors.aliyun.com/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk add --no-cache git

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

#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o release src/main.go
#
#FROM scratch
#
#ARG project=/go/src/github.com/yuetsh/Hackathon2018
#
#COPY --from=builder ${project}/release /
#
#EXPOSE 3010
#
#ENTRYPOINT ["/release"]
