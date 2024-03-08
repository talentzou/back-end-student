package main

import (
	"back-end/core"
)
func main() {
    r:=core.RunWindowServer()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
