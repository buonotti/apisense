FROM ubuntu:20.04

RUN apt update
RUN apt upgrade -y
RUN apt install git curl tar -y

WORKDIR /

RUN curl -OL https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.19.4.linux-amd64.tar.gz

COPY ./ /odh-data-monitor

WORKDIR /odh-data-monitor

RUN /usr/local/go/bin/go build

RUN /usr/local/go/bin/go install

ENV PATH="$PATH:/root/go/bin"

RUN mkdir -p /root/.config/odh-data-monitor

EXPOSE 23232

ENTRYPOINT ["odh-data-monitor", "ssh"]