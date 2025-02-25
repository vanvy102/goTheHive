package main

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
	thehive "github.com/frikky/hive4go"
	uuid "github.com/satori/go.uuid"
)

func CreateAlert() {
	hive := thehive.CreateLogin("http://localhost:9000/thehive", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)

	// Missing file
	artifacts := []thehive.Artifact{
		thehive.AlertArtifact("ip", "8.8.8.8", 0, []string{}, false),
		thehive.AlertArtifact("domain", "google.com", 0, []string{}, false),
		//thehive.AlertArtifact("file", "pic.png", 0, []string{}, 0)
		//thehive.AlertArtifact("file", "sample.txt", 0, []string{}, 0)
	}

	sourceRef := uuid.NewV4().String()
	fmt.Println("Create Alert")
	fmt.Println("--------------------------")
	alert, err := hive.CreateAlert(
		artifacts,                     // Artifacts
		"helo",                        // Title
		"HELO",                        // Description
		1,                             // TLP
		1,                             // Severity
		[]string{"hive4go", "sample"}, // Tags
		"SIEM",                        // Type
		"Carbon black",                // Source
		sourceRef,                     // SourceRef
		"20250219",
		"New", // Status (ví dụ: "New", "InProgress", "Closed")
	)

	jsonData, err := simplejson.NewJson(alert.Raw)
	ret, _ := jsonData.EncodePretty()
	fmt.Println(string(ret))

	id := jsonData.Get("id").MustString()

	// Get all the details of the created alert
	fmt.Printf("Get created alert %s\n", id)
	fmt.Println("--------------------------")

	response, err := hive.GetAlert(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonData, err = simplejson.NewJson(response.Raw)
	ret, _ = jsonData.EncodePretty()
	fmt.Println(string(ret))
}
