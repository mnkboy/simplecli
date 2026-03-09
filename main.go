package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Operating System: %s\n Architecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
