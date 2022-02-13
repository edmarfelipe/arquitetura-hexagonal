package main

import (
	"github.com/edmarfelipe/go-hexagonal/adapters/httpserver"
)

func main() {
	// cmd.ExecuteCli()
	httpserver.MakeNewServcer().Serve(":3000")
}
