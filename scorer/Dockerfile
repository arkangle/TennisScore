FROM golang:1.14-buster
ARG USERID=0
RUN adduser --disabled-password --gecos "" --uid $USERID app
WORKDIR /opt/app/scorer
USER app
CMD [ "go","run","cmd/kafkanated/main.go" ]