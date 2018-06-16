FROM golang:alpine

RUN mkdir -p $GOPATH/src/git.hocngay.com/techmaster/service-complier

ADD . $GOPATH/src/git.hocngay.com/techmaster/service-complier

WORKDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN go build -o service-complier .

CMD ["./service-complier"]