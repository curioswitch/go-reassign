package shadow

import "fmt"

func use() {
	fmt.Println("hi")
}

func shadow() any {
	fmt := struct{ EOF int }{}
	fmt.EOF = 5 // want "foo"
	return fmt
}
