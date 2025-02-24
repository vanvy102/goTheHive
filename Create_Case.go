package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bitly/go-simplejson"
)

func main() {
	// hive := thehive.CreateLogin("http://127.0.0.1:9000", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)
	hive, err := thehive.NewClient("http://localhost:9000/thehive", &thehive.ClientOpts{
		Auth: &thehive.APIAuth{
			APIKey: "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	tasks := []thehive.CaseTask{
		thehive.CaseTask{Title: "Tracking"},
		thehive.CaseTask{Title: "Communication"},
		thehive.CaseTask{Title: "Investigation", Status: "Waiting", Flag: true},
	}

	fmt.Println("Create Case")
	fmt.Println("--------------------------")
	Case, err := hive.CreateCase(
		"From hive4go", // Title
		"N/A",          // Description
		3,              // TLP
		1,              // Severity
		tasks,          // Tasks
		[]string{},     // Tags
		true,           // Flag
	)

	if err != nil || Case.StatusCode != 201 {
		fmt.Println(err, Case)
		os.Exit(1)
	}

	jsonData, err := simplejson.NewJson(Case.Bytes())
	ret, _ := jsonData.EncodePretty()
	fmt.Println(string(ret))

	id := jsonData.Get("id").MustString()

	// Get all the details of the created case
	fmt.Printf("Get created case %s\n", id)
	fmt.Println("--------------------------")
	response, err := hive.GetCase(id)
	if err != nil {
		fmt.Println(err, response.StatusCode)
		os.Exit(1)
	}

	jsonData, err = simplejson.NewJson(response.Bytes())
	ret, _ = jsonData.EncodePretty()
	fmt.Println(string(ret))

	// Add a new task to the created case
	fmt.Printf("Add a task %s\n", id)
	fmt.Println("--------------------------")
	response, err = hive.CreateCaseTask(
		id,
		thehive.CaseTask{
			Title:  "Yet another Task",
			Status: "InProgress",
			Owner:  "admin",
			Flag:   true,
		},
	)

	if err != nil || response.StatusCode != 201 {
		fmt.Println(err, response.StatusCode)
		os.Exit(1)
	}

	jsonData, err = simplejson.NewJson(response.Bytes())
	ret, _ = jsonData.EncodePretty()
	fmt.Println(string(ret))
}
