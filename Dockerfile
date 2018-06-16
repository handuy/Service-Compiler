FROM golang:alpine

RUN mkdir -p $GOPATH/src/git.hocngay.com/techmaster/service-complier

ADD . $GOPATH/src/git.hocngay.com/techmaster/service-complier

WORKDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN go build -o service-complier .

RUN apk add docker --repository http://dl-cdn.alpinelinux.org/alpine/latest-stable/community &&apk update

RUN rc-update add docker boot

RUN service docker start


CMD ["./service-complier"]