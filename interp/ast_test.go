package interp

import (
	"fmt"
	"testing"
)

func TestAst(t *testing.T) {
	src := `
package main

func main() {
	println(1)
}
`
	root, _ := Ast(src, nil)
	fmt.Println(root.index)
}
