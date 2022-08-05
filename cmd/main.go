package main

import (
	"github.com/curioswitch/go-noreassign/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.New())
}
