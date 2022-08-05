package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

var gosImports = sh.RunCmd("go", "run", fmt.Sprintf("github.com/rinchsan/gosimports/cmd/gosimports@%s", gosImportsVer))

func Build() error {
	return sh.Run("go", "build", "-o", "build/reassign", ".")
}

func Test() error {
	return sh.RunV("go", "test", "./...")
}

func Format() error {
	return gosImports("-w", ".")
}

var Default = Build
