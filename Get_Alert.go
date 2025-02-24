package main

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
	thehive "github.com/frikky/hive4go"
)

func main() {
	// Kết nối đến TheHive
	hive := thehive.CreateLogin("http://localhost:9000/thehive", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)

	// Lấy log của task từ TheHive
	alertID := "~32800" // Thay bằng Alert ID thực tế
	ret, err := hive.GetAlert(alertID)
	if err != nil {
		fmt.Println("Lỗi khi lấy alert Logs:", err)
		os.Exit(1)
	}

	// Parse JSON phản hồi
	jsonData, err := simplejson.NewJson(ret.Raw)
	if err != nil {
		fmt.Println("Lỗi khi parse JSON:", err)
		os.Exit(1)
	}

	// In kết quả logs
	formattedJSON, _ := jsonData.EncodePretty()
	fmt.Println("Alert Logs:")
	fmt.Println(string(formattedJSON))

	// Lấy ID của alert từ JSON response
	alertID = jsonData.Get("id").MustString() // Sửa lại từ ":=" thành "="
	if alertID == "" {
		fmt.Println("Không tìm thấy Alert ID trong response!")
		os.Exit(1)
	}
	response, err := hive.GetAlert(alertID)
	if err != nil {
		fmt.Println("Lỗi khi lấy Alert:", err)
		os.Exit(1)
	}
	// Parse JSON phản hồi alert
	jsonData, err = simplejson.NewJson(response.Raw)
	if err != nil {
		fmt.Println("Lỗi khi parse JSON alert:", err)
		os.Exit(1)
	}

	// Lấy thông tin chi tiết của alert
	fmt.Printf("Get alert %s\n", alertID)
	fmt.Println("--------------------------")

	// In thông tin chi tiết của alert
	formattedJSON, _ = jsonData.EncodePretty()
	fmt.Println("Alert Details:")
	fmt.Println(string(formattedJSON))
}
