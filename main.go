package main

import (
	"github.com/masibw/aidx/cmd"
	"os"
)

func main() {
	query := os.Args[1]

	cmd.Run(query)
}
