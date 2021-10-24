package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module = "github.com/diamondburned/gotk4/pkg"
	thisModule  = "github.com/diamondburned/gotk4-handy/pkg"
)

var packages = []gendata.Package{
	{PkgName: "libhandy-1", Namespaces: nil},
}

var pkgGenerated = []string{
	"handy",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
	// "_examples",
}

var preprocessors = []types.Preprocessor{
	// types.PreprocessorFunc(func(repos gir.Repositories) {
	// 	layershell := repos.FromPkg("gtk-layer-shell-0")
	// 	layershell.Packages = append(layershell.Packages, gir.Package{
	// 		Name: "gtk+-3.0",
	// 	})
	// }),
}

var filters = []types.FilterMatcher{}
