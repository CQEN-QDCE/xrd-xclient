package main

import (
	"log"
	"github.com/brynjarh/xclient/pkg/web"
)
func main() {
	addr := ":5000"
	log.Printf("Starting server on %s", addr)
	web.StartWWW(addr)
}