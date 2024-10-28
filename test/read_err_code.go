package main

import (
	"encoding/json"
	"fmt"
	"gvb-server/models/res"
	"os"
)

const file = "models/res/err_code.json"

type ErrorMap map[res.ErrorCode]string

func main() {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return
	}
	errorMap := ErrorMap{}
	err = json.Unmarshal(bytes, &errorMap)
	if err != nil {
		return
	}
	fmt.Println(errorMap)
	fmt.Println(errorMap[res.SettingsError])
}
