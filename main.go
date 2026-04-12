package main

import (
	"fmt"
	"os"
	"runtime"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("simplecli", version)
		fmt.Println("commit:", commit)
		fmt.Println("built:", date)
		return
	}

	fmt.Printf("Operating System: %s\nArchitecture: %s.\n", runtime.GOOS, runtime.GOARCH)
}
