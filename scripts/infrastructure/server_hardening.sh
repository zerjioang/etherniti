#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "installing sudo"
apt install sudo -y

echo "creating non root user"
adduser eth
usermod -aG sudo eth
su - eth
whoami
sudo whoami

echo "printing system information"
lsb_release -a
echo "cpu information"
lscpu
echo "hardware information"
lshw
echo "pci information"
lcpci

echo "checking CPU NUMA status"
sudo apt install numactl -y
numactl --hardware

echo "printing hard disk information"
echo "ROTA = 1 means HDD, ROTA = 0 means SDD"
lsblk -d -o name,rota

echo "updating"
sudo apt update
sudo apt upgrade -y

echo "finding installed packages"
apt list --installed

echo "checking current server entropy"
cat /proc/sys/kernel/random/entropy_avail

echo "installing haveged"
sudo apt install haveged rng-tools -y

echo "benchmarking server entropy"
cat /dev/random | rngtest -c 1000

echo "checking current server entropy again"
cat /proc/sys/kernel/random/entropy_avail

echo "checking docker compatibility"
$ curl https://raw.githubusercontent.com/docker/docker/master/contrib/check-config.sh > check-config.sh
bash check-config.sh
rm check-config.sh

echo "installing docker pre-requisites"
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update

echo "removing old dependencies"
sudo apt-get remove docker docker-engine docker.io containerd runc
echo "installing docker"
sudo apt-get install docker docker-compose -y
echo "checking docker status"
systemctl status docker.service

echo "post docker installation steps"
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
mkdir $HOME/.docker
sudo chown "$USER":"$USER" /home/$USER/.docker -R
sudo chmod g+rwx "$HOME/.docker" -R
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker.service

echo "installing required software"
sudo apt install tmux git

echo "installing low latency kernel"
sudo apt-get install linux-headers-lowlatency
sudo apt-get install linux-lowlatency
sudo update-grub

echo "checking kernel parameters"
sysctl -a 

echo "checking system ulimit"
limit=$(ulimit -n)
if  [ $x -le 1024 ]; then
	echo "server ulimit is too small"
	echo "updating file file /etc/security/limits.conf"
	limitfile="/etc/security/limits.conf"
	echo "*    soft nofile 1048576" > $limitfile
	echo "*    hard nofile 1048576" > $limitfile
	echo "root soft nofile 1048576" > $limitfile
	echo "root hard nofile 1048576" > $limitfile
fi

echo "session limits"
echo "session required pam_limits.so" > /etc/pam.d/common-session

echo "sysctl.conf"
echo "fs.file-max = 2097152" > /etc/sysctl.conf
echo "fs.nr_open = 1048576" > /etc/sysctl.conf
echo "" > /etc/sysctl.conf
echo "net.ipv4.netfilter.ip_conntrack_max = 1048576" > /etc/sysctl.conf
echo "net.nf_conntrack_max = 1048576" > /etc/sysctl.conf
echo "net.core.somaxconn = 1048576" > /etc/sysctl.conf

echo "increasing your network stack buffers"

net.core.rmem_default = 10000000
net.core.wmem_default = 10000000
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216

echo "removing orphaned packages using deborphan"
sudo apt-get install deborphan -y
deborphan
sudo orphaner

echo "removing system fonts"
sudo apt remove fonts-*

sudo apt-get autoclean
sudo apt-get clean all
sudo apt-get autoremove
sudo apt-get autoremove --purge