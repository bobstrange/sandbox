#!/bin/bash

WEBDIR=./www/docs
CONFDIR=./httpd/conf.d
TEMPLATE=./vh_template.txt

[[ -d ${CONFDIR} ]] || mkdir -p ${CONFDIR}

sed s/dumnmy-host.example.com/"$1"/ ${TEMPLATE} > "${CONFDIR}/${1}.conf"
mkdir -p "${WEBDIR}/$1"
echo "New site for $1 > ${WEBDIR}/$1/index.html"
