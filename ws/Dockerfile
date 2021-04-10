FROM ubuntu:20.04

ARG USERID=0

RUN apt-get update \
&& DEBIAN_FRONTEND=noninteractive apt-get install -y \
nodejs \
npm \
&& npm install -g yarn \
&& apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
&& adduser --disabled-password --gecos "" --uid $USERID app

WORKDIR /opt/app/ws

USER app

CMD ["/usr/bin/npm", "run", "dev"]
