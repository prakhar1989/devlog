FROM ubuntu:16.04

# Install dependencies
RUN apt-get update
RUN apt-get install -y nginx wget supervisor gccgo nodejs npm g++
RUN npm install npm -g
RUN apt-get upgrade -y

# Install node stuff
RUN ln -s /usr/bin/nodejs /usr/bin/node
ENV PATH=/usr/bin/:$PATH

# Install hugo
RUN cd /tmp/ && wget https://github.com/spf13/hugo/releases/download/v0.16/hugo_0.16_linux-64bit.tgz && tar -xzf /tmp/*.tgz && cp /tmp/hugo /usr/bin/

# Install influx
RUN wget https://s3.amazonaws.com/influxdb/influxdb_0.6.5_amd64.deb -O /tmp/influxdb.deb && dpkg -i /tmp/influxdb.deb && rm /tmp/influxdb.deb

# Setup stuff to serve
RUN mkdir /dynamic
RUN mkdir /work/
COPY config.yaml Gruntfile.js nginx.conf package.json /work/
COPY ./content /work/content/
COPY ./layouts /work/layouts/
COPY ./static /work/static/
COPY ./dynamic /work/dynamic/
RUN cd /work/dynamic/ambience && gccgo -g main.go -o /dynamic/ambience
RUN cd /work/dynamic/getip && gccgo -g main.go -o /dynamic/getip
#ADD dynamic/ambience/ambience /dynamic/ambience
#ADD dynamic/getip/getip /dynamic/getip
RUN cd /work/ && npm install grunt execSync
RUN cd /work/ && ./node_modules/grunt/bin/grunt generate

# Setup supervisord
COPY supervisord.conf /etc/supervisor/supervisord.conf

# Application specific dependencies
ADD nginx.conf /etc/nginx/nginx.conf

CMD supervisord -n

# Weather port
EXPOSE 8084
