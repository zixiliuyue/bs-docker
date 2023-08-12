

package alert

import (
	"os"
	"testing"
	"time"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd/config"
)

var defaultOptions = &config.AlertServerOptions{
	Server:      "http://xxx/prod",
	AppCode:     "xxx",
	AppSecret:   "xxx",
	ServerDebug: true,
}

func TestNewAlertServer(t *testing.T) {
	client := NewAlertServer(defaultOptions, WithTestDebug(true))

	data := []AlarmReqData{
		{
			StartsTime:   time.Now(),
			EndsTime:     time.Now().Add(60 * time.Hour),
			GeneratorURL: "http://xxx",
			Annotations: map[string]string{
				"uuid":    "cee84faf-7ee3-11ea-xxx",
				"message": "0.gseagent.gse.30012.1586932748085923931()status changed:Staging->Failed",
			},
			Labels: map[string]string{
				"alertname":       "测试cee84faf",
				"project_id":      "5805f1b824134fa39318fb0cf59f694b",
				"cluster_id":      "BCS-K8S-40185",
				"namespace":       "gse",
				"ip":              "xx.xx.xx.xx",
				"module_name":     "scheduler",
				"app_alarm_level": "important",
				"reason":          "scheduler",
			},
		},
	}

	err := client.SendAlarmInfoToAlertServer(data, 10*time.Second)
	if err != nil {
		t.Fatal(err)
		os.Exit(1)
	}

	t.Log("call SendAlarmInfoToAlertServer successful")
}
