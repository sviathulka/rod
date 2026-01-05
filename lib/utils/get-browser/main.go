// Package main ...
package main

import (
	"fmt"

	"github.com/sviathulka/rod/lib/launcher"
	"github.com/sviathulka/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
