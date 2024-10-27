package main

import (
	"fmt"
	"log"
	"os"
)

const version = "1.0.0"

func main() {
	printLogo()
	printVersion()
	printLicense()
}

func printLogo() {
	data, err := os.ReadFile("logo.txt")
	if err != nil {
		log.Fatalf("Error reading logo file: %v", err)
	}
	fmt.Println(string(data))
}

func printVersion() {
	fmt.Printf("client app: %s\n", version)
}

func printLicense() {
	data, err := os.ReadFile("LICENSE")
	if err != nil {
		log.Fatalf("Error reading LICENSE file: %v", err)
	}
	fmt.Println(string(data))
}
