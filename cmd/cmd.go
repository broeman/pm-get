// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Implementation of CLI command structure
package cmd

import (
	"fmt"
	"github.com/broeman/gopack/pack" // using Package struct
	"github.com/codegangsta/cli"
)

var Install = cli.Command{
	Name:        "install",
	Usage:       "Installs a package",
	Description: `Installs a package, that isn't already installed`,
	Action:      runInstall,
	Flags:       []cli.Flag{},
}

var UnInstall = cli.Command{
	Name:        "uninstall",
	Usage:       "Uninstalls a package",
	Description: `Uninstalls a package, that is already installed`,
	Action:      runUninstall,
	Flags:       []cli.Flag{},
}

var Show = cli.Command{
	Name:        "show",
	Usage:       "shows a package",
	Description: `Shows a package, that is in the database`,
	Action:      showPackage,
	Flags:       []cli.Flag{},
}

func showPackage(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Println("You need to specify which package you want to query")
		return
	}
	apack := pack.RetrievePackage(ctx.Args().First())
	if apack.Name() == "" {
		fmt.Println("Package not found")
		return
	}
	fmt.Println(apack.Name(), apack.Installed())
}

func runUninstall(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Println("You need to specify which package you want to uninstall")
		return
	}

	curpackage := pack.RetrievePackage(ctx.Args().First())

	if curpackage.Name() == "" {
		fmt.Println("Package not found")
		return
	}

	if !curpackage.Installed() {
		fmt.Println("Package " + curpackage.Name() + " is not installed")
		return
	}

	fmt.Println("Uninstalling package " + curpackage.Name())
	curpackage.SetInstalled(false)
}

func runInstall(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Println("You need to specify which package you want to install")
		return
	}

	curpackage := pack.RetrievePackage(ctx.Args().First())

	if curpackage.Name() == "" {
		fmt.Println("Package not found")
		return
	}

	if curpackage.Installed() {
		fmt.Println("Package " + curpackage.Name() + " is already installed")
		return
	}

	fmt.Println("Installing package " + curpackage.Name())
	curpackage.SetInstalled(true)
}
