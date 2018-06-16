FROM golang:alpine

RUN mkdir -p $GOPATH/src/git.hocngay.com/techmaster/service-complier

ADD . $GOPATH/src/git.hocngay.com/techmaster/service-complier

WORKDIR $GOPATH/src/git.hocngay.com/techmaster/service-complier

RUN go build -o service-complier .

RUN sed -e 's;^#http\(.*\)/v3.6/community;http\1/v3.6/community;g' -i /etc/apk/repositories
RUN apk update
RUN apk add docker

RUN memb=$(grep "^docker:" /etc/group | sed -e 's/^.*:\([^:]*\)$/\1/g')[ "${memb}x" = "x" ] && memb=${USER} || memb="${memb},${USER}"
RUN sed -e "s/^docker:\(.*\):\([^:]*\)$/docker:\1:${memb}/g" -i /etc/group

RUN rc-update add docker

CMD ["./service-complier"]