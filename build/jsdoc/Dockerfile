FROM node:alpine

RUN mkdir $HOME/app

RUN mkdir $HOME/temp

ADD . $HOME/app

WORKDIR $HOME/app

RUN adduser -D dev

USER dev

CMD ["node","run.js"]
