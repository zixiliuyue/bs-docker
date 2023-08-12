

package qcloud

import (
	"reflect"
	"testing"

	loadbalance "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
)

func TestGetClusterIDPostfix(t *testing.T) {
	cases := []struct {
		info   string
		inStr  string
		outStr string
	}{
		{
			info:   "",
			inStr:  "BCS-MESOS-3344",
			outStr: "3344",
		},
		{
			info:   "",
			inStr:  "BCS-MESOS-",
			outStr: "",
		},
		{
			info:   "",
			inStr:  "xxxxxxx",
			outStr: "xxxxxxx",
		},
	}
	for i, c := range cases {
		t.Logf("index %d, info %s", i, c.info)
		tmpStr := GetClusterIDPostfix(c.inStr)
		if tmpStr != c.outStr {
			t.Errorf("expect %s, but get %s", c.outStr, tmpStr)
		} else {
			t.Log("success")
		}
	}
}

func TestGetBackendSegement(t *testing.T) {
	cases := []struct {
		info         string
		inSlice      []*loadbalance.Backend
		inCur        int
		inSegmentLen int
		outSlice     []*loadbalance.Backend
	}{
		{
			info:         "nil slice",
			inSlice:      nil,
			inCur:        1,
			inSegmentLen: 10,
			outSlice:     nil,
		},
		{
			info: "normal test",
			inSlice: []*loadbalance.Backend{
				{
					IP: "a",
				},
				{
					IP: "b",
				},
				{
					IP: "c",
				},
				{
					IP: "d",
				},
				{
					IP: "e",
				},
			},
			inCur:        1,
			inSegmentLen: 3,
			outSlice: []*loadbalance.Backend{
				{
					IP: "b",
				},
				{
					IP: "c",
				},
				{
					IP: "d",
				},
			},
		},
		{
			info: "normal test",
			inSlice: []*loadbalance.Backend{
				{
					IP: "a",
				},
				{
					IP: "b",
				},
				{
					IP: "c",
				},
				{
					IP: "d",
				},
				{
					IP: "e",
				},
			},
			inCur:        2,
			inSegmentLen: 5,
			outSlice: []*loadbalance.Backend{
				{
					IP: "c",
				},
				{
					IP: "d",
				},
				{
					IP: "e",
				},
			},
		},
	}

	for i, c := range cases {
		t.Logf("index %d, info %s", i, c.info)
		ret := GetBackendsSegment(c.inSlice, c.inCur, c.inSegmentLen)
		if !reflect.DeepEqual(ret, c.outSlice) {
			t.Errorf("expect %v, but get %v", c.outSlice, ret)
		} else {
			t.Log("success")
		}
	}
}
