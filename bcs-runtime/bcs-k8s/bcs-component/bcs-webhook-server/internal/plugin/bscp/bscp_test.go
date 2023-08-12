

package bscp

import (
	"testing"
)

func TestAddPathIntoAppInfoMode(t *testing.T) {
	tests := []struct {
		inputStr  string
		inputPath string
		outputStr string
		isError   bool
	}{
		{
			"[{\"business\":\"mars\",\"app\":\"app\",\"cluster\":\"sz\",\"zone\":\"zone1\",\"dc\":\"szdc\"}]",
			"/aaa/bbb",
			"[{\"business\":\"mars\",\"app\":\"app\",\"cluster\":\"sz\",\"zone\":\"zone1\",\"dc\":" +
				"\"szdc\",\"path\":\"/aaa/bbb\"}]",
			false,
		},
		{
			"[{\"business\":\"mars\",\"app\":\"app\"}]",
			"/a/a/b",
			"",
			true,
		},
	}
	for index, test := range tests {
		t.Logf("%d test", index)
		tmpStr, err := AddPathIntoAppInfoMode(test.inputStr, test.inputPath)
		if err != nil && !test.isError {
			t.Errorf("expect no error, but get err %s", err.Error())
		}
		if err == nil && test.isError {
			t.Errorf("expect error, but no error")
		}
		if tmpStr != test.outputStr {
			t.Errorf("expect %s, but get %s", test.outputStr, tmpStr)
		}
	}
}
