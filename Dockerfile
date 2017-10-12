FROM golang:1.7-alpine

#ENV GIN_MODE=release
# https://github.com/alecthomas/gometalinter/issues/149
ENV CGO_ENABLED=0

RUN apk update && apk add git
COPY . $GOPATH/src/gitlab.com/MikaXII/recalbox-api
RUN go get -u github.com/tools/godep
RUN go get -u github.com/alecthomas/gometalinter
RUN gometalinter --install --update

WORKDIR $GOPATH/src/gitlab.com/MikaXII/recalbox-api
RUN godep restore -v
RUN ls
# Disable gotype -> give an error
RUN gometalinter --aggregate . --disable gotype 
RUN go install
RUN mkdir /etc/recalbox-api
RUN cp ./config/config.toml /etc/recalbox-api/

CMD ["recalbox-api"]
