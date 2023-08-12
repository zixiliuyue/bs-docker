

package alertmanager

import (
	"testing"
	"time"
)

func newAlertManager() AlertManageInterface {
	options := Options{
		Server:     "https://xxx:xxx/xx/v4",
		ClientAuth: true,
		Debug:      true,
		Token:      "dFYn6pFOouFePmpKlfBPoBaNbFbnoSJX",
	}

	alert, _ := NewAlertManager(options)
	return alert
}

func TestAlertManager_CreateAlertInfo(t *testing.T) {
	alertClient := newAlertManager()

	req := &CreateBusinessAlertInfoReq{
		Starttime:    time.Now().Unix(),
		Endtime:      time.Now().Add(24 * time.Hour).Unix(),
		Generatorurl: "http://123456",
		AlarmType:    "module",
		ClusterID:    "bcs-2048",
		AlertAnnotation: &AlertAnnotation{
			Message: "cpu test",
		},
		ModuleAlertLabel: &ModuleAlertLabel{
			ModuleName: "bcs-scheduler",
			ModuleIP:   "1.1.1.1",
			AlarmName:  "cpu负载过高",
			AlarmLevel: "warning",
		},
	}
	err := alertClient.CreateAlertInfoToAlertManager(req, time.Second*10)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}
