package main

import (
	"fmt"
	"os"

	thehive "github.com/Duongdot/hive4go"
	"github.com/bitly/go-simplejson"
)

func CreateCase() {
	hive := thehive.CreateLogin("http://127.0.0.1:9000/thehive", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)

	fmt.Println("Create Case")
	fmt.Println("--------------------------")
	response, err := hive.CreateCase(
		"hive4go",  // Title
		"N/A",      // Description
		1,          // TLP
		1,          // Severity
		[]string{}, // Tags
		true,
	)

	if err != nil {
		fmt.Println("Error creating case:", err)
		os.Exit(1)
	}

	jsonData, err := simplejson.NewJson(response.Raw)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		os.Exit(1)
	}

	ret, _ := jsonData.EncodePretty()
	fmt.Println(string(ret))

	id := jsonData.Get("id").MustString()
	fmt.Printf("Get created case %s\n", id)
	fmt.Println("--------------------------")

	caseResponse, err := hive.GetCase(id)
	if err != nil {
		fmt.Println("Error getting case:", err)
		os.Exit(1)
	}

	jsonData, err = simplejson.NewJson(caseResponse.Raw)
	if err != nil {
		fmt.Println("Error parsing case response:", err)
		os.Exit(1)
	}

	ret, _ = jsonData.EncodePretty()
	fmt.Println(string(ret))
}

// Hàm main để chạy chương trình
func main() {
	CreateCase()
}
