package main

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/frikky/hive4go"
)

func main() {
	// Kết nối đến TheHive
	hive := thehive.CreateLogin("http://localhost:9000/thehive", "qqsYzkVEoEA/FHRW7ecB4llNOKIeEPm+", false)

	// Task ID cần lấy thông tin
	taskID := "~45072" // Thay bằng Task ID thực tế

	// Gọi API để lấy thông tin Task
	ret, err := hive.GetTask(taskID)
	if err != nil {
		fmt.Println("❌ Lỗi khi lấy Task:", err)
		os.Exit(1)
	}

	// Kiểm tra phản hồi có dữ liệu không
	if len(ret.Raw) == 0 {
		fmt.Println("❌ Phản hồi trống, không có dữ liệu Task!")
		os.Exit(1)
	}

	// Parse JSON phản hồi
	jsonData, err := simplejson.NewJson(ret.Raw)
	if err != nil {
		fmt.Println("❌ Lỗi khi parse JSON:", err)
		os.Exit(1)
	}

	// Lấy ID của Task từ JSON
	parsedTaskID := jsonData.Get("id").MustString()
	if parsedTaskID == "" {
		fmt.Println("❌ Không tìm thấy Task ID trong response!")
		os.Exit(1)
	}

	// In thông tin chi tiết của Task
	formattedJSON, _ := jsonData.EncodePretty()
	fmt.Println("✅ Task Details:")
	fmt.Println(string(formattedJSON))
}
