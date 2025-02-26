package main

import (
	"encoding/json"
	"fmt"
	"os"

	thehive "github.com/Duongdot/hive4go"
	"github.com/bitly/go-simplejson"
)

func main() {
	hive := thehive.CreateLogin("http://10.11.161.24:9000", "R1xbVBu7RZh/C9xa6nsePkuNvabKlp2S", false)
	query := `{"query": {"_in": {"_field": "title", "_values": ["Domain"]}}}`
	resp, err := hive.FindCases([]byte(query))
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		os.Exit(1)
	}
	temp, _ := json.Marshal(resp.Detail)
	fmt.Println(string(temp))
	jsondata, err := simplejson.NewJson(resp.Raw)
	fmt.Println(jsondata)
	ret, _ := jsondata.EncodePretty()
	fmt.Println(string(ret))
}
