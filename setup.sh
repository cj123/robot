#!/bin/bash
# get servod from 4tronix

wget -q http://4tronix.co.uk/initio/servod.xxx -O servod
chmod +x servod
sudo mv servod /usr/bin/servod