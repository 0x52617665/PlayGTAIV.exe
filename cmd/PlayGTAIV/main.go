package main

import (
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("LaunchGTAIV").Run()
	if err != nil {
		log.Fatal(err)
	}
}
