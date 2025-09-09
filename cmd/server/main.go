// Copyright (c) 2024 Go Project Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

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
