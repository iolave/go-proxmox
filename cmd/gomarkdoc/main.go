package main

import (
	"fmt"
	"go/build"
	"os"
	"strings"

	"github.com/princjef/gomarkdoc"
	"github.com/princjef/gomarkdoc/lang"
	"github.com/princjef/gomarkdoc/logger"
)

func main() {
	packages := []string{
		"proxmox_api",
		"cloudflare",
	}

	for i := 0; i < len(packages); i++ {
		pkg := packages[i]
		md, err := getPackageMd(pkg)
		handleErr(err)
		err = writePackageMd(pkg, md)
		handleErr(err)
	}
}

func handleErr(err error) {
	if err == nil {
		return
	}

	fmt.Println("error:", err.Error())
	os.Exit(1)
}

func getFuncString(pkg *lang.Func) string {
	return fmt.Sprint(
		fmt.Sprintf("####"),
	)
}

func writePackageMd(pkgName string, md string) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/docs/reference/pkg/%s.md", wd, pkgName), []byte(md), 0644)
}

func getPackageMd(pkgName string) (string, error) {
	// Create a renderer to output data
	out, err := gomarkdoc.NewRenderer(
		gomarkdoc.WithTemplateOverride("index", ""),
	)

	if err != nil {
		return "", err
	}

	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// sets the project package path
	pkgPath := fmt.Sprintf("%s/pkg/%s", wd, pkgName)

	buildPkg, err := build.ImportDir(pkgPath, build.ImportComment)

	if err != nil {
		return "", err
	}

	// Create a documentation package from the build representation of our
	// package.
	log := logger.New(logger.DebugLevel)
	pkg, err := lang.NewPackageFromBuild(
		log,
		buildPkg,
		lang.PackageWithRepositoryOverrides(&lang.Repo{
			DefaultBranch: "master",
			Remote:        "https://github.com/iolave/go-proxmox",
			PathFromRoot:  "/",
		}),
	)

	if err != nil {
		return "", err
	}

	s, _ := out.Package(pkg)
	s = strings.Replace(s, "## Index", "", 1)

	return s, nil
}
