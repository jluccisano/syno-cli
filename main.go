package main

import (
	"log"
	"os"
	"gopkg.in/alecthomas/kingpin.v1"
	"frister.net/go/syno-cli/synoapi"
)

var (
	enable          = kingpin.Command("enable", "Enable camera")
	disable          = kingpin.Command("disable", "Disable camera")
)

func main() {
	// something like https://myds.example.net:5001
	api_base := os.Getenv("SYNO_BASE_URL")
	user := os.Getenv("SYNO_USER")
	password := os.Getenv("SYNO_PASSWORD")

	command := kingpin.Parse()

	client := synoapi.NewClient(api_base)
	err := client.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	default:
		kingpin.Usage()
	case "enable":
		enableCamera(client)
	case "disable":
		disableCamera(client)
	}
}

func enableCamera(client *synoapi.Client) {
	err := client.Enable()
	if err != nil {
		log.Fatalf("Locking failed: %v", err)
	}
}

func disableCamera(client *synoapi.Client) {
	err := client.Disable()
	if err != nil {
		log.Fatalf("Locking failed: %v", err)
	}
}
