FROM golang:alpine

RUN mkdir $GOPATH/src/server

RUN mkdir $GOPATH/src/temp

ADD . $GOPATH/src/server

WORKDIR $GOPATH/src/server

RUN go build -o svc .

RUN export PATH=$PATH:$GOROOT/bin

VOLUME $GOPATH/src/server

CMD ["./svc"]
