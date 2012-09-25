#!/usr/bin/bash

apt-get update
apt-get -y install python-dev git-core python-setuptools libevent-dev vim
easy_install pip
pip install virtualenv
