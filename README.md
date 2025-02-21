# Hive4Go - TheHive API Client for Go

Hive4Go là một thư viện Go giúp tương tác với API của TheHive, một nền tảng quản lý sự cố an ninh mạng mã nguồn mở.

## Cài đặt

Để sử dụng thư viện Hive4Go, trước tiên bạn cần cài đặt Go và chạy lệnh sau để tải thư viện:

```sh
go get github.com/frikky/hive4go
```

Ngoài ra, bạn cũng cần đảm bảo các dependencies sau được cài đặt:

```sh
github.com/bitly/go-simplejson v0.5.1 // indirect
github.com/frikky/hive4go v0.0.0-20200903125636-d374532a56cd // indirect
github.com/google/go-querystring v1.1.0 // indirect
github.com/levigross/grequests v0.0.0-20231203190023-9c307ef1f48d // indirect
github.com/satori/go.uuid v1.2.0 // indirect
golang.org/x/net v0.19.0 // indirect
```

## Sử dụng

### 1. Import thư viện

```go
package main

import (
    "fmt"
    "github.com/frikky/hive4go"
)

func main() {
    client := hive4go.NewClient("https://your-thehive-instance.com", "your-api-key")
    fmt.Println("Hive4Go client initialized.")
}
```

### 2. Lấy danh sách các cases

```go
cases, err := client.GetCases()
if err != nil {
    fmt.Println("Lỗi khi lấy cases:", err)
} else {
    fmt.Println("Danh sách cases:", cases)
}
```

### 3. Tạo một case mới

```go
caseData := hive4go.Case{
    Title: "New Incident",
    Description: "Investigation required",
    Severity: 2,
}

newCase, err := client.CreateCase(caseData)
if err != nil {
    fmt.Println("Lỗi khi tạo case:", err)
} else {
    fmt.Println("Case mới đã được tạo:", newCase)
}
```

## Thông tin thêm
- Tài liệu chính thức của TheHive API: [https://docs.thehive-project.org](https://docs.thehive-project.org)
- Repository Hive4Go: [https://github.com/frikky/hive4go](https://github.com/frikky/hive4go)

