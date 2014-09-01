// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Implementation of CLI command structure
package cmd

import (
	"fmt"
	"github.com/broeman/gopack/db"   // using DB struct for placeholder REMOVE
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

var Installed = cli.Command{
	Name:        "installed",
	Usage:       "Shows installed packages",
	Description: `Shows all currently installed packages`,
	Action:      runInstalled,
	Flags:       []cli.Flag{},
}

var Init = cli.Command{
	Name:        "init",
	Usage:       "Placeholder initialization",
	Description: "Placeholdering",
	Action:      runPlaceholder,
	Flags:       []cli.Flag{},
}

// Init DB placeholder, should be a pm-tools setup function REMOVE
func runPlaceholder(ctx *cli.Context) {
	db.InitDB()
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

// Placeholder to see if things works REMOVE
func runInstalled(*cli.Context) {
	rows, err := db.QueryAllPackages()
	if err != nil {
		fmt.Println("Query Error: : %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var versionregex string
		var installed bool
		rows.Scan(&id, &name, &versionregex, &installed)
		fmt.Println(id, name, installed)
	}
	rows.Close()
}
