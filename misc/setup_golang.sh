#!/bin/sh

mkdir -p ~/projects/{bin,pkg,src}

curl -LO https://golang.org/dl/go1.16.2.linux-amd64.tar.gz

sudo tar -C /usr/local -xvzf go1.16.2.linux-amd64.tar.gz

//sudo vi /etc/profile.d/path.sh
//export PATH=$PATH:/usr/local/go/bin

//vi ~/.bash_profile
//export GOBIN="$HOME/projects/bin"
//export GOPATH="$HOME/projects/src"

//go install $GOPATH/hello.go
