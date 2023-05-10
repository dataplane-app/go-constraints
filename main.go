package main

import (
	"hello/functions"
	"log"
	"runtime"
)

func main() {
	log.Println(runtime.GOOS)
	functions.Hello()
}
