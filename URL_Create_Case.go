package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// type Hive_Acc struct {
// 	url    string
// 	apik   string
// 	hder   map[string]string
// 	verify bool
// }
// type CaseData struct {
// 	Title            string                 `json:"title"`
// 	Description      string                 `json:"description"`
// 	Tlp              int                    `json:"tlp"`
// 	Severity         int                    `json:"severity"`
// 	Date             int64                  `json:"date,omitempty"`
// 	Tags             []string               `json:"tags"`
// 	Tasks            []CaseTask             `json:"tasks"`
// 	Flag             bool                   `json:"flag"`
// 	Owner            string                 `json:"owner"`
// 	Status           string                 `json:"status"`
// 	CreatedAt        int64                  `json:"createdAt"`
// 	CustomFields     map[string]interface{} `json:"customFields"`
// 	Id               string                 `json:"id"`
// 	Summary          string                 `json:"summary"`
// 	ResolutionStatus string                 `json:"resolutionStatus"`
// 	ImpactStatus     string                 `json:"impactStatus"`
// 	Raw              []byte                 `json:"-"`
// }

//	func Login(url string, api string, verify_t bool) Hive_Acc {
//		formattedApi := fmt.Sprintf("Bearer %s", api)
//		return Hive_Acc{
//			url:  url,
//			apik: api,
//			hder: map[string]string{
//				"Content-Type":  "application/json", //header: Content-Type
//				"Authorization": formattedApi,    //header: Authorization
//			},
//			verify: verify_t,
//		}
//	}
func URLCreateCase() {
	// 1️⃣ Định nghĩa URL API
	url := "http://localhost:9000/thehive/api/case"

	// 2️⃣ Tạo dữ liệu JSON (Case mới)
	// caseData := map[string]interface{}{
	// 	"title":       "VanvyTest",
	// 	"description": "testing",
	// 	"tags":        []string{"golang", "thehive"},
	// }
	caseData, err := os.ReadFile("case.json")
	fmt.Println(caseData)
	// var jsonData_t map[string]interface{}
	// err = json.NewDecoder(caseData).Decode(&jsonData_t)
	// 3️⃣ Chuyển struct thành JSON
	// jsonData, err := json.Marshal(jsonData_t)
	// if err != nil {
	// 	fmt.Println("Error encoding JSON:", err)
	// 	return
	// }

	// 4️⃣ Tạo HTTP request với method POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(caseData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 5️⃣ Thêm headers (API Key & Content-Type)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer dxA/NEvjl/hjcvv5rQ9aG/juLExgrHn6")

	// 6️⃣ Gửi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 7️⃣ Đọc phản hồi từ server
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response:", err)
	// 	return
	// }

	// // 8️⃣ In kết quả
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))
	fmt.Println("Response Status:", resp)
}
