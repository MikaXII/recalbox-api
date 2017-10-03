FROM golang:1.7-alpine

#ENV GIN_MODE=release

RUN apk update && apk add git
COPY . /go/src/gitlab.com/MikaXII/recalbox-api
RUN go get -u github.com/tools/godep

WORKDIR /go/src/gitlab.com/MikaXII/recalbox-api
RUN godep restore
RUN go install
RUN mkdir /etc/recalbox-api
RUN cp ./config/config.toml /etc/recalbox-api/

CMD ["recalbox-api"]
