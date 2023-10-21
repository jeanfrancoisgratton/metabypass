package main

import (
	"fmt"
	"metabypass/cmd"
	"os"
)

func main() {
	currentdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
	os.Chdir(currentdir)
}
