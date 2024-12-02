package main

import (
	"fmt"

	"github.com/curioswitch/go-build"
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/boot"
	"github.com/goyek/x/cmd"
)

func main() {
	build.DefineTasks()

	goyek.Define(goyek.Task{
		Name: "snapshot",
		Action: func(a *goyek.A) {
			cmd.Exec(a, fmt.Sprintf("go run github.com/goreleaser/goreleaser/v2@%s release --snapshot --clean", verGoReleaser))
		},
	})

	goyek.Define(goyek.Task{
		Name: "releasea",
		Action: func(a *goyek.A) {
			cmd.Exec(a, fmt.Sprintf("go run github.com/goreleaser/goreleaser/v2@%s release --clean", verGoReleaser))
		},
	})

	boot.Main()
}
