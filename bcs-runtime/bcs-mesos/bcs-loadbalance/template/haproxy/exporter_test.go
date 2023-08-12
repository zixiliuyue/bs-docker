

package haproxy

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	s := &Server{
		Name:           "test_server",
		MaxQueue:       238473,
		MaxSessionRate: 23824,
	}
	if getValue(s, "MaxQueue") != 238473 {
		t.Errorf("test failed")
	}
	if getValue(s, "MaxSessionRate") != 23824 {
		t.Errorf("test failed")
	}
	if getValue(s, "Name") != 0 {
		t.Errorf("test failed")
	}
}

func TestConvertStatus(t *testing.T) {
	tests := []struct {
		str   string
		value float64
	}{
		{
			str:   "UP",
			value: 1,
		},
		{
			str:   "UP 1/3",
			value: 1,
		},
		{
			str:   "UP 2/3",
			value: 1,
		},
		{
			str:   "no check",
			value: 1,
		},
		{
			str:   "DOWN",
			value: 0,
		},
		{
			str:   "DOWN 1/2",
			value: 0,
		},
		{
			str:   "NOLB",
			value: 0,
		},
		{
			str:   "MAINT",
			value: 0,
		},
		{
			str:   "invalidstr",
			value: 0,
		},
	}
	for _, test := range tests {
		realValue := convertStatus(test.str)
		if realValue != test.value {
			t.Errorf("test failed, convertStatus(%s) expected %f but get %f", test.str, test.value, realValue)
		}
	}
}

func TestConvertCheckStatus(t *testing.T) {
	tests := []struct {
		str   string
		value float64
	}{
		{
			str:   "L4OK",
			value: 1,
		},
		{
			str:   "L6OK",
			value: 1,
		},
		{
			str:   "L7OK",
			value: 1,
		},
		{
			str:   "invalidstr",
			value: 0,
		},
	}
	for _, test := range tests {
		realValue := convertCheckStatus(test.str)
		if realValue != test.value {
			t.Errorf("test failed, convertCheckStatus(%s) expected %f but get %f", test.str, test.value, realValue)
		}
	}
}
