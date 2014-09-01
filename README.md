pm-get - Package Manager in Go 
==============================

## User Application
BSD-License applies, see LICENSE for more information.

### Current Version: 0.1a
Running in alpha, and not that useful. At this stage it is just me tinkering. Haven't even acted professional yet.

### Purpose
The idea is to get it to install source packages, defined in a database, with install-scripts based on LinuxFromScratch.

### Overview
- [Documentation](http://godoc.org/github.com/broeman/pm-get)

### Installation
You shouldn't, but do it the go way:

```
$ go get github.com/broeman/gopack
$ go get github.com/broeman/pm-get
```

Make sure that you have set $GOPATH and source it in your PATH (e.g. ~/.bashrc):
```
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### About
Go Packge is written by [Jesper Brodersen](http://jesperbrodersen.dk)
