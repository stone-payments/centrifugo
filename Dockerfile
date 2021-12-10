FROM ubuntu:latest

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update && apt install build-essential curl unzip git -y

ENV VERSION 1.7.4

WORKDIR /centrifugo

ENV CENTRIFUGO_SHA256 7a895e97ceee1064ebbff0bfedcab78dd424166a9000ccd6c3e710eb88c5c5ec

ENV DOWNLOAD https://github.com/centrifugal/centrifugo/releases/download/v$VERSION/centrifugo-$VERSION-linux-amd64.zip

RUN curl -sSL "$DOWNLOAD" -o /tmp/centrifugo.zip && \
    echo "${CENTRIFUGO_SHA256}  /tmp/centrifugo.zip" | sha256sum -c - && \
    unzip -jo /tmp/centrifugo.zip -d /tmp/ && \
    mv /tmp/centrifugo /usr/bin/centrifugo && \
    rm -f /tmp/centrifugo.zip && \
    echo "centrifugo - nofile 65536" >> /etc/security/limits.d/centrifugo.nofiles.conf

RUN groupadd -r centrifugo && useradd -r -g centrifugo centrifugo

RUN chown centrifugo:centrifugo /centrifugo && \
    mkdir /var/log/centrifugo && chown centrifugo:centrifugo /var/log/centrifugo

VOLUME ["/centrifugo", "/var/log/centrifugo"]

WORKDIR /centrifugo

COPY config.sh .
COPY Makefile .

RUN ls -la

USER centrifugo

CMD ["make", "runcentrifugo"]

EXPOSE 8888

