FROM golang:alpine AS builder

RUN apk add --no-cache git

RUN go get github.com/lib/pq

RUN go get github.com/pilu/fresh

ARG project=/go/src/github.com/yuetsh/Hackathon2018

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
