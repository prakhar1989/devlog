FROM ubuntu:14.04

# Install dependencies
RUN apt-get update &&apt-get upgrade -y
RUN apt-get install -y nginx wget supervisor gccgo node npm

# Install node stuff
RUN ln -s /usr/bin/nodejs /usr/bin/node
ENV PATH=/usr/bin/:$PATH

# Install hugo
RUN cd /tmp/ && wget https://github.com/spf13/hugo/releases/download/v0.13/hugo_0.13_linux_amd64.tar.gz && tar -xzf /tmp/*.tar.gz && cd /usr/bin/ && ln -s /tmp/hugo_0.13_linux_amd64/hugo_0.13_linux_amd64 hugo

# Install influx
RUN wget https://s3.amazonaws.com/influxdb/influxdb_0.6.5_amd64.deb -O /tmp/influxdb.deb && dpkg -i /tmp/influxdb.deb && rm /tmp/influxdb.deb

# Setup stuff to serve
RUN mkdir /dynamic
RUN mkdir /work/
COPY config.yaml Gruntfile.js nginx.conf package.json /work/
COPY ./content /work/content/
COPY ./layouts /work/layouts/
COPY ./static /work/static/
ADD dynamic/ambience/ambience /dynamic/ambience
ADD dynamic/getip/getip /dynamic/getip
RUN cd /work/ && npm install grunt execSync && ./node_modules/grunt/bin/grunt generate

# Setup supervisord
COPY supervisord.conf /etc/supervisor/supervisord.conf

# Application specific dependencies
ADD nginx.conf /etc/nginx/nginx.conf

CMD service supervisor start

# Weather port
EXPOSE 8084
