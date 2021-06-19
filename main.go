package main

import (
	_ "alfred/boot"
	_ "alfred/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
