FROM ubuntu:20.04

ARG USERID=0

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
python3-pip \
vim-tiny \
&& apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
&& adduser --disabled-password --gecos "" --uid $USERID app

WORKDIR /opt/app/api

COPY requirements.txt /opt/app/requirements

RUN pip3 install -r /opt/app/requirements

USER app

CMD [ "uvicorn","main:app","--reload","--host", "0.0.0.0" ]
