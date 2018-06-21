FROM golang:alpine AS build-compiler

ENV HOST 0.0.0.0

ENV WDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN mkdir -p $WDIR

ADD . $WDIR

WORKDIR $WDIR

RUN go build -o cron ./cmd/cron

RUN go build -o complier ./cmd/run

FROM minhcuong/alpine-consul

RUN mkdir -p /app/build/ && mkdir -p /app/build/ && mkdir /app/temp

ADD root /

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/build /app/build/

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/cron /app/

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/root /app/root/

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-complier/complier /app/

RUN chmod +x /app/complier

RUN apk update && apk add docker && apk add openrc

RUN rc-update add docker boot


EXPOSE 8888
