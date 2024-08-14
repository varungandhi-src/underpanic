package main

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func main() {
	mode := packages.NeedTypes | packages.NeedTypesInfo | packages.NeedName
	cfg := &packages.Config{Mode: mode}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue
		}

		for _, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}
			obj, ok := obj.(*types.TypeName)
			if !ok {
				continue
			}
			objType, ok := obj.Type().(*types.Named)
			if !ok {
				continue
			}
			fmt.Printf("processing type: %s\n", obj.String())
			_ = types.NewMethodSet(objType)
		}
	}
}
