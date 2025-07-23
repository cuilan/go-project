package utils

import (
	"go-project/internal/utils"
	"testing"
	"time"
)

func TestPrintProgress(t *testing.T) {
	for i := 0; i < 100; i++ {
		utils.PrintProgress(i)
		time.Sleep(1000 * time.Millisecond)
	}
}
