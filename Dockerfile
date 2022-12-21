FROM ubuntu:20.04

RUN apt update
RUN apt upgrade -y

RUN mkdir /odh-data-monitor
WORKDIR /odh-data-monitor

COPY release/odh-data-monitor /odh-data-monitor/odh-data-monitor

ENTRYPOINT ["/odh-data-monitor/odh-data-monitor"]