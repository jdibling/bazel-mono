package main

import (
	"log"

	"github.com/jdibling/bazel-mono/src/greetings"
)

func main() {
	log.Printf("%s, World!", greetings.Greeting())
}
