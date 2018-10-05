FROM golang:1.11.1-stretch AS builder

WORKDIR /go/src/github.com/yuetsh/Hackathon2018_API

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go install -ldflags "-s -w" .

FROM jrottenberg/ffmpeg:4.0-ubuntu

WORKDIR /opt/api

RUN apt-get install -y ttf-wqy-microhei

COPY --from=builder /go/bin/Hackathon2018_API .

COPY --from=builder /go/src/github.com/yuetsh/Hackathon2018_API/templates .

EXPOSE 3010

ENTRYPOINT ["./Hackathon2018_API"]
