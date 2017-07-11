FROM golang:1.7-alpine

RUN apk update && apk add git
RUN go-wrapper download github.com/tools/godep
RUN go-wrapper install github.com/tools/godep

#RUN go-wrapper download github.com/Masterminds/glide
#RUN go-wrapper install github.com/Masterminds/glide
