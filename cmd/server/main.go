package main

import (
	"go-project/internal/utils"
	"time"
)

func main() {
	// fmt.Println("Hello, World!")
	for i := 1; i <= 100; i++ {
		utils.PrintProgress(i)
		time.Sleep(10 * time.Millisecond)
	}
}
