package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.Run("go", "build", "-o", "build/reassign", "./cmd")
}

func Test() error {
	return sh.RunV("go", "test", "./...")
}

func Format() error {
	return sh.RunV("go", "run", fmt.Sprintf("github.com/rinchsan/gosimports/cmd/gosimports@%s", gosImportsVer), "-w", ".")
}

func Check() error {
	return sh.RunV("go", "run", fmt.Sprintf("github.com/golangci/golangci-lint/cmd/golangci-lint@%s", golangCILintVer), "run")
}

var Default = Build
