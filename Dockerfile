FROM ubuntu:14.04

# Setup stuff to serve
ADD public/ srv/
RUN mkdir /dynamic
ADD dynamic/ambience/ambience /dynamic/ambience
ADD dynamic/getip/getip /dynamic/getip

# Install dependencies
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y nginx
RUN apt-get install -y wget

# Setup supervisord
RUN apt-get install -y supervisor
COPY supervisord.conf /etc/supervisor/supervisord.conf


# Application specific dependencies
RUN apt-get install -y gccgo
RUN wget https://s3.amazonaws.com/influxdb/influxdb_0.6.5_amd64.deb -O /tmp/influxdb.deb
RUN dpkg -i /tmp/influxdb.deb
RUN rm /tmp/influxdb.deb
ADD nginx.conf /etc/nginx/nginx.conf

CMD service supervisor start

# Weather port
EXPOSE 8084
