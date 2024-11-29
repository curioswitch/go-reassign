package main

import (
	"fmt"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/boot"
	"github.com/goyek/x/cmd"

	"github.com/curioswitch/go-build"
)

func main() {
	build.DefineTasks()

	goyek.Define(goyek.Task{
		Name: "snapshot",
		Action: func(a *goyek.A) {
			cmd.Exec(a, fmt.Sprintf("go run github.com/goreleaser/goreleaser@%s release --snapshot --rm-dist", verGoReleaser))
		},
	})

	goyek.Define(goyek.Task{
		Name: "releasea",
		Action: func(a *goyek.A) {
			cmd.Exec(a, fmt.Sprintf("go run github.com/goreleaser/goreleaser@%s release --rm-dist", verGoReleaser))
		},
	})

	boot.Main()
}
