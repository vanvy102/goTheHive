package main

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/frikky/hive4go"
)

func CreateTask() {
	// Kết nối TheHive
	hive := thehive.CreateLogin("http://localhost:9000/thehive", "qqsYzkVEoEA/FHRW7ecB4llNOKIeEPm+", false)

	fmt.Println("Create Task")
	fmt.Println("--------------------------")

	// Tạo task với trạng thái hợp lệ
	task := thehive.CaseTask{
		Title:       "Task from python",
		Status:      "Waiting",  // ✅ Giá trị hợp lệ
		Owner:       "testcase@thehive",
		Description: "created from Golang",
		Flag:        false,
	}

	// Gửi request tạo task
	newTask, err := hive.CreateCaseTask("2", task)
	if err != nil {
		fmt.Println("Error creating task:", err)
		os.Exit(1)
	}

	// Đọc JSON response
	jsonData, err := simplejson.NewJson(newTask.Raw)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		os.Exit(1)
	}
	ret, _ := jsonData.EncodePretty()
	fmt.Println(string(ret))

	// Lấy ID của task
	id := jsonData.Get("id").MustString()
	if id == "" {
		fmt.Println("Error: Task ID is empty")
		os.Exit(1)
	}
	fmt.Println("Created Task ID:", id)  // Kiểm tra ID

	// Lấy thông tin task vừa tạo
	fmt.Printf("Get created task %s\n", id)
	fmt.Println("--------------------------")

	// response, err := hive.GetTask(id)
	// if err != nil {
	// 	fmt.Println("Error retrieving task:", err)
	// 	os.Exit(1)
	// }

	// jsonData, err = simplejson.NewJson(response.Raw)
	// if err != nil {
	// 	fmt.Println("Error parsing JSON response:", err)
	// 	os.Exit(1)
	// }
	// ret, _ = jsonData.EncodePretty()
	// fmt.Println(string(ret))
}
