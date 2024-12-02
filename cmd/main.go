package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/curioswitch/go-reassign/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
