FROM node:alpine

RUN mkdir $HOME/app

RUN mkdir $HOME/temp

ADD . $HOME/app

WORKDIR $HOME/app

CMD ["node","run.js"]
