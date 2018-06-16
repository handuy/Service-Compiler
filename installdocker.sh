#!/bin/sh

set -e

sed -e 's;^#http\(.*\)/v3.6/community;http\1/v3.6/community;g' \
     -i /etc/apk/repositories
apk update
apk add docker

memb=$(grep "^docker:" /etc/group | sed -e 's/^.*:\([^:]*\)$/\1/g')
[ "${memb}x" = "x" ] && memb=${USER} || memb="${memb},${USER}"

sed -e "s/^docker:\(.*\):\([^:]*\)$/docker:\1:${memb}/g" -i /etc/group

rc-update add docker
