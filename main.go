package main

import (
	"aidx/cmd"
	"os"
)

func main() {
	query := os.Args[1]

	cmd.Run(query)
}
