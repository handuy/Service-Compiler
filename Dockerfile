FROM ubuntu:16.04

RUN apt-get -y update&& apt-get -y upgrade

RUN chmod +x ./installdocker.sh && chmod +x ./installgo.sh

