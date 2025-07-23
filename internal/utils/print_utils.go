package utils

import (
	"fmt"
	"strings"
)

func PrintProgress(progress int) {
	fmt.Printf("\r[%s%s] %d%%", strings.Repeat("=", progress/10), strings.Repeat("-", 10-progress/10), progress)
}
