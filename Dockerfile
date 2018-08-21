FROM golang:alpine AS builder

RUN cp /etc/apk/repositories /etc/apk/repositories.bak

RUN echo "http://mirrors.aliyun.com/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk add --no-cache git

RUN go get github.com/lib/pq

RUN go get github.com/pilu/fresh

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
#EXPOSE 3000
#
#ENTRYPOINT ["/release"]
