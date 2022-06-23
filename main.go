package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/curioswitch/reassign/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
