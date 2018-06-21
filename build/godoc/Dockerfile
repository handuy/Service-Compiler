FROM golang:alpine

ENV WDIR /home/dev/app

RUN adduser -D dev

USER dev

RUN mkdir $WDIR

RUN mkdir $GOPATH/src/temp

WORKDIR $WDIR

ADD . $WDIR

RUN export PATH=$PATH:$GOROOT/bin

CMD ["/home/dev/app/cron","-path=/go/src/temp","-user=dev"]
