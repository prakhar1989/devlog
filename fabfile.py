#!/usr/bin/env python

from fabric.api import task, run, put
from fabric.operations import local

@task()
def deploy():
    run('cd /work/devlog/ && git pull')
    run('cd /work/devlog/ && docker build -t devlog .')

@task()
def restart():
    output = run("docker ps | grep 8084 | awk '{print$1}'")
    run('docker stop %s' % output)
