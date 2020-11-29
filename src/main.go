package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("TASKKILL", "/F", "/IM", "System.exe")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// res := exec.Command("explorer.exe")

	// erre := res.Run()

	// if erre != nil {
	// 	log.Fatal(err)
	// }
}
