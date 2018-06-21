FROM golang:alpine

RUN adduser -D dev

USER dev

RUN mkdir $GOPATH/src/server

RUN mkdir $GOPATH/src/temp &&mkdir $HOME/app/

WORKDIR $HOME/app/

ADD . $HOME/app/


RUN ls&& pwd&&whoami

RUN export PATH=$PATH:$GOROOT/bin

VOLUME $GOPATH/src/server


CMD ["/home/dev/app/cron"]
