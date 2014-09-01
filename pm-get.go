// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// pm-get is a package manger written in Go Language
// this is the user application for running daily tasks on packages,
// like browsing, installing, uninstalling and updating
package main

import (
	"github.com/broeman/gopack/cmd" // using CLI command args
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

const APP_VER = "0.1 Alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "pm-get"
	app.Usage = "Package Manager in Go"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Install,   // install a package
		cmd.UnInstall, // uninstall a package
		cmd.Show,      // show package
		cmd.Installed, // shows current installed packages, placeholder
		//cmd.Update,	// update packages
		cmd.Init, // placeholder initialization
	}
	app.Run(os.Args)

}
