#!/bin/bash

#install ganeti deb package
sudo dpkg -i ganeti.deb

sudo mkdir /srv/ganeti/ /srv/ganeti/os /srv/ganeti/export
#sudo cp doc/examples/ganeti.initd /etc/init.d/ganeti

#Setup Ganeti service

sudo update-rc.d ganeti defaults 20 80

sudo mkdir default
sudo ln -s /usr/local/share/ganeti/2.11/ ./default 


