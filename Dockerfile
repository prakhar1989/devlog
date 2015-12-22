FROM ubuntu:14.04

ADD public/ srv/
RUN mkdir /dynamic
ADD dynamic/ambience/ambience /dynamic/ambience

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y nginx
ADD nginx.conf /etc/nginx/nginx.conf
RUN apt-get install -y wget
RUN wget https://s3.amazonaws.com/influxdb/influxdb_0.6.5_amd64.deb -O /tmp/influxdb.deb
RUN dpkg -i /tmp/influxdb.deb
RUN rm /tmp/influxdb.deb
RUN apt-get install -y gccgo

# Weather port
EXPOSE 8085
