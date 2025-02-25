// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	thehive "github.com/Duongdot/hive4go"
// )

// type HiveCaseResp struct {
// 	Id    string `json:"id"`
// 	Title string `json:"title"`
// }

// func main() {
// 	// Đăng nhập TheHive
// 	hive := thehive.CreateLogin("http://localhost:9000/thehive", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)

// 	// Kiểm tra nếu đăng nhập thất bại
// 	if hive.Url == "" {
// 		fmt.Println("Lỗi khi đăng nhập vào TheHive.")
// 		os.Exit(1)
// 	}

// 	// Tạo query để lấy danh sách cases
// 	query := []byte(`{"query": {"_name": "listCases"}}`)

// 	// Gửi truy vấn để tìm cases
// 	cases, err := hive.FindCases(query)
// 	if err != nil {
// 		fmt.Println("Lỗi khi lấy danh sách cases:", err)
// 		os.Exit(1)
// 	}

// 	// Kiểm tra có cases nào không
// 	if len(cases.Detail) == 0 {
// 		fmt.Println("Không có case nào trong hệ thống.")
// 		return
// 	}

// 	// Lấy case đầu tiên
// 	caseID := cases.Detail[0].Id // Sửa lỗi tại đây
// 	fmt.Println("Lấy thông tin case với ID:", caseID)

// 	caseData, err := hive.GetCase(caseID)
// 	if err != nil {
// 		fmt.Println("Lỗi khi lấy thông tin case:", err)
// 		return
// 	}

// 	// Parse JSON response
// 	var caseResp HiveCaseResp
// 	err = json.Unmarshal(caseData.Raw, &caseResp)
// 	if err != nil {
// 		fmt.Println("Lỗi khi parse dữ liệu case:", err)
// 		return
// 	}

// 	// Hiển thị thông tin case
// 	fmt.Println("Thông tin case:")
// 	fmt.Printf("ID: %s\n", caseResp.Id)
// 	fmt.Printf("Tên: %s\n", caseResp.Title)
// }

package main

import (
	"fmt"

	thehive "github.com/frikky/hive4go"
)

func main() {
	hive := thehive.CreateLogin("http://localhost:9000/thehive", "dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6", false)

	ret, err := hive.GetCase("2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ret)
}
