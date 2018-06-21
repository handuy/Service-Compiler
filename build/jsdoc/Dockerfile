FROM node:alpine

ENV WDIR /home/dev/app

RUN adduser -D dev

USER dev

RUN mkdir $WDIR && mkdir -p /home/dev/temp

WORKDIR $WDIR

ADD . $WDIR

CMD ["/home/dev/app/cron","-path=/home/dev/temp"]
