FROM golang:alpine AS builder

RUN apk add --no-cache git

ARG project=/go/src/github.com/yuetsh/Hackathon2018

WORKDIR ${project}

RUN go get github.com/lib/pq

ADD src src

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o release src/main.go

FROM scratch

ARG project=/go/src/github.com/yuetsh/Hackathon2018

COPY --from=builder ${project}/release /

EXPOSE 3000

ENTRYPOINT ["/release"]
