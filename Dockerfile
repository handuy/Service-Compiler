FROM golang:alpine AS build-compiler

ENV HOST 0.0.0.0

ENV WDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier


RUN mkdir -p $WDIR

ADD . $WDIR

WORKDIR $WDIR

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN git checkout development

RUN go build -o compiler .

FROM minhcuong/alpine-consul

# Add s6 service
ADD root /

RUN mkdir -p /app/data/
RUN mkdir /app/temp

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-compiler/build /app/build/

COPY --from=build-compiler /go/src/git.hocngay.com/techmaster/service-compiler/compiler /app/

RUN apk update && apk add docker

RUN rc-update add docker boot

RUN service docker start


EXPOSE 8888
