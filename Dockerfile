FROM golang:alpine AS build-compiler

ENV HOST 0.0.0.0

ENV WDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN mkdir -p $WDIR

ADD . $WDIR

WORKDIR $WDIR

RUN go build -o complier ./cmd/run

FROM minhcuong/alpine-consul

RUN mkdir -p /app/build/
RUN mkdir /app/temp

ADD root /

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/root /app/root/

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/complier /app/

RUN apk update && apk add docker

RUN rc-update add docker boot

RUN service docker start


EXPOSE 8888
