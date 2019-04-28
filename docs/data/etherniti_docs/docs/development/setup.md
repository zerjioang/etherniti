# Environment setup

Follow this guide to have your enviornment ready for etherniti

## Requirements

Following, etherniti requirements are listed:

* Linux like computer. If you are using Windows, hurry up and grab a Linux powered laptop.
* Have access to an Ubuntu 18.04 system, configured with a non-root user with sudo privileges

## Go installation

### Introduction

Go is a modern programming language developed at Google. It is increasingly popular for many applications and at many companies, and offers a robust set of libraries. This tutorial will walk you through downloading and installing the latest version of Go (Go 1.10 at the time of this article's publication), as well as building a simple Hello World application.

If you dont have **Go** already installed (`go version`), please install it!

```bash
curl -O https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar -xvf go1.11.5.linux-amd64.tar.gz
sudo mv go /usr/local
```
### Testing your go installation

Now that Go is installed and the paths are set for your machine, you can test to ensure that Go is working as expected.

Easy and simplest way: type

```bash
go version
```

and it should print the installed go version

```bash
go version go1.11 linux/amd64
```

!!! note 
	If you have troubles installing Go, please visit one of the many tutorials available in Internet. here is one, provided by [DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04)

## Etherniti sources download

### Clone the repository from Gitlab

Download the latest version of the Go project from Gitlab:

```bash
git clone git@github.com:etherniti etherniti
```

or if you prefer HTTPS:

```bash
git clone https://github.com/etherniti etherniti
```

After a successful clone you will see something similar to:

```bash
Clonando en 'etherniti'...

UNAUTHORIZED ACCESS TO THIS DEVICE IS PROHIBITED
You must have explicit, authorized permission to access or configure this device.
Unauthorized attempts and actions to access or use this system may result in civil and/or 
criminal penalties.
All activities performed on this device are logged and monitored.

remote: Counting objects: 1045, done.
remote: Compressing objects: 100% (307/307), done.
remote: Total 1045 (delta 312), reused 315 (delta 197)
Recibiendo objetos: 100% (1045/1045), 283.17 KiB | 5.45 MiB/s, listo.
Resolviendo deltas: 100% (561/561), listo.
```

## Importing cloned project in Goland

After that, go to created folder called	`etherniti` via `cd etherniti` and there should be all source files of the go contract.

### GOPATH check

In order to open this project successfully, the recommendation is to place `etherniti` files under your `GOPATH` which usually will be located under `$HOME/go`. To make sure where your `GOPATH` is run:

```bash
echo $GOPATH
```

If the result is empty, that means you dont have the `GOPATH` set.

### Setting `GOPATH` in your system

Edit your `~/.bash_profile` to add the following line:

```bash
export GOPATH=$HOME/go # or use your prefered gopath
```

Save and exit your editor. Then, source your `~/.bash_profile` to commit the changes.

```bash
source ~/.bash_profile
```

!!! Note "A tip for ZSH terminal users"
    If you are using **ZSH** or any other non-native terminal, you should do the same process but with that terminal configuration file.
