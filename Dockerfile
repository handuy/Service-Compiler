FROM golang:alpine

RUN mkdir -p $GOPATH/src/git.hocngay.com/techmaster/service-complier

ADD . $GOPATH/src/git.hocngay.com/techmaster/service-complier

WORKDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN go build -o service-complier .

RUN apk install sudo

RUN chmod +x ./installdocker.sh && ./installdocker.sh

CMD ["./service-complier"]