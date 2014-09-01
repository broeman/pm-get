pm-get - Package Manager in Go 
==============================

## User Application
BSD-License applies, see LICENSE for more information.

### Purpose
The idea is to get it to install source packages, defined in a database, with install-scripts based on LinuxFromScratch.

### Overview
- [Documentation](http://godoc.org/github.com/broeman/pm-get)

### Current Version 0.1b
Can install, uninstall and show package.

### Roadmap to 0.1
- Update: update packages that is marked 'stable'
- Install-script available for testing

### Installation
GoPack is the library used, pm-tools currently initializes the db, so quite useful for testing:

```
$ go get github.com/broeman/gopack
$ go get github.com/broeman/pm-tools
$ go get github.com/broeman/pm-get
```

Make sure that you have set $GOPATH and source it in your PATH (e.g. ~/.bashrc):
```
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### About
Go Packge is written by [Jesper Brodersen](http://jesperbrodersen.dk)
