FROM ubuntu:16.04

# Install dependencies
RUN apt-get update
RUN apt-get install -y nginx wget supervisor g++
RUN apt-get upgrade -y

# Install node stuff
ENV PATH=/usr/bin/:$PATH

# Install hugo
RUN cd /tmp/ && wget https://github.com/spf13/hugo/releases/download/v0.16/hugo_0.16_linux-64bit.tgz && tar -xzf /tmp/*.tgz && cp /tmp/hugo /usr/bin/

# Setup stuff to serve
RUN mkdir /work/
COPY config.yaml Gruntfile.js nginx.conf package.json /work/
COPY ./content /work/content/
COPY ./layouts /work/layouts/
COPY ./static /work/static/
RUN cd /work/ && /usr/bin/hugo

# Setup supervisord
COPY supervisord.conf /etc/supervisor/supervisord.conf

# Application specific dependencies
ADD nginx.conf /etc/nginx/nginx.conf

CMD supervisord -n

# Weather port
EXPOSE 8084
