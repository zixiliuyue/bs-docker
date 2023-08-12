

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Output2Json 格式化输出json
func Output2Json(i interface{}) {
	jsonStr, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, jsonStr, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
