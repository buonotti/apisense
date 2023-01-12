FROM ubuntu:20.04

# get the branch to use
ARG BRANCH

# update system and install required software
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install git curl tar supervisor zsh -y

# set workdir
WORKDIR /

# install the go tools
RUN curl -OL https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.19.4.linux-amd64.tar.gz

# clone the project and cd into it
RUN git clone -b $BRANCH https://github.com/buonotti/odh-data-monitor
WORKDIR /odh-data-monitor

# build the project install it and add it to path
RUN /usr/local/go/bin/go build
RUN /usr/local/go/bin/go install
ENV PATH="$PATH:/root/go/bin"

# create app directories
RUN mkdir -p /root/.config/odh-data-monitor
RUN mkdir /root/odh-data-monitor

# copy supervisor config
COPY docker/ssh.supervisor.conf /etc/supervisor/conf.d/ssh.supervisor.conf
COPY docker/api.supervisor.conf /etc/supervisor/conf.d/api.supervisor.conf

COPY docker/startup.sh /startup.sh

RUN mkdir /comp
RUN odh-data-monitor completion zsh > /comp/_odh-data-monitor
RUN chmod +x /comp/_odh-data-monitor
ENV FPATH="$FPATH:/comp"

# expose ssh and api port
EXPOSE 23232
EXPOSE 8080

# start supervisord
ENTRYPOINT ["sh", "/startup.sh"]