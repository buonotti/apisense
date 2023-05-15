FROM ubuntu:20.04

# get the branch to use
ARG BRANCH

# update system and install required software
RUN apt-get update
RUN apt-get install curl tar supervisor -y

# set workdir
WORKDIR /

# install the go tools
RUN curl -sLO https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.19.4.linux-amd64.tar.gz

RUN mkdir /apisense
COPY . /apisense

WORKDIR /apisense

# build the project install it and add it to path
ENV PATH="$PATH:/root/go/bin:/usr/local/go/bin"
RUN go build
RUN go install

# create app directories
RUN mkdir -p /root/apisense
RUN mkdir -p /root/.config/apisense

COPY examples/apisense.config.yml /root/.config/apisense/apisense.config.yml
COPY examples/bluetooth.apisensedef.yml /root/apisense/definitions/bluetooth.apisensedef.yml

# copy supervisor config
COPY docker/ssh.supervisor.conf /etc/supervisor/conf.d/ssh.supervisor.conf
COPY docker/api.supervisor.conf /etc/supervisor/conf.d/api.supervisor.conf
COPY docker/daemon.supervisor.conf /etc/supervisor/conf.d/daemon.supervisor.conf

COPY docker/startup.sh /startup.sh

# expose ssh and api port
EXPOSE 23232
EXPOSE 8080

WORKDIR /root

# start supervisord
ENTRYPOINT ["sh", "/startup.sh"]