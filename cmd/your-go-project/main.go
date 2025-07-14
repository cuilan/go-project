//go:build (darwin && amd64) || (linux && amd64) || (darwin && arm64)

package main

import (
	"log"
)

func main() {
	if err := RunApp(nil); err != nil {
		log.Fatalf("Application run error: %v", err)
	}
}
