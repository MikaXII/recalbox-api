FROM golang:1.7-alpine

RUN apk update && apk add git
 COPY . /go/src/gitlab.com/MikaXII/recalbox-api
RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/gitlab.com/MikaXII/recalbox-api
RUN govendor sync


CMD ["go", "run", "main.go"]
