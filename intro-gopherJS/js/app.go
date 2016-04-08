package main

import (
	"github.com/gopherjs/gopherjs/js"
	"fmt"
)

func main() {
	js.Global.Get("document").Call("write", "Say hello to GopherJS")
	fmt.Println("GopherJS")
}
