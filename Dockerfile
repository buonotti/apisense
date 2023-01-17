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
RUN git clone -b $BRANCH https://github.com/buonotti/apisense
WORKDIR /apisense

# build the project install it and add it to path
ENV PATH="$PATH:/root/go/bin"
RUN /usr/local/go/bin/go get -u github.com/go-bindata/go-bindata/...
RUN /usr/local/go/bin/go install github.com/go-bindata/go-bindata/...
RUN go-bindata -o assets.go assets/
RUN /usr/local/go/bin/go build
RUN /usr/local/go/bin/go install

# create app directories
RUN mkdir -p /root/.config/apisense
RUN mkdir /root/apisense

# copy supervisor config
COPY docker/ssh.supervisor.conf /etc/supervisor/conf.d/ssh.supervisor.conf
COPY docker/api.supervisor.conf /etc/supervisor/conf.d/api.supervisor.conf

COPY docker/startup.sh /startup.sh

# expose ssh and api port
EXPOSE 23232
EXPOSE 8080

WORKDIR /root

# start supervisord
ENTRYPOINT ["sh", "/startup.sh"]