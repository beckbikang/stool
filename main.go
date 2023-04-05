package main

import (
	"log"
	"stool/cmd"
)

func main() {
	log.Println("start running stool")
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
